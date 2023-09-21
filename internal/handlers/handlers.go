package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"truncate/internal/helpers"

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

	shortened := "http://" + os.Getenv("HANDLER_IP") + ":8080/~" + c.Param("shortened") // Get the shortened URL
	log.Println("shortened")

	var original string // Original URL returned from the database

	// Get the original URL from the database, and store in the
	// original string
	err := db.QueryRow("SELECT original FROM urls WHERE shortened = $1;", shortened).Scan(&original)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Couldn't find the URL corresponding to the link")
		} else {
			log.Fatal(err)
		}
	}

	// Redirect the user to the original URL returned from the database
	c.Redirect(http.StatusPermanentRedirect, original)
}

// Handler for getting the correct label and link for the login/logout
// button
func GetLogin(c *gin.Context) {
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

	// Stores the JSON sent from the frontend
	var reqData struct {
		URL string `json:"url"`
	}

	// Get the JSON from the front end
	err := c.BindJSON(&reqData)
	if err != nil {
		log.Fatal(err)
	}

	// Get the original URL from the JSON
	original := reqData.URL

	// Add "https://" to the original URL if it isn't there already
	if !strings.HasPrefix(original, "http://") && !strings.HasPrefix(original, "https://") {
		original = "https://" + original
	}

	// Parse the original URL to check it's valid
	parsedURL, err := url.Parse(original)

	// If the original URL isn't a valid URL
	if err != nil || parsedURL.Host == "" || !strings.Contains(parsedURL.Host, ".") {
		// Return an error to the frontend
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	// Generate a random string to use as the unique part of the shortened URL
	shortened := "http://" + os.Getenv("HANDLER_IP") + ":8080/~" + helpers.GenerateShortened()

	log.Println(original + " shortened to " + shortened)

	id := helpers.GetID(c) // Get the user's ID

	// Add the original URL and shortened URL into the database
	_, err = db.Exec("INSERT INTO urls (id, original, shortened) VALUES ($1, $2, $3);", id, original, shortened)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Added URL to database")
	}

	// Return the shortened URL in a JSON
	c.JSON(http.StatusOK, gin.H{"shortened": shortened})
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
		result, err := db.Query("SELECT original, shortened FROM urls WHERE id=$1;", id)
		if err != nil {
			log.Fatal(err)
		}
		defer result.Close()

		// Loop over each row i.e. original and shortened URL pair
		for result.Next() {
			var original string  // To store the original URL
			var shortened string // To store the shortened URL

			// Put the original and shortened URLs into the variables
			err := result.Scan(&original, &shortened)
			if err != nil {
				log.Fatal(err)
			} else {
				// Put the variables into an array
				pair := [2]string{original, shortened}

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
	_, err = db.Exec("DELETE FROM urls WHERE shortened =$1;", shortened)
	if err != nil {
		log.Fatal(err)
	}
}
