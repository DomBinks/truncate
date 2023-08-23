package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	db := connectToDatabase() // Get the database
	defer db.Close()

	router := gin.Default() // Get a router for the web server

	// Default route
	router.GET("/", func(c *gin.Context) {
		// Display the Angular frontend
		c.File("web/dist/web/index.html")
	})

	// Static files needed by Angular frontend
	router.GET("/:file", func(c *gin.Context) {
		// Return the static file from the Angular frontend
		c.File("web/dist/web/" + c.Param("file"))
	})

	// Shortened link
	router.GET("/link/:short", func(c *gin.Context) {
		short := c.Param("short") // Get the URL number
		var original string       // URL returned from the database

		// Get the URL from the database
		err := db.QueryRow("SELECT original FROM urls WHERE short = '" + short + "';").Scan(&original)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("Couldn't find the URL corresponding to the link")
			} else {
				log.Fatal(err)
			}
		}

		// Redirect the user to the URL returned from the database
		c.Redirect(http.StatusPermanentRedirect, "http://"+original)
	})

	router.POST("/shorten", func(c *gin.Context) {
		// JSON sent from the frontend
		var reqData struct {
			URL string `json:"url"`
		}

		// Get the JSON from the front end
		if err := c.BindJSON(&reqData); err != nil {
			log.Fatal(err)
		}

		url := reqData.URL // Get the URL from the JSON

		// Generate a random number to use as the shortened link
		short := rand.Intn(10000000000)

		fmt.Println("url: " + url + " short: " + strconv.Itoa(short))

		// Add this URL to the database with the generated number
		_, err := db.Exec("INSERT INTO urls (name, original, short) VALUES ('test', '" + url + "', '" + strconv.Itoa(short) + "');")
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Added url to database")
		}

		c.JSON(http.StatusOK, gin.H{"short": short})
	})

	router.Run("localhost:8080") // Run the web server
}

// Connects to the PostgreSQL database and returns it as an sql.DB
// pointer
func connectToDatabase() *sql.DB {
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

	fmt.Println("Connecting to database")

	// Connect to the database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Getting rows")

	// Get all the rows in the database
	rows, err := db.Query("SELECT name, original, short FROM urls;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Rows:")

	// Print all the rows retrieved from the database
	for rows.Next() {
		// Stores each field of the row
		var name string
		var original string
		var short string

		// Get the fields
		if err := rows.Scan(&name, &original, &short); err != nil {
			log.Fatal(err)
		} else {
			// Print the row
			fmt.Println(name, original, short)
		}
	}

	return db
}
