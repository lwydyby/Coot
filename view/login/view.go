package login

import (
	"Coot/core/dbUtil"
	"Coot/error"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func findLoginStatus() string {
	sql := `select status from coot_setting where type="login";`
	result := dbUtil.Query(sql)

	status := strconv.FormatInt(result[0]["status"].(int64), 10)

	return status
}

func Html(c *gin.Context) {

	status := findLoginStatus()

	if status == "1" {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	} else {
		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func Jump(c *gin.Context) {

	status := findLoginStatus()

	if status == "1" {
		c.Redirect(http.StatusMovedPermanently, "/login")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func Login(c *gin.Context) {
	loginName := c.PostForm("loginName")
	loginPwd := c.PostForm("loginPwd")

	sql := `select info,status from coot_setting where type="login";`
	result := dbUtil.Query(sql)

	info := result[0]["info"].(string)

	infoArr := strings.Split(info, "&&")

	if loginName == infoArr[0] && loginPwd == infoArr[1] {

		cookie := &http.Cookie{
			Name:     "is_login",
			Value:    infoArr[0],
			Path:     "/",
			HttpOnly: true,
		}

		http.SetCookie(c.Writer, cookie)

		c.JSON(http.StatusOK, error.ErrSuccessNull())
		return
	}

	c.JSON(http.StatusOK, error.ErrLoginFail())
}