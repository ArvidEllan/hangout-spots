package handlers

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"mpango-wa-cuddles/internal/models"
)

type authRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Register creates a user and returns a JWT.
func Register(c *gin.Context) {
	var req authRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user := models.User{
		Email:    req.Email,
		Password: string(hashed),
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not create user"})
		return
	}

	token := issueToken(user.ID)
	c.JSON(http.StatusCreated, gin.H{"token": token})
}

// Login issues JWT on valid credentials.
func Login(c *gin.Context) {
	var req authRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token := issueToken(user.ID)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// SaveLocation adds a location to the user's cuddle list.
func SaveLocation(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	locID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid location id"})
		return
	}

	entry := models.Saved{UserID: userID, LocationID: locID}
	if err := db.Where("user_id = ? AND location_id = ?", userID, locID).
		FirstOrCreate(&entry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save location"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "saved"})
}

// SavedPlaces returns the user's cuddle list.
func SavedPlaces(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		// fallback to session-backed cuddle list
		session := sessions.Default(c)
		if raw := session.Get("cuddle_list"); raw != nil {
			if ids, ok := raw.([]string); ok && len(ids) > 0 {
				uuidIDs := make([]uuid.UUID, 0, len(ids))
				for _, idStr := range ids {
					if id, err := uuid.Parse(idStr); err == nil {
						uuidIDs = append(uuidIDs, id)
					}
				}
				if len(uuidIDs) > 0 {
					var locations []models.Location
					if err := db.Where("id IN ?", uuidIDs).Find(&locations).Error; err == nil {
						c.JSON(http.StatusOK, gin.H{"data": locations, "count": len(locations), "source": "session"})
						return
					}
				}
			}
		}
		c.JSON(http.StatusOK, gin.H{"data": []models.Location{}, "count": 0, "source": "session"})
		return
	}

	var saved []models.Saved
	if err := db.Where("user_id = ?", userID).Find(&saved).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load saved"})
		return
	}

	locIDs := make([]uuid.UUID, 0, len(saved))
	for _, s := range saved {
		locIDs = append(locIDs, s.LocationID)
	}

	var locations []models.Location
	if len(locIDs) > 0 {
		if err := db.Where("id IN ?", locIDs).Find(&locations).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load locations"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": locations, "count": len(locations)})
}

// SessionSaveLocation stores a cuddle list entry in the cookie session.
func SessionSaveLocation(c *gin.Context) {
	session := sessions.Default(c)
	locID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid location id"})
		return
	}

	raw := session.Get("cuddle_list")
	var ids []string
	if raw != nil {
		if existing, ok := raw.([]string); ok {
			ids = existing
		}
	}

	for _, id := range ids {
		if id == locID.String() {
			c.JSON(http.StatusOK, gin.H{"message": "already saved"})
			return
		}
	}

	ids = append(ids, locID.String())
	session.Set("cuddle_list", ids)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "saved"})
}

// SessionList returns cuddle list stored in session.
func SessionList(c *gin.Context) {
	session := sessions.Default(c)
	raw := session.Get("cuddle_list")
	ids, ok := raw.([]string)
	if !ok || len(ids) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": []models.Location{}, "count": 0, "source": "session"})
		return
	}

	uuidIDs := make([]uuid.UUID, 0, len(ids))
	for _, idStr := range ids {
		if id, err := uuid.Parse(idStr); err == nil {
			uuidIDs = append(uuidIDs, id)
		}
	}

	var locations []models.Location
	if len(uuidIDs) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": []models.Location{}, "count": 0, "source": "session"})
		return
	}
	if err := db.Where("id IN ?", uuidIDs).Find(&locations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load locations"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": locations, "count": len(locations), "source": "session"})
}

func issueToken(id uuid.UUID) string {
	claims := jwt.MapClaims{
		"sub": id.String(),
		"exp": time.Now().Add(72 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, _ := token.SignedString(jwtSecret)
	return signed
}

func getUserID(c *gin.Context) (uuid.UUID, bool) {
	raw, ok := c.Get("user_id")
	if !ok {
		return uuid.UUID{}, false
	}
	str, ok := raw.(string)
	if !ok {
		return uuid.UUID{}, false
	}
	id, err := uuid.Parse(str)
	if err != nil {
		return uuid.UUID{}, false
	}
	return id, true
}
