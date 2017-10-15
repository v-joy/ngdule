package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"strconv"
)

var Db *sql.DB
var users1 []string
var users2 []string

const vid = 1

func init() {
	Db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ngdule?charset=utf8")
	users1 = []string{
		"chenxiao",
		"liujunling03",
		"test",
		"测试",
	}
	users2 = []string{
		"chenxiao2",
		"liujunling032",
		"test2",
		"测试2",
	}
	if err != nil {
		log.Fatalln(err)
	}

	Db.SetMaxIdleConns(2)
	Db.SetMaxOpenConns(2)

	if err := Db.Ping(); err != nil {
		log.Fatalln(err)
	}
}

func Teams(c *gin.Context) {
	// TODO:  need get data from database
	teams := map[string][]string{
		"应该玩游戏":  users1,
		"不应该玩游戏": users2,
	}
	c.JSON(http.StatusOK, gin.H{"teams": teams})
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "vote.html", gin.H{
		"title": "投票系统",
	})
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func getDb() *sql.DB {
	Db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ngdule?charset=utf8")
	return Db
}

func IsVoted(c *gin.Context) {
	cid := c.Param("cid")
	// Todo: 为什么在此处打印db 为nil?
	Db := getDb();
	if err := Db.Ping(); err != nil {
		log.Fatalln(err)
	}
	log.Println("SELECT COUNT(*) as nums FROM vote where vid=" + strconv.Itoa(vid) +" and cid=" + cid)
	rows, err := Db.Query("SELECT COUNT(*) as nums FROM vote where vid=" + strconv.Itoa(vid) +" and cid=" + cid)
	defer rows.Close()
	voted := true;
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusOK, gin.H{"voted": voted})
		return ;
	}
	for rows.Next() {
		var nums int
		rows.Columns()
		err = rows.Scan(&nums)
		_ = err
		fmt.Println(nums) // 为什么在此处访问 Db为空？
		if nums == 0 {
			voted = false
		}
		break
	}
	c.JSON(http.StatusOK, gin.H{"voted": voted})
}

func SetCompetitorScore(c *gin.Context) {

}

func SetTeemScore(c *gin.Context) {

}

func Summery(c *gin.Context) {

}
