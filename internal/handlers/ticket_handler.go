package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"mpango-wa-cuddles/internal/models"
	"mpango-wa-cuddles/internal/services"
)

type ticketRequest struct {
	LocationID uuid.UUID `json:"location_id" binding:"required"`
	Phone      string    `json:"phone" binding:"required"`
}

// InitiateTicket simulates STK push initiation.
func InitiateTicket(c *gin.Context) {
	var req ticketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var loc models.Location
	if err := db.First(&loc, "id = ?", req.LocationID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "location not found"})
		return
	}

	now := time.Now()
	ticket := models.Ticket{
		ID:          uuid.New(),
		LocationID:  req.LocationID,
		Price:       loc.CostPerPerson,
		MpesaRef:    "PENDING",
		UserPhone:   req.Phone,
		Status:      "pending",
		CreatedAt:   now,
		LastUpdated: now,
	}
	if ref, err := services.InitiateSTK(req.Phone, loc.CostPerPerson, loc.Name); err == nil {
		ticket.MpesaRef = ref
	}
	if err := db.Create(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create ticket"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"ticket_id": ticket.ID, "status": ticket.Status})
}

// TicketStatus returns a mocked payment status.
func TicketStatus(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var ticket models.Ticket
	if err := db.First(&ticket, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ticket not found"})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

// MpesaCallback updates ticket status after STK push.
func MpesaCallback(c *gin.Context) {
	var payload struct {
		TicketID uuid.UUID `json:"ticket_id"`
		Status   string    `json:"status"`
		MpesaRef string    `json:"mpesa_ref"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if payload.Status == "" {
		payload.Status = "paid"
	}

	if err := db.Model(&models.Ticket{}).
		Where("id = ?", payload.TicketID).
		Updates(map[string]interface{}{
			"status":       payload.Status,
			"mpesa_ref":    payload.MpesaRef,
			"last_updated": time.Now(),
		}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "callback processed"})
}
