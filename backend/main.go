package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_18_practice/handler"
	"github.com/traP-jp/h24s_18_practice/model"
)

func main() {
	user := os.Getenv("MARIADB_USERNAME")
	pass := os.Getenv("MARIADB_PASSWORD")
	host := os.Getenv("MARIADB_HOSTNAME")
	dbname := os.Getenv("MARIADB_DATABASE")

	err := model.Init(user, pass, host, dbname)
	if err != nil {
		panic(err)
	}

	// Echoの新しいインスタンスを作成
	e := echo.New()

	// 「/hello」というエンドポイントを設定する
	e.GET("/hello", handler.GetHello)

	// Webサーバーをポート番号8080で起動し、エラーが発生した場合はログにエラーメッセージを出力する
	e.Logger.Fatal(e.Start(":8080"))
}


func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}