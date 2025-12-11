package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"mpango-wa-cuddles/internal/handlers"
)

// New configures the HTTP router with public endpoints.
func New() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://localhost:3000",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Public API
	r.GET("/healthz", handlers.Health)
	r.GET("/locations", handlers.ListLocations)
	r.GET("/locations/:id", handlers.GetLocation)
	r.GET("/ads", handlers.ListAds)
	r.POST("/tickets/initiate", handlers.InitiateTicket)
	r.GET("/tickets/status/:id", handlers.TicketStatus)

	// Placeholder admin group
	admin := r.Group("/admin")
	{
		admin.POST("/locations", handlers.CreateLocation)
		admin.PUT("/locations/:id", handlers.UpdateLocation)
		admin.DELETE("/locations/:id", handlers.DeleteLocation)
		admin.POST("/ads", handlers.CreateAd)
	}

	return r
}


