package view

import (
	"Coot/view/dashboard"
	"Coot/view/plug"
	"Coot/view/setting"
	"Coot/view/task"
	"github.com/gin-gonic/gin"
)

func LoadUrl(r *gin.Engine) {
	//r.GET("/login", login.Html)
	//r.POST("/login", login.Login)
	// 仪表盘
	r.GET("/", dashboard.Html)
	r.GET("/dashboard", dashboard.Html)

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
	r.GET("/setting", setting.Html)
	r.GET("/get/setting/info", setting.GetSettingInfo)
	r.POST("/post/setting/update", setting.UpdateEmailInfo)
	r.POST("/post/setting/login", setting.UpdateLoginInfo)
	r.POST("/post/setting/checkSetting", setting.UpdateStatusSetting)
}
