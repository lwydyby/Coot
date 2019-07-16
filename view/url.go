package view

import (
	"Coot/error"
	"Coot/view/dashboard"
	"Coot/view/login"
	"Coot/view/plug"
	"Coot/view/setting"
	"Coot/view/task"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/favicon.ico" {
			c.Next()
			return
		}
		tokenStr, err := c.Cookie("userToken")
		fmt.Println(tokenStr, err)
		if err != nil {
			c.JSON(http.StatusUnauthorized, error.ErrSuccessCustom(10002, "success", nil))
			c.Abort()
			return
		}
		c.Next()
	}
}
func LoadUrl(r *gin.Engine) {
	r.GET("/login", login.Html)
	r.POST("/login", login.Login)
	//r.Use(middleware())
	// 仪表盘
	r.GET("/", dashboard.Html)
	r.GET("/dashboard", dashboard.Html)
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
	r.GET("/setting", setting.Html)
	r.GET("/get/setting/info", setting.GetSettingInfo)
	r.POST("/post/setting/update", setting.UpdateEmailInfo)
	r.POST("/post/setting/login", setting.UpdateLoginInfo)
	r.POST("/post/setting/checkSetting", setting.UpdateStatusSetting)
}
