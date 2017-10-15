package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	 vote "ngdule/controller/vote"
)

var DB = make(map[string]string)

func main() {

	votePrefix := "/vote"

	app := gin.Default()
	app.Use(favicon.New("./favicon.ico"))

	app.Static("/assets", "./assets")
	app.LoadHTMLGlob("templates/*")

	app.GET(votePrefix + "/", vote.Index)
	app.GET(votePrefix + "/summery", vote.Summery)
	app.GET(votePrefix + "/teams", vote.Teams)
	app.GET(votePrefix + "/voted/:cid", vote.IsVoted)

	app.POST(votePrefix + "/login", vote.Login)
	app.POST(votePrefix + "/score/competitor/:cid", vote.SetCompetitorScore)
	app.POST(votePrefix + "/score/team/:cid", vote.SetTeemScore)


	app.Run("192.168.31.248:8080")
}