package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB = make(map[string]string)

func main() {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
	defer db.Close()
	if err != nil{
		log.Fatalln(err)
	}


	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil{
		log.Fatalln("error")
	}

	app := gin.Default()
	app.Use(favicon.New("./favicon.ico"))

	app.GET("/ping", func(c *gin.Context) {
		c.String(200, "Hello favicon.")
	})

	app.POST("/login", func(c *gin.Context) {

	})

	app.POST("/counter/competitor/score", func(c *gin.Context) {

	})

	app.POST("/counter/team/score", func(c *gin.Context) {

	})

	app.GET("/counter/competitors", func(c *gin.Context) {

	})

	app.GET("/counter/judgers", func(c *gin.Context) {

	})

	app.GET("/counter/team/score", func(c *gin.Context) {

	})

	app.GET("/counter/competitor/score", func(c *gin.Context) {

	})


	app.Run(":8080")
}