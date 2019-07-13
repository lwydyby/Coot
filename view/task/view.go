package task

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Coot/core/job"
	"Coot/core/dbUtil"
	"Coot/error"
	"Coot/utils/file"
	"Coot/utils/md5"
	"time"
	"github.com/satori/go.uuid"
)

// Task List 页面
func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "task.html", gin.H{})
}

// 查询任务列表
func GetTaskList(c *gin.Context) {
	sql := `select id,task_name,task_explain,task_id,task_time_type,task_time,last_exec_time,script_type,script_path,create_time from coot_tasks ORDER BY id desc;`
	result := dbUtil.Query(sql)
	c.JSON(http.StatusOK, error.ErrSuccess(result))
}

// Task Add 页面
func HtmlAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "taskAdd.html", gin.H{})
}

// 添加任务
func PostTaskAdd(c *gin.Context) {
	is_plug_script := c.PostForm("is_plug_script")
	taskName := c.PostForm("taskName")
	taskExplain := c.PostForm("taskExplain")
	taskTimeType := c.PostForm("taskTimeType")
	taskTime := c.PostForm("taskTime")
	taskLanuage := c.PostForm("taskLanuage")
	code := c.PostForm("code")

	// 获取时间戳，生成MD5
	currTimeStr := time.Now().Format("2006-01-02 15:04")
	uid := uuid.NewV4()
	fileName := md5.Md5(currTimeStr + uid.String())

	var fileType = ""

	if taskLanuage == "Python" {
		fileType = "py"
	} else if taskLanuage == "Shell" {
		fileType = "sh"
	} else {
		c.JSON(http.StatusOK, error.ErrFailFileType())
	}

	// 写入文件
	filePath := "./scripts/" + fileName + "." + fileType
	file.Output(code, filePath)

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

	dbUtil.Insert(sql, taskName, taskExplain, "", taskTimeType, taskTime, "", is_plug_script, taskLanuage, filePath, "1", currTimeStr)

	c.JSON(http.StatusOK, error.ErrSuccessNull())
}

// 删除任务
func PostTaskDel(c *gin.Context) {
	id := c.PostForm("id")

	sql := `select id,task_name,task_explain,task_id,task_time_type,task_time,last_exec_time,script_type,script_path,create_time from coot_tasks WHERE id = ?;`
	result := dbUtil.Query(sql, id)

	taskId := result[0]["task_id"]

	job.StopJob(taskId.(string))

	sqlDel := `delete from coot_tasks where id=?;`
	dbUtil.Delete(sqlDel, id)

	c.JSON(http.StatusOK, error.ErrSuccessNull())
}

func updateTaskId(task_id string, id string) {
	sql := `
		UPDATE coot_tasks
		SET task_id = ?
		WHERE
			id = ?;
		`
	dbUtil.Update(sql, task_id, id)
}

func UpdateTaskAll() {
	sql := `
		UPDATE coot_tasks
		SET task_id = "";
		`
	dbUtil.Update(sql)
}

// 启动任务
func TaskStart(c *gin.Context) {
	id := c.PostForm("id")

	sql := `select id,task_name,task_explain,task_id,task_time_type,task_time,last_exec_time,script_type,script_path,create_time from coot_tasks WHERE id = ?;`
	result := dbUtil.Query(sql, id)

	taskTimeType := result[0]["task_time_type"]
	taskTime := result[0]["task_time"]
	scriptType := result[0]["script_type"]
	scriptPath := result[0]["script_path"]

	// 启动任务
	taskId := job.AddJob(&job.Task{
		id,
		"",
		taskTimeType.(string),
		taskTime.(string),
		scriptType.(string),
		scriptPath.(string),
	})

	// 更新数据库
	updateTaskId(taskId, id)

	c.JSON(http.StatusOK, error.ErrSuccessNull())
}

// 关闭任务
func TaskStop(c *gin.Context) {
	id := c.PostForm("id")

	sql := `select id,task_name,task_explain,task_id,task_time_type,task_time,last_exec_time,script_type,script_path,create_time from coot_tasks WHERE id = ?;`
	result := dbUtil.Query(sql, id)

	taskId := result[0]["task_id"]

	// 停止任务
	job.StopJob(taskId.(string))

	// 更新数据库
	updateTaskId("", id)

	c.JSON(http.StatusOK, error.ErrSuccessNull())
}
