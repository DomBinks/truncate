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

	router := gin.Default()

	router.LoadHTMLGlob("web/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
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

		/*
			c.HTML(200, "link.html", gin.H{
				"original": original,
			})
		*/

		c.Redirect(http.StatusPermanentRedirect, "http://"+original)
	})

	router.NoRoute(func(c *gin.Context) {
		c.File("./web/static/index.html")
	})

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
	})

	router.Run("localhost:8080")
}
