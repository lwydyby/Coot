package setting

import (
	"Coot/core/dbUtil"
	"Coot/error"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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

/*检查是否配置信息*/
func checkInfo(id string) bool {
	sql := "select id,info,type from coot_setting where id = ?"
	result := dbUtil.Query(sql, id)
	info := result[0]["info"].(string)
	typeStr := result[0]["type"].(string)
	infoArr := strings.Split(info, "&&")
	num := len(infoArr)
	if num == 4 && typeStr == "mail" {
		return true
	}
	if num == 2 && typeStr == "login" {
		return true
	}
	return false
}

/*更新邮件通知*/
func UpdateEmailInfo(c *gin.Context) {
	email := c.PostForm("email")
	id := c.PostForm("id")
	pass := c.PostForm("pass")
	host := c.PostForm("host")
	port := c.PostForm("port")
	//subType := c.PostForm("type")
	and := "&&"
	info := host + and + port + and + email + and + pass
	sql := `
		UPDATE  coot_setting 
		set	info = ?,
			status = ?,
			update_time = ?
		where id = ?;`
	dbUtil.Update(sql, info, 0, time.Now().Format("2006-01-02 15:04"), id)
	c.JSON(http.StatusOK, error.ErrSuccessNull())
}

/*更新登录用户信息*/
func UpdateLoginInfo(c *gin.Context) {
	loginName := c.PostForm("loginName")
	loginPwd := c.PostForm("loginPwd")
	id := c.PostForm("id")
	and := "&&"
	info := loginName + and + loginPwd
	sql := `
		UPDATE  coot_setting 
		set	info = ?,
			status = ?,
			update_time = ?
		where id = ?;`
	dbUtil.Update(sql, info, 0, time.Now().Format("2006-01-02 15:04"), id)
	c.JSON(http.StatusOK, error.ErrSuccessNull())
}

/*更新设置状态*/
func UpdateStatusSetting(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")
	if !checkInfo(id) && status == "1" {
		c.JSON(http.StatusOK, gin.H{"code": 10003, "msg": "请配置后在启用", "data": nil})
		return
	}
	sql := `update coot_setting
		set status = ?,
			update_time=?
		where id = ?`
	dbUtil.Update(sql, status, time.Now().Format("2006-01-02 15:04"), id)
	c.JSON(http.StatusOK, error.ErrSuccessNull())
}

/*根据id获取设置详情*/
func GetSettingInfo(c *gin.Context) {
	id, _ := c.GetQuery("id")
	sql := `select id,type,info,status from coot_setting where id = ?`
	result := dbUtil.Query(sql, id)
	c.JSON(http.StatusOK, error.ErrSuccess(result))
}
