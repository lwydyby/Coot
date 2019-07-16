package setting

import (
	"Coot/core/dbUtil"
	"Coot/error"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Html(c *gin.Context) {
	data := getSetting() //订阅通知等
	c.HTML(http.StatusOK, "setting.html", gin.H{
		"dataList": data,
	})
}

/*获取配置*/
func getSetting() []map[string]interface{} {
	sql := "select id,type,info,setting_name,setting_dis,update_time,status from coot_setting"
	result := dbUtil.Query(sql)
	return result
}

/*更新邮件通知*/
func UpdateEmailInfo(c *gin.Context) {
	email := c.PostForm("email")
	id := c.PostForm("id")
	pass := c.PostForm("pass")
	host := c.PostForm("host")
	port := c.PostForm("port")
	subType := c.PostForm("type")
	and := "&&"
	info := host + and + port + and + email + and + pass
	fmt.Println(email, pass, host, port, subType)
	sql := `
		UPDATE  coot_setting 
		set	info = ?,
			status = ?,
			update_time = ?
		where id = ?;`
	dbUtil.Update(sql, info, 1, time.Now().Format("2006-01-02 15:04"), id)
	c.JSON(http.StatusOK, error.ErrSuccessNull())
}

/*更新登录用户信息*/
func UpdateLoginInfo(c *gin.Context) {
	loginName := c.PostForm("loginName")
	loginPwd := c.PostForm("loginPwd")
	id := c.PostForm("id")
	and := "&&"
	info := loginName + and + loginPwd
	fmt.Println(info)
	sql := `
		UPDATE  coot_setting 
		set	info = ?,
			status = ?,
			update_time = ?
		where id = ?;`
	dbUtil.Update(sql, info, 1, time.Now().Format("2006-01-02 15:04"), id)
	c.JSON(http.StatusOK, error.ErrSuccessNull())
}

/*更新设置状态*/
func UpdateStatusSetting(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")
	fmt.Println(id, status)
	sql := `update coot_setting
		set status = ?,
			update_time=?
		where id = ?`
	dbUtil.Update(sql, status, time.Now().Format("2006-01-02 15:04"), id)
	c.JSON(http.StatusOK, error.ErrSuccessNull())
}

//
//func UpdateAlertInfo(c *gin.Context) {
//	id := c.PostForm("alert_id")
//	status := c.PostForm("status")
//	sql := `
//		UPDATE coot_alert
//		SET status = ?
//		WHERE
//			id = ?;
//		`
//	dbUtil.Update(sql, status, id)
//	c.JSON(http.StatusOK, error.ErrSuccessNull())
//}
