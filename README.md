# Setzna

## usage
```bash
$ make run
```
localhost:8080で動きます。

```bash
$ make debug
```
リモートデバッグしたいときはこちら  
IntelliJ IDEA, Golandで動作確認済み

## library

- [Gin Web Framework](https://github.com/gin-gonic/gin)
    - 軽量Webフレームワーク
    - pythonのflaskみたいなもの
- [melody](https://github.com/olahol/melody)
    - Websocketフレームワーク
    - メッセージのフィルタリングが簡単にできるので採用
- [gorm](https://github.com/jinzhu/gorm)
    - ORM
    - [日本語ドキュメント](https://gorm.io/ja_JP/)あります
- [realize](https://github.com/oxequa/realize)
    - ホットリロードできるようになるやつ
- [delve](https://github.com/go-delve/delve)
    - リモートデバッグに必要

## ディレクトリ構成

ディレクトリ構成は↓を参考にしています。  
https://github.com/golang-standards/project-layout

## テスト環境(Heroku)

Herokuにデプロイしています。自由に使ってください。  
https://setzna.herokuapp.com/

## ローカル開発環境

ローカルで動かすときはWebsocketプロトコルをwssからwsに変更してください。

## Migration
```bash
$ docker-compose exec app go run cmd/migration/migration.go
# or
$ make migration
```
gormのマイグレーション機能を使用

## Test

```bash
$ docker-compose exec app go test -v ./...
# or
$ make test
```

## リクエストjson

テスト環境(Heroku)に投げる場合はwss  
wss://setzna.herokuapp.com/room/ws  

ローカルに投げる場合はws  
ws://localhost:8080/room/ws

### ポストリクエスト

ユーザーが投稿する度にサーバーに送信する想定

```
{
    "type": "post",
    "name": "hoge",
    "content": "fugafugafuga",
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
    "content": "",
    "latitude": 65.123123,
    "longitude": 123.123123
}
```