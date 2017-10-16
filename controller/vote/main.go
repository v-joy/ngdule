package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

var db *sql.DB
var users1 []string
var users2 []string

const vid = 1

func init() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ngdule?charset=utf8")
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

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(2)

	if err := db.Ping(); err != nil {
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

func IsVoted(c *gin.Context) {
	uid := c.Query("username"); // Todo 需要改为从session 获取
	cid := c.Param("cid")
	// Todo: 为什么在此处打印db 为nil?
	db := getDb();
	rows, err := db.Query("SELECT COUNT(*) as nums FROM vote where vid=" + strconv.Itoa(vid) + " and cid='" + cid + "' and uid='" + uid + "'")
	checkErr(err)
	defer rows.Close()
	voted := true;
	for rows.Next() {
		var nums int
		rows.Columns()
		err = rows.Scan(&nums)
		checkErr(err)
		if nums == 0 {
			voted = false
		}
		break
	}
	c.JSON(http.StatusOK, gin.H{"voted": voted})
}

func SetCompetitorScore(c *gin.Context) {
	uid := c.Query("username"); // Todo 需要改为从session 获取
	cid := c.Param("cid")
	score := 1
	db := getDb();
	stmt, err := db.Prepare(`INSERT vote (uid,cid,vid,score,type) values (?,?,?,?,?)`)
	checkErr(err)
	_, err = stmt.Exec(uid, cid, vid, score, 100)
	checkErr(err)
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func SetTeemScore(c *gin.Context) {
	uid := c.Query("username"); // Todo 需要改为从session 获取
	cid := c.Param("cid")
	scores := c.PostFormArray("scores")
	db := getDb();
	stmt, err := db.Prepare(`INSERT vote (uid,cid,vid,score,type) values (?,?,?,?,?)`)
	checkErr(err)
	for i, score := range scores {
		_, err = stmt.Exec(uid, cid, vid, score, i)
		checkErr(err)
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func Summery(c *gin.Context) {
	db := getDb();
	rows, err := db.Query("SELECT cid, sum(score) as score FROM vote where vid=" + strconv.Itoa(vid) + " group by cid")
	defer rows.Close()
	result := map[string]string{}
	for rows.Next() {
		var score string
		var uid string
		rows.Columns()
		err = rows.Scan(&uid, &score)
		checkErr(err)
		result[uid] = score
		break
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func getDb() *sql.DB {
	if db == nil {
		db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ngdule?charset=utf8")
		log.Println("reopen db")
	}
	return db
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}