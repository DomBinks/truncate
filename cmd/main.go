package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	db := connectToDatabase()

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "/")
		c.Header("Access-Control-Allow-Methods", "POST")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	router.GET("/", func(c *gin.Context) {
		c.File("web/dist/web/index.html")
	})

	router.GET("/:file", func(c *gin.Context) {
		c.File("web/dist/web/" + c.Param("file"))
	})

	router.GET("/link/:short", func(c *gin.Context) {
		short := c.Param("short")
		var original string

		fmt.Println("SELECT original FROM urls WHERE short = '" + short + "';")
		err := db.QueryRow("SELECT original FROM urls WHERE short = '" + short + "';").Scan(&original)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("Couldn't find link row")
			} else {
				log.Fatal(err)
			}
		}

		c.Redirect(http.StatusPermanentRedirect, "http://"+original)
	})

	/*
		router.POST("/shorten", func(c *gin.Context) {
			url := c.PostForm("url")
				short := rand.Intn(10000000000)

				_, err = db.Exec("INSERT INTO urls (name, original, short) VALUES ('test', '" + url + "', '" + strconv.Itoa(short) + "');")
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println("Added url to database")
				}

				c.HTML(200, "shortened.html", gin.H{
					"short": short,
				})

			fmt.Println("aaa " + url)
		})
	*/

	router.POST("/shorten", func(c *gin.Context) {
		fmt.Println("IN GO")
		var reqData struct {
			URL string `json:"url"`
		}

		if err := c.BindJSON(&reqData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
			return
		}

		fmt.Println("out " + reqData.URL)
	})

	router.Run("localhost:8080")
}

func connectToDatabase() *sql.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "urlshortener"
		password = "golang"
		dbname   = "urlshortener"
	)

	connectionString := "host=" + host + " port =" +
		strconv.Itoa(port) + " user=" + user + " password=" +
		password + " dbname=" + dbname + " sslmode=disable"

	fmt.Println("Connecting to database")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Getting rows")
	rows, err := db.Query("SELECT name, original, short FROM urls;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Rows:")
	for rows.Next() {
		var name string
		var original string
		var short string
		if err := rows.Scan(&name, &original, &short); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(name, original, short)
		}
	}

	return db
}
