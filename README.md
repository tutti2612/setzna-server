# Setzna

## usage
```
$ go run main.go
```
localhost:8080で動きます。

## library

- Gin Web Framework https://github.com/gin-gonic/gin
- melody https://github.com/olahol/melody

## テスト環境(Heroku)

Herokuにデプロイしています。自由に使ってください。  
https://setzna.herokuapp.com/

## ローカル開発環境

ローカルで動かすときはWebsocketプロトコルをwssからwsに変更してください。

## リクエストjson

テスト環境(Heroku)に投げる場合はwss  
wss://setzna.herokuapp.com/room/ws  

ローカルに投げる場合はws  
ws://localhost:8080/room/ws

### メッセージリクエスト

ユーザーが投稿する度にサーバーに送信する想定

```
{
    "type": "message",
    "name": "hoge",
    "message": "fugafugafuga",
    "latitude": 65.123123,
    "longitude": 123.123123
}
```

### ロケーションリクエスト

ユーザーの位置情報をリアルタイムに知るために一定間隔(5秒とか？)ごとにロケーション情報(緯度、経度)をサーバーに送信する想定。

```
{
    "type": "location",
    "name": "",
    "message": "",
    "latitude": 65.123123,
    "longitude": 123.123123
}
```