package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	melody "gopkg.in/olahol/melody.v1"
)

// 2点間の距離を計算して半径50m以内だったらtrue
func isClose(lat1, lng1, lat2, lng2 interface{}) bool {
	const setznaDistance float64 = 50
	// TODO: 三平方の定理などで２点間の距離を出す
	// 緯度1度あたり110.9463km、経度1度あたり90.4219km
	// https://s-giken.info/distance/distance.php

	// HACK: 地球の丸みを考慮すると、三平方の定理では不十分なので改善する。
	// distance := math.Sqrt(math.Pow((lat1-lat2)*110.9463, 2) + math.Pow((lng1-lng2)*90.4219, 2))
	// return distance < setznaDistance
	return true
}

func main() {
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
		// TODO: コネクション接続時に初期位置情報をセットする
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		// TODO: メッセージの度に位置情報を更新する

		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			lat1, isExistLat1 := s.Get("latitude")
			lng1, isExistLng1 := s.Get("longitude")
			lat2, isExistLat2 := q.Get("latitude")
			lng2, isExistLng2 := q.Get("longitude")
			if !isExistLat1 || !isExistLng1 || !isExistLat2 || !isExistLng2 {
				return false
			}

			return isClose(lat1, lng1, lat2, lng2)
		})
	})
	r.Run(":8080")
}
