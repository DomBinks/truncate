package helpers

import (
	"database/sql"
	"log"
	"math/rand"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Get a pointer to a struct that represents the database
func GetDatabase() *sql.DB {
	// Set the credentials for connecting to the database
	const (
		host     = "localhost"
		port     = 5432
		user     = "urlshortener"
		password = "golang"
		dbname   = "urlshortener"
	)

	// Put the credentials into a string
	connectionString := "host=" + host + " port =" +
		strconv.Itoa(port) + " user=" + user + " password=" +
		password + " dbname=" + dbname + " sslmode=disable"

	log.Println("Connecting to database")

	// Connect to the database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return db // Return a pointer to the database
}

// Get the unique identifier for the current user
func GetID(c *gin.Context) string {
	session := sessions.Default(c)    // Get the current session
	profile := session.Get("profile") // Get the session's profile data

	// If the user is logged in
	if profile != nil {
		// Put the profile data from the session into a map
		profileMap := profile.(map[string]interface{})

		// Return the value of the sub key in profileMap as it
		// uniquely identifies the user
		return profileMap["sub"].(string)
	} else {
		// If the user isn't logged in
		return "default"
	}
}

// Generate the random string used to create the shortened URL
func GenerateShortened() string {
	// Possible symbols to choose from
	symbols := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var shortened string // The string to output

	// Infinite loop
	for {
		shortened = "" // Start with an empty string

		// Add 6 random symbols to the string
		for i := 0; i < 6; i++ {
			shortened += string(symbols[rand.Intn(len(symbols))])
		}

		// Conenct to the database
		db := GetDatabase()
		rows, err := db.Query("SELECT name FROM urls WHERE short='" + shortened + "';")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// If this string hasn't been used yet
		if !rows.Next() {
			break // Break out the infinite loop
		}
	}

	return shortened // Return the generated string
}
