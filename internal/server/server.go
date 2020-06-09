package server

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	melody "gopkg.in/olahol/melody.v1"

	"setzna/internal/model"
	"setzna/pkg/location"
)

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	m := melody.New()

	// フロントが完成したら不要
	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/template/*.html")

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
		var message model.Message
		json.Unmarshal(msg, &message)
		if message.Type == "location" {
			// 緯度、経度をセットする
			s.Set("latitude", message.Latitude)
			s.Set("longitude", message.Longitude)
		}
		if message.Type == "post" {
			m.BroadcastFilter(msg, func(q *melody.Session) bool {
				return isSessionNear(s, q)
			})

			// ここらへんでフロント側から受け取った名前、メッセージなどなどを保存する。
			go message.Save()
		}
	})
	r.Run(":" + port)
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

	const setznaDistance float64 = 0.05 // 0.05km = 50m
	return location.Distance(lat1F64, lng1F64, lat2F64, lng2F64) < setznaDistance
}
