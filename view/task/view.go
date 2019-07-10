package task

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Coot/utils/job"
	"fmt"
)

// 返回 Task 页面
func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "task.html", gin.H{
		"title": "Main website",
	})
}

// 启动任务
func Start(c *gin.Context) {
	data := map[string]interface{}{
		"k": "测试",
	}

	taskId := job.AddJob(&job.Task{
		"",
		"shell",
		"/xd/ad/adas/1.sh",
		1,
		"1",
	})

	fmt.Println(taskId)

	c.JSONP(http.StatusOK, data)
}

// 关闭任务
func Stop(c *gin.Context) {
	data := map[string]interface{}{
		"k": "测试",
	}

	taskId := c.Query("taskId")
	job.StopJob(taskId)

	c.JSONP(http.StatusOK, data)
}
