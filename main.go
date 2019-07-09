package main

import (
	"github.com/gin-gonic/gin"
	"Coot/view"
)

func main() {
	// 引入gin
	r := gin.Default()

	// 引入html资源
	r.LoadHTMLGlob("web/*")

	// 引入静态资源
	r.Static("/static", "./static")

	// 加载路由
	view.LoadUrl(r)

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run()
}
