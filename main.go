package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	melody "gopkg.in/olahol/melody.v1"
)

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
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			// ここで半径50mのユーザーに絞る
			return true
		})
	})
	r.Run(":8080")
}
