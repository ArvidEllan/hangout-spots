package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"mpango-wa-cuddles/internal/config"
	"mpango-wa-cuddles/internal/handlers"
	"mpango-wa-cuddles/internal/middleware"
)

// New configures the HTTP router with public endpoints.
func New(cfg config.Config) *gin.Engine {
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

	store := cookie.NewStore([]byte(cfg.JWTSecret))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 14, // 14 days
		HttpOnly: true,
	})
	r.Use(sessions.Sessions("cuddle_session", store))

	// Public API
	r.GET("/healthz", handlers.Health)
	r.GET("/locations", handlers.ListLocations)
	r.GET("/locations/:id", handlers.GetLocation)
	r.GET("/ads", handlers.ListAds)
	r.POST("/ads/:id/click", handlers.TrackAdClick)
	r.POST("/tickets/initiate", handlers.InitiateTicket)
	r.GET("/tickets/status/:id", handlers.TicketStatus)
	r.POST("/tickets/mpesa/callback", handlers.MpesaCallback)
	r.POST("/cuddlelist/:id", handlers.SessionSaveLocation)
	r.GET("/cuddlelist", handlers.SessionList)

	// Placeholder admin group
	auth := r.Group("/user")
	auth.Use(middleware.RequireAuth([]byte(cfg.JWTSecret)))
	{
		auth.POST("/save/:id", handlers.SaveLocation)
		auth.GET("/saved", handlers.SavedPlaces)
	}

	r.POST("/user/register", handlers.Register)
	r.POST("/user/login", handlers.Login)

	admin := r.Group("/admin")
	admin.Use(middleware.RequireAuth([]byte(cfg.JWTSecret)))
	{
		admin.POST("/locations", handlers.CreateLocation)
		admin.PUT("/locations/:id", handlers.UpdateLocation)
		admin.DELETE("/locations/:id", handlers.DeleteLocation)
		admin.POST("/ads", handlers.CreateAd)
		admin.PUT("/ads/:id", handlers.UpdateAd)
		admin.POST("/upload", handlers.UploadImage)
	}

	return r
}


