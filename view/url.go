package view

import (
	"github.com/gin-gonic/gin"
	"Coot/view/task"
	"Coot/view/plug"
)

func LoadUrl(r *gin.Engine) {
	// 仪表盘
	r.GET("/", task.Html)
	//r.GET("/dashboard", dashboard.Html)
	//r.GET("/dashboard/get/data", dashboard.Get)

	// 任务
	r.GET("/task", task.Html)
	r.GET("/task/add", task.HtmlAdd)
	r.GET("/get/task/list", task.GetTaskList)
	r.POST("/post/task/add", task.PostTaskAdd)
	r.POST("/post/task/del", task.PostTaskDel)
	r.POST("/task/start", task.TaskStart)
	r.POST("/task/stop", task.TaskStop)

	// 插件
	r.GET("/plugs", plug.Html)

	// 设置
	//r.GET("/setting", setting.Html)
}
