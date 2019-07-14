package main

import (
	"Coot/view"
	"Coot/view/task"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create("./logs/coot.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 引入gin
	r := gin.Default()

	// 引入html资源
	r.LoadHTMLGlob("web/*")

	// 引入静态资源
	r.Static("/static", "./static")

	// 加载路由
	view.LoadUrl(r)
	task.UpdateTaskAll()
	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run()
}
