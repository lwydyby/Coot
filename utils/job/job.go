package job

import (
	"github.com/domgoer/gotask"
	"time"
	"fmt"
	"strconv"
	"Coot/error"
)

type Task struct {
	/*
		taskId     任务ID  添加的时候为空
		ScriptType 脚本类型 shell 、python
		ScriptPath 脚本路径 应该是当前程序下的 script 目录
		TimeType   执行类型 1 秒执行，2 分钟执行，3 小时执行 ，4 每天指定时间执行，5 每月指定天和时间执行，6 年执行
		Time   	   周期时间
	*/
	TaskId     string
	ScriptType string
	ScriptPath string
	TimeType   int
	Time       string
}

// 执行任务
func exec(t *Task) {
	fmt.Print(t)
}

func mTask(t *Task, typs string) string {
	var taskId string

	// 创建任务
	switch t.TimeType {
	case 1:
		// 秒执行
		number, err := strconv.Atoi(t.Time)
		error.Check(err, "秒时间格式化失败")

		if typs == "add" {
			task := gotask.NewTask(time.Second*time.Duration(number), func() { exec(t) })
			gotask.AddToTaskList(task)
			taskId = task.ID()
		} else if typs == "update" {
			gotask.ChangeInterval(t.TaskId, time.Second*time.Duration(number))
		}
	case 2:
		// 分钟执行
		number, err := strconv.Atoi(t.Time)
		error.Check(err, "分钟时间格式化失败")

		if typs == "add" {
			task := gotask.NewTask(time.Minute*time.Duration(number), func() { exec(t) })
			gotask.AddToTaskList(task)
			taskId = task.ID()
		} else if typs == "update" {
			gotask.ChangeInterval(t.TaskId, time.Minute*time.Duration(number))
		}
	case 3:
		// 小时执行
		number, err := strconv.Atoi(t.Time)
		error.Check(err, "小时时间格式化失败")

		if typs == "add" {
			task := gotask.NewTask(time.Hour*time.Duration(number), func() { exec(t) })
			gotask.AddToTaskList(task)
			taskId = task.ID()
		} else if typs == "update" {
			gotask.ChangeInterval(t.TaskId, time.Hour*time.Duration(number))
		}
	case 4:
		// 天执行
		task, err := gotask.NewDayTask(t.Time, func() { exec(t) })
		error.Check(err, "")
		gotask.AddToTaskList(task)
		taskId = task.ID()
	case 5:
		// 月执行
		task, err := gotask.NewMonthTask(t.Time, func() { exec(t) })
		error.Check(err, "")
		taskId = task.ID()
	case 6:
		// 年执行
		task := gotask.NewTask(time.Second*2, func() { exec(t) })
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
