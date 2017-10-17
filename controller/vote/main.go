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
var items map[string]int

const vid = 1

func init() {
	users1 = []string{
		"周从军",
		"于鹏",
		"赵一瀚",
		"蒋红",
	}
	users2 = []string{
		"李广亭",
		"高阳",
		"莫康波",
		"蒋廉",
	}
	items = map[string]int{"陈词":10,"质询":30,"小结":10,"自由辩论":30,"总结陈词":20}
	db = getDb()
}

func Teams(c *gin.Context) {
	// TODO:  need get data from database
	teams := map[string][]string{
		"生活的暴击值得感激":  users1,
		"生活的暴击不值得感激": users2,
	}
	c.JSON(http.StatusOK, gin.H{"teams": teams, "items":items})
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "vote.html", gin.H{
		"title": "投票系统",
	})
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": 200,"success":true})
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
	cids := c.PostFormArray("users")
	score := 1
	db := getDb();
	stmt, err := db.Prepare(`INSERT vote (uid,cid,vid,score,type) values (?,?,?,?,?)`)
	checkErr(err)
	for cid := range cids {
		_, err = stmt.Exec(uid, cid, vid, score, 100)
		checkErr(err)
	}
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

	// 个人得分
	rows, err := db.Query("SELECT cid, sum(score) as score FROM vote where vid=" + strconv.Itoa(vid) + " and cid not in ('正方','反方') group by cid")
	userResult := map[string]int{}
	for rows.Next() {
		var score int
		var cid string
		rows.Columns()
		err = rows.Scan(&cid, &score)
		checkErr(err)
		userResult[cid] = score
	}
	rows.Close()

	// 团队得分
	rows, err = db.Query("SELECT count(id) as total FROM vote where vid=" + strconv.Itoa(vid) + " and cid  in ('正方','反方')")
	var total int
	for rows.Next() {
		rows.Columns()
		err = rows.Scan(&total)
		checkErr(err)
	}
	rows.Close()
	rows, err = db.Query("SELECT cid, sum(score) as score FROM vote where vid=" + strconv.Itoa(vid) + " and cid in ('正方','反方') group by cid")
	teamResult := map[string]int{}
	var score int
	var cid string
	for rows.Next() {
		rows.Columns()
		err = rows.Scan(&cid, &score)
		checkErr(err)
		teamResult[cid] = score*5/total
	}
	rows.Close()
	c.JSON(http.StatusOK, gin.H{"userResult": userResult,"teamResult":teamResult})
}

func getDb() *sql.DB {
	if db == nil {
		var err error
		db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ngdule?charset=utf8")
		checkErr(err)
		db.SetMaxIdleConns(2)
		db.SetMaxOpenConns(2)

		if err := db.Ping(); err != nil {
			log.Fatalln(err)
		}
	}
	return db
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}