package main

import (
	"github.com/bai615/Go-000/Week04/goblog/app/http/middlewares"
	"github.com/bai615/Go-000/Week04/goblog/bootstrap"
	"github.com/bai615/Go-000/Week04/goblog/config"
	c "github.com/bai615/Go-000/Week04/goblog/pkg/config"
	"net/http"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	// 初始化 SQL
	bootstrap.SetupDB()

	// 初始化路由绑定
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
