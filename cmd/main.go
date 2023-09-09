package main

import (
	"encoding/gob"
	"log"
	"url-shortener/cmd/authenticator"
	"url-shortener/cmd/callback"
	"url-shortener/cmd/handlers"
	"url-shortener/cmd/login"
	"url-shortener/cmd/logout"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
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
	router.GET("/", handlers.Index)
	router.GET("/:file", handlers.Static)
	router.GET("/link/:shortened", handlers.URL)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/logout", logout.Handler)
	router.GET("/invalid", handlers.Index)
	router.GET("/profile", handlers.Index)
	router.GET("/get-login-UI", handlers.GetLoginUI)

	router.POST("/shorten", handlers.Shorten)
	router.POST("/get-urls", handlers.GetURLs)
	router.POST("/delete-row", handlers.DeleteRow)

	router.Run("localhost:8080") // Run the web server
}
