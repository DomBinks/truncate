package main

import (
	"encoding/gob"
	"log"
	"os"
	"truncate/internal/authenticator"
	"truncate/internal/callback"
	"truncate/internal/handlers"
	"truncate/internal/login"
	"truncate/internal/logout"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	/* gin.SetMode(gin.ReleaseMode) // Set Gin to use the release mode */
	router := gin.Default() // Get a router for the web server

	// Register custom types to store in cookies
	gob.Register(map[string]interface{}{})
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load environment variables")
	}

	// Get an authenticator
	auth, err := authenticator.New()
	if err != nil {
		log.Fatal("Unable to initialise the authenticator")
	}

	// Set router paths to use handlers

	// Paths to pages viewed by the user
	router.GET("/", handlers.Index) // Index as handled by Angular
	router.GET("/:file", handlers.Static)
	router.GET("/~:shortened", handlers.URL)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/logout", logout.Handler)
	router.GET("/invalid", handlers.Index) // Index as handled by Angular
	router.GET("/profile", handlers.Index) // Index as handled by Angular
	router.GET("/get-login", handlers.GetLogin)

	// Paths used to provide an API to the frontend
	router.POST("/shorten", handlers.Shorten)
	router.POST("/get-urls", handlers.GetURLs)
	router.POST("/delete-row", handlers.DeleteRow)

	router.Run(os.Getenv("ROUTER") + ":8080") // Run the web server
}
