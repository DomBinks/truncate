package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello, World!")

	const (
		host     = "localhost"
		port     = 5432
		user     = "urlshortener"
		password = "golang"
		dbname   = "urls"
	)

	connectionString := "host=" + host + " port =" +
		strconv.Itoa(port) + " user=" + user + " password=" +
		password + " dbname=" + dbname + " sslmod=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT new, original, name FROM urls")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var from string
		var to string
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name, from, to)
	}

	router := gin.Default()

	router.LoadHTMLGlob("web/templates/*")

	router.Static("/", "./web/static/")

	router.NoRoute(func(c *gin.Context) {
		c.File("./web/static/index.html")
	})

	router.POST("/shorten", func(c *gin.Context) {
		url := c.PostForm("url")
		c.HTML(200, "response.html", gin.H{
			"url": url,
		})
	})

	router.Run("localhost:8080")
}
