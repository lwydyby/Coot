package login

import (
	"Coot/core/dbUtil"
	"Coot/error"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func findLoginStatus() string {
	sql := `select status,info from coot_setting where type="login";`
	result := dbUtil.Query(sql)

	status := strconv.FormatInt(result[0]["status"].(int64), 10)
	info:=result[0]["info"].(string)
	infoArr:=strings.Split(info,"&&")
	return status+"&&"+infoArr[0]
}

func Html(c *gin.Context) {

	loginInfo := findLoginStatus()
	loginArr:=strings.Split(loginInfo,"&&")
	if loginArr[0] == "1"{
		c.HTML(http.StatusOK, "login.html", gin.H{})
	} else {
		c.Redirect(http.StatusFound, "/task")
	}
}

func Jump(c *gin.Context) {
	loginInfo := findLoginStatus()
	loginArr:=strings.Split(loginInfo,"&&")
	loginCookie,_:=c.Cookie("is_login")
	fmt.Println(loginCookie,"是否登录")
	if loginArr[0] == "1"&&loginArr[1]!=loginCookie {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}else{
		c.Next()
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
		c.SetCookie("is_login",infoArr[0] ,60*60*24, "/", "127.0.0.1", false, true)
		c.JSON(http.StatusOK, error.ErrSuccessNull())
		return
	}
	c.JSON(http.StatusOK, error.ErrLoginFail())
}
