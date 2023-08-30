package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"url-shortener/cmd/authenticator"
	"url-shortener/cmd/callback"
	"url-shortener/cmd/login"
	"url-shortener/cmd/logout"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	db := connectToDatabase() // Get the database
	defer db.Close()          // Defer the closing of the database until the program ends

	router := gin.Default() // Get a router for the web server

	gob.Register(map[string]interface{}{})
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load environment variables")
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatal("Unable to initialise the authenticator")
	}

	// Default route
	router.GET("/", func(c *gin.Context) {
		// Display the Angular frontend
		c.File("web/dist/web/index.html")
	})

	// Static files needed by Angular frontend
	router.GET("/:file", func(c *gin.Context) {
		// Return the static file from the Angular frontend folder
		c.File("web/dist/web/" + c.Param("file"))
	})

	// Redirect a shortened link to the original URL
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

		// Redirect the user to the original URL returned from the database
		c.Redirect(http.StatusPermanentRedirect, "http://"+original)
	})

	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/logout", logout.Handler)
	router.GET("/profile", func(c *gin.Context) {
		c.File("web/dist/web/index.html")
	})

	// Get the URL to shorten from the frontend and send back
	// the generated number
	router.POST("/shorten", func(c *gin.Context) {
		// Stores the JSON sent from the frontend
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

		id := getID(c)

		// Add this URL to the database with the generated number
		_, err := db.Exec("INSERT INTO urls (name, original, short) VALUES ('" + id + "', '" + url + "', '" + strconv.Itoa(short) + "');")
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Added url to database")
		}

		// Send back the generated number as a response
		c.JSON(http.StatusOK, gin.H{"short": short})
	})

	router.POST("/get-profile", func(c *gin.Context) {
		id := getID(c)

		fmt.Println(id)
		if id == "default" {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not signed in."})
		} else {
			var arr [][2]string

			rows, err := db.Query("SELECT original, short FROM urls WHERE name='" + id + "';")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			for rows.Next() {
				var original string
				var short string

				if err := rows.Scan(&original, &short); err != nil {
					log.Fatal(err)
				} else {
					pair := [2]string{original, short}
					arr = append(arr, pair)
				}
			}

			c.JSON(http.StatusOK, arr)
		}
	})

	router.POST("/delete-row", func(c *gin.Context) {
		var reqData struct {
			ROW string `json:"row"`
		}

		if err := c.BindJSON(&reqData); err != nil {
			log.Fatal(err)
		}

		original := reqData.ROW

		_, err := db.Exec("DELETE FROM urls WHERE short = '" + original + "';")
		if err != nil {
			log.Fatal(err)
		}
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

	return db // Return a pointer to the database
}

func getID(c *gin.Context) string {
	session := sessions.Default(c)
	profile := session.Get("profile")
	if profile != nil {
		profileMap := profile.(map[string]interface{})
		return profileMap["sub"].(string)
	} else {
		return "default"
	}
}
