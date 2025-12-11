package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mpango-wa-cuddles/internal/models"
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

	ticket := models.Ticket{
		ID:          uuid.New(),
		LocationID:  req.LocationID,
		Price:       500, // demo price
		MpesaRef:    "MOCKREF",
		UserPhone:   req.Phone,
		Status:      "pending",
		CreatedAt:   time.Now(),
		LastUpdated: time.Now(),
	}
	demoTickets = append(demoTickets, ticket)

	c.JSON(http.StatusAccepted, gin.H{"ticket_id": ticket.ID, "status": ticket.Status})
}

// TicketStatus returns a mocked payment status.
func TicketStatus(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	for _, t := range demoTickets {
		if t.ID == id {
			c.JSON(http.StatusOK, t)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "ticket not found"})
}


