package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	vote "ngdule/controller/vote"
)

func main() {
	votePrefix := "/vote"

	app := gin.Default()
	app.Use(favicon.New("./favicon.ico"))
	app.Use(gin.Recovery())
	// app.Use(gin.Logger())

	app.Static("/assets", "./assets")
	app.LoadHTMLGlob("templates/*")

	app.GET(votePrefix+"/", vote.Index)
	app.GET(votePrefix+"/summery", vote.Summery)
	app.GET(votePrefix+"/teams", vote.Teams)
	app.GET(votePrefix+"/score/team", vote.GetTeemScore)
	app.GET(votePrefix+"/score/competitor", vote.GetCompetitorScore)

	app.POST(votePrefix+"/login", vote.Login)
	app.POST(votePrefix+"/score/competitor", vote.SetCompetitorScore)
	app.POST(votePrefix+"/score/team", vote.SetTeemScore)

	app.Run("172.24.125.199:8080")
}