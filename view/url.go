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
	r.GET("/dashboard/get/data", dashboard.Get)

	// 任务
	r.GET("/task", task.Html)
	r.GET("/task/add", task.HtmlAdd)
	r.GET("/get/task/list", task.GetTaskList)

	// 定时任务测试 + 执行脚本测试
	r.GET("/task/start", task.Start)
	r.GET("/task/stop", task.Stop)
	// 数据库测试
	r.GET("/task/insert", task.Insert)
	r.GET("/task/update", task.Update)
	r.GET("/task/del", task.Delete)
	r.GET("/task/query", task.Query)

	// 插件
	r.GET("/plugs", plug.Html)

	// 设置
	r.GET("/setting", setting.Html)
}
