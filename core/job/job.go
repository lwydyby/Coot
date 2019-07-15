package job

import (
	"Coot/core/dbUtil"
	"Coot/core/exec"
	"Coot/error"
	"fmt"
	"github.com/domgoer/gotask"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"Coot/utils/send"
	"strings"
)

type Task struct {
	/*
	 * Id         	数据库ID
	 * Name		  	任务名称
	 * TaskId     	任务ID  添加的时候为空
	 * TimeType   	执行类型 1 秒执行，2 分钟执行，3 小时执行 ，4 每天指定时间执行，5 每月指定天和时间执行，6 年执行
	 * Time   	  	周期时间
	 * ScriptType 	脚本语言
	 * ScriptPath 	脚本路径
	*/
	Id           string
	Name         string
	TaskId       string
	TimeType     string
	Time         string
	ScriptType   string
	ScriptPath   string
	AlertType    string
	AlertRecMail string
}

func updateExecTime(id string) {
	sql := `
		UPDATE coot_tasks
		SET last_exec_time = ?
		WHERE
			id = ?;
		`

	currTimeStr := time.Now().Format("2006-01-02 15:04:05")
	dbUtil.Update(sql, currTimeStr, id)
}

// 执行任务
func execute(t *Task) {
	var id = t.Id
	var cmd string

	exeResult := "开始执行"

	if t.ScriptType == "Python" {
		cmd = "python " + t.ScriptPath
	} else if t.ScriptType == "Shell" {
		cmd = "sh " + t.ScriptPath
	}

	fmt.Fprintln(gin.DefaultWriter, time.Now().Format("2006-01-02 15:04:05")+"		"+exeResult+" id:"+id+"		任务名称:"+t.Name+"		执行命令:"+cmd)
	result, err := exec.Execute(cmd)

	if err != nil {
		exeResult = "执行失败"
	} else {
		exeResult = "执行完成"
	}

	updateExecTime(id)

	fmt.Fprintln(gin.DefaultWriter, time.Now().Format("2006-01-02 15:04:05")+"		"+exeResult+" id:"+id+"		任务名称:"+t.Name+"		脚本结果:"+result)

	if t.AlertType == "mail" {
		rec_list := strings.Split(t.AlertRecMail, ",")
		send.SendMail(rec_list, "Coot["+t.Name+"]提醒你", result)
	}
}

func mTask(t *Task, typs string) string {
	var taskId string

	// 创建任务
	switch t.TimeType {
	case "1":
		// 秒执行
		number, err := strconv.Atoi(t.Time)
		error.Check(err, "秒时间格式化失败")

		if typs == "add" {
			task := gotask.NewTask(time.Second*time.Duration(number), func() { execute(t) })
			gotask.AddToTaskList(task)
			taskId = task.ID()
		} else if typs == "update" {
			gotask.ChangeInterval(t.TaskId, time.Second*time.Duration(number))
		}
	case "2":
		// 分钟执行
		number, err := strconv.Atoi(t.Time)
		error.Check(err, "分钟时间格式化失败")

		if typs == "add" {
			task := gotask.NewTask(time.Minute*time.Duration(number), func() { execute(t) })
			gotask.AddToTaskList(task)
			taskId = task.ID()
		} else if typs == "update" {
			gotask.ChangeInterval(t.TaskId, time.Minute*time.Duration(number))
		}
	case "3":
		// 小时执行
		number, err := strconv.Atoi(t.Time)
		error.Check(err, "小时时间格式化失败")

		if typs == "add" {
			task := gotask.NewTask(time.Hour*time.Duration(number), func() { execute(t) })
			gotask.AddToTaskList(task)
			taskId = task.ID()
		} else if typs == "update" {
			gotask.ChangeInterval(t.TaskId, time.Hour*time.Duration(number))
		}
	case "4":
		// 天执行
		task, err := gotask.NewDayTask(t.Time, func() { execute(t) })
		error.Check(err, "")
		gotask.AddToTaskList(task)
		taskId = task.ID()
	case "5":
		// 月执行
		task, err := gotask.NewMonthTask(t.Time, func() { execute(t) })
		error.Check(err, "")
		taskId = task.ID()
	case "6":
		// 年执行
		task := gotask.NewTask(time.Second*2, func() { execute(t) })
		gotask.AddToTaskList(task)
		taskId = task.ID()
	}
	return taskId
}

// 创建定时任务
func AddJob(t *Task) string {
	// 创建任务
	taskId := mTask(t, "add")

	// 返回 任务id
	return taskId
}

// 停止定时任务
func StopJob(taskId string) {
	// 停止任务
	gotask.Stop(taskId)
}

// 更新任务运行时间
func UpdateJobTime(t *Task) {
	mTask(t, "update")
}
