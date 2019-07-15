package setting

import (
	"Coot/core/dbUtil"
	"Coot/error"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Html(c *gin.Context) {
	data := getSetting()
	c.HTML(http.StatusOK, "setting.html", gin.H{
		"settingList": data,
	})
}

func getSetting() []map[string]interface{} {
	sql := "select id,type,info,create_time,status from coot_alert"
	result := dbUtil.Query(sql)
	return result
}

func AddAlertInfo(c *gin.Context) {
	email := c.PostForm("email")
	pass := c.PostForm("pass")
	host := c.PostForm("host")
	port := c.PostForm("port")
	subType := c.PostForm("type")
	fmt.Println(email, pass, host, port, subType)
	//sql := `
	//	INSERT INTO coot_alert (
	//		type,
	//		info,
	//		status,
	//		create_time,
	//	)
	//	VALUES
	//		(?,?,?,?);`
	//dbUtil.Insert(sql, info, subType, 1, time.Now().Format("2006-01-02 15:04"))
	c.JSON(http.StatusOK, error.ErrSuccessNull())
}

func UpdateAlertInfo(c *gin.Context) {
	id := c.PostForm("alert_id")
	status := c.PostForm("status")
	sql := `
		UPDATE coot_alert
		SET status = ?
		WHERE
			id = ?;
		`
	dbUtil.Update(sql, status, id)
	c.JSON(http.StatusOK, error.ErrSuccessNull())
}
