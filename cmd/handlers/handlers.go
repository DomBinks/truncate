package handlers

import (
	"database/sql"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"url-shortener/cmd/helpers"

	"github.com/gin-gonic/gin"
)

// Default handler
func Index(c *gin.Context) {
	// Get the page for this path using Angular routing
	c.File("web/dist/web/index.html")
}

// Handler for static files used by the front-end
func Static(c *gin.Context) {
	// Load the static file from the Angular folder
	c.File("web/dist/web/" + c.Param("file"))
}

// Handler for shortened URLs
func URL(c *gin.Context) {
	db := helpers.GetDatabase() // Get the database
	defer db.Close()

	url := c.Param("url") // Get the shortened URL
	var original string   // Original URL returned from the database

	// Get the original URL from the database, and store in the
	// original string
	err := db.QueryRow("SELECT original FROM urls WHERE short = '" + url + "';").Scan(&original)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Couldn't find the URL corresponding to the link")
		} else {
			log.Fatal(err)
		}
	}

	// Redirect the user to the original URL returned from the database
	c.Redirect(http.StatusPermanentRedirect, "http://"+original)

}

// Handler for getting the correct label and link for the login/logout
// button
func GetLoginUI(c *gin.Context) {
	id := helpers.GetID(c) // Get the user's ID

	// Return the label and link in a JSON
	if id != "default" {
		c.JSON(http.StatusOK, gin.H{"label": "Logout",
			"link": "/logout"})
	} else {
		c.JSON(http.StatusOK, gin.H{"label": "Login",
			"link": "/login"})
	}
}

// Handler for shortening a URL
func Shorten(c *gin.Context) {
	db := helpers.GetDatabase() // Get the database
	defer db.Close()

	// Stores the JSON sent from the front-end
	var reqData struct {
		URL string `json:"url"`
	}

	// Get the JSON from the front end
	err := c.BindJSON(&reqData)
	if err != nil {
		log.Fatal(err)
	}

	// Generate a random number to use as the shortened URL
	url := rand.Intn(10000000000)

	id := helpers.GetID(c) // Get the user's ID

	// Add the original URL and shortened URL into the database
	_, err = db.Exec("INSERT INTO urls (name, original, short) VALUES ('" + id + "', '" + reqData.URL + "', '" + strconv.Itoa(url) + "');")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Added url to database")
	}

	// Return the shortened URL in a JSON
	c.JSON(http.StatusOK, gin.H{"url": url})
}

// Handler for getting the user's shortened URLS
func GetURLs(c *gin.Context) {
	db := helpers.GetDatabase() // Get the database
	defer db.Close()

	id := helpers.GetID(c) // Get the user's ID

	if id == "default" {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not signed in."})
	} else {
		var rows [][2]string // Array to store original and shortened URLs

		// Query the database to get the rows of original and
		// shortened URLs that match the user
		result, err := db.Query("SELECT original, short FROM urls WHERE name='" + id + "';")
		if err != nil {
			log.Fatal(err)
		}
		defer result.Close()

		// Loop over each row i.e. original and shortened URL pair
		for result.Next() {
			var original string // To store the original URL
			var short string    // To store the shortened URL

			// Put the original and shortened URLs into the variables
			err := result.Scan(&original, &short)
			if err != nil {
				log.Fatal(err)
			} else {
				// Put the variables into an array
				pair := [2]string{original, short}

				// Append this array to the array of rows
				rows = append(rows, pair)
			}
		}

		// Return the array of rows in a JSON
		c.JSON(http.StatusOK, rows)
	}
}

// Handler for deleting a row selected by the user
func DeleteRow(c *gin.Context) {
	db := helpers.GetDatabase() // Get the database
	defer db.Close()

	// Stores the JSON sent from the front-end
	var reqData struct {
		SHORTENED string `json:"shortened"`
	}

	// Get the JSON from the front-end
	err := c.BindJSON(&reqData)
	if err != nil {
		log.Fatal(err)
	}

	shortened := reqData.SHORTENED // Get the shortened URL from the JSON

	// Execute the SQL command to remove the row with this original URL
	_, err = db.Exec("DELETE FROM urls WHERE short = '" + shortened + "';")
	if err != nil {
		log.Fatal(err)
	}
}
