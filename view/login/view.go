package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Coot/core/dbUtil"
	"strconv"
	"fmt"
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

	sql := `select info from coot_setting where type="login";`
	result := dbUtil.Query(sql)

	info := result[0]["status"].(string)

	fmt.Println(info)
}
