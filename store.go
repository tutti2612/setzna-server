package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var connectionError error

// DB の接続情報
const (
	DRIVER_NAME = "mysql"
	// user:password@tcp(container-name:port)/dbname ※mysql はデフォルトで用意されているDB
	DATA_SOURCE_NAME = "root:golang@tcp(mysql-container:3306)/mysql"
)

// 初期化処理
func init() {
	db, connectionError = sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	if connectionError != nil {
		log.Fatal("error connecting to database: ", connectionError)
	}
	println(db)
}

func getDBInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get DB Info")
	// DB の情報を取得
	rows, err := db.Query("SELECT SUBSTRING_INDEX(USER(), '@', -1) AS ip, @@hostname as hostname, @@port as port, DATABASE() as current_dtabase;")
	if err != nil {
		log.Print("error executing database query: ", err)
		return
	}
	var buffer bytes.Buffer
	for rows.Next() {
		var ip string
		var hostname string
		var port string
		var current_database string
		err = rows.Scan(&ip, &hostname, &port, &current_database)
		buffer.WriteString("IP::" + ip + "| HostName::" + hostname + "|Port::" + port + "|CurrentDatabase::" + current_database)
	}
	fmt.Fprint(w, buffer.String())
}

func main() {
	// gorilla mux でルーティングする
	// router := mux.NewRouter()
	defer db.Close()
}
