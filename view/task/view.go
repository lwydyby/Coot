package task

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Coot/core/job"
	"fmt"
	"Coot/core/dbUtil"
	"Coot/error"
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
		0,
		"",
		1,
		"3",
	})

	fmt.Println(taskId)

	c.JSON(http.StatusOK, data)
}

// 关闭任务
func Stop(c *gin.Context) {
	data := map[string]interface{}{
		"k": "测试",
	}

	taskId := c.Query("taskId")
	job.StopJob(taskId)

	c.JSON(http.StatusOK, data)
}

// 插入数据
func Insert(c *gin.Context) {
	data := map[string]interface{}{
		"k": "测试",
	}

	sql := `
		INSERT INTO coot_tasks (
			task_name,
			task_explain,
			task_id,
			task_time_type,
			task_time,
			last_exec_time,
			is_plug_script,
			script_type,
			script_path,
			alert_type,
			create_time
		)
		VALUES
			(?,?,?,?,?,?,?,?,?,?,?);
	`
	dbUtil.Insert(sql, "插入任务测试", "测试说明", "", 1, "2", "", "1", "shell", "/plugs/myscript/test.sh", "1", "2019-07-10 16:12")

	c.JSON(http.StatusOK, data)
}

// 更新数据
func Update(c *gin.Context) {
	data := map[string]interface{}{
		"k": "测试",
	}

	sql := `
		UPDATE coot_tasks
		SET task_name = ?
		WHERE
			id = ?;
	`
	dbUtil.Update(sql, "任务更新测试", 1)

	c.JSON(http.StatusOK, data)
}

// 删除数据
func Delete(c *gin.Context) {
	data := map[string]interface{}{
		"k": "测试",
	}

	sql := `delete from coot_tasks where id=?;`
	dbUtil.Delete(sql, 2)

	c.JSON(http.StatusOK, data)
}

// 查询数据
func Query(c *gin.Context) {
	sql := `select * from coot_tasks;`
	result := dbUtil.Query(sql)
	//sql := `select * from coot_tasks where id=?;`
	//result := dbUtil.Query(sql, 1)

	c.JSON(http.StatusOK, error.ErrSuccess(result))
}
