package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	melody "gopkg.in/olahol/melody.v1"
)

type post struct {
	PostType  string `json:"type"`
	Name      string `json:"name"`
	Message   string `json:"message"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// 2点間の距離を計算して半径50m以内だったらtrue
func isNear(lat1, lng1, lat2, lng2 float64) bool {
	const setznaDistance float64 = 0.05 // 0.05km = 50m
	// 緯度1度あたり110.9463km、経度1度あたり90.4219km
	// https://s-giken.info/distance/distance.php
	// HACK: 地球の丸みを考慮すると、三平方の定理では不十分なので改善する。
	distance := math.Sqrt(math.Pow((lat1-lat2)*110.9463, 2) + math.Pow((lng1-lng2)*90.4219, 2))
	return distance < setznaDistance
}

func isSessionNear(s, q *melody.Session) bool {
	lat1, isExistLat1 := s.Get("latitude")
	lng1, isExistLng1 := s.Get("longitude")
	lat2, isExistLat2 := q.Get("latitude")
	lng2, isExistLng2 := q.Get("longitude")
	if !isExistLat1 || !isExistLng1 || !isExistLat2 || !isExistLng2 {
		return false
	}

	lat1F64, _ := strconv.ParseFloat(lat1.(string), 64)
	lng1F64, _ := strconv.ParseFloat(lng1.(string), 64)
	lat2F64, _ := strconv.ParseFloat(lat2.(string), 64)
	lng2F64, _ := strconv.ParseFloat(lng2.(string), 64)

	return isNear(lat1F64, lng1F64, lat2F64, lng2F64)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	m := melody.New()

	// フロントが完成したら不要
	r.Static("/static", "./view/static")
	r.LoadHTMLGlob("view/*.html")

	// フロントが完成したら不要
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// フロントが完成したら不要
	r.GET("/room", func(c *gin.Context) {
		c.HTML(http.StatusOK, "room.html", gin.H{})
	})

	r.GET("/room/ws", func(c *gin.Context) {
		// コネクション接続時に初期位置情報をセットする
		m.HandleRequestWithKeys(
			c.Writer,
			c.Request,
			map[string]interface{}{"latitude": c.Query("latitude"), "longitude": c.Query("longitude")},
		)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		var p post
		json.Unmarshal(msg, &p)
		if p.PostType == "location" {
			// 緯度、経度をセットする
			s.Set("latitude", p.Latitude)
			s.Set("longitude", p.Longitude)
		}
		if p.PostType == "message" {
			m.BroadcastFilter(msg, func(q *melody.Session) bool {
				return isSessionNear(s, q)
			})

			// ここらへんでフロント側から受け取った名前、メッセージなどなどを保存する。
			go saveDB(p)
		}
	})
	r.Run(":" + port)
}

func saveDB(p post) bool {
	const (
		DRIVER_NAME = "mysql"
		// user:password@tcp(container-name:port)/dbname ※mysql はデフォルトで用意されているDB
		DATA_SOURCE_NAME = "setzna:setzna@tcp(setzna_mysql:3306)/mysql"
	)

	// DB の接続情報
	var db *sql.DB
	var err error

	db, err = sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	// データ保存処理
	stmtIns, err := db.Prepare(fmt.Sprintf("INSERT INTO messages (message, latitude, longitude) VALUES (?, ?, ?)"))
	if err != nil {
		panic(err.Error())
	}
	_, err = stmtIns.Exec(p.Content, p.Latitude, p.Longitude)
	defer db.Close()

	if err != nil {
		log.Print("error connecting to database: ", err)
		return false
	}

	return true
}

// messagesテーブル
// id
// message
// latitude
// longitude
// created_at
