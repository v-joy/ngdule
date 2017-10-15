package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var DB = make(map[string]string)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?parseTime=true")
	defer db.Close()
	if err != nil{
		log.Fatalln(err)
	}


	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil{
		log.Fatalln(err)
	}

	vote := "/vote";

	app := gin.Default()
	app.Use(favicon.New("./favicon.ico"))

	app.Static("/assets", "./assets")

	app.GET(vote + "/", func(c *gin.Context) {
		c.String(200, "Hello favicon.")
	})
	app.LoadHTMLGlob("templates/*")
	app.GET(vote + "/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "vote.html", gin.H{
			"title": "投票系统",
		})
	})

	app.POST(vote + "/login", func(c *gin.Context) {

	})

	app.POST(vote + "/competitor/score", func(c *gin.Context) {

	})

	app.POST(vote + "/team/score", func(c *gin.Context) {

	})

	app.GET(vote + "/competitors", func(c *gin.Context) {

	})

	app.GET(vote + "/judgers", func(c *gin.Context) {

	})

	app.Run(":8080")
}