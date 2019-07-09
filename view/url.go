package view

import (
	"github.com/gin-gonic/gin"
	"Coot/view/dashboard"
	"Coot/view/task"
	"Coot/view/plug"
	"Coot/view/setting"
)

func LoadUrl(r *gin.Engine) {
	// 仪表盘
	r.GET("/", dashboard.Html)
	r.GET("/dashboard", dashboard.Html)
	r.GET("/dashboard/get/data", dashboard.GetData)

	// 任务
	r.GET("/task", task.Html)

	// 插件
	r.GET("/plug", plug.Html)

	// 设置
	r.GET("/setting", setting.Html)
}
