package send

import (
	"strconv"
	"gopkg.in/gomail.v2"
	"Coot/core/dbUtil"
	"strings"
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
)

func findMailConfig() []map[string]interface{} {
	sql := `select info from coot_alert where type="mail";`
	result := dbUtil.Query(sql)
	return result
}

func SendMail(mailTo []string, subject string, body string) error {

	// smtp&&post&&user&&password
	info := findMailConfig()[0]["info"]

	config := strings.Split(info.(string), "&&")

	port, _ := strconv.Atoi(config[1])
	m := gomail.NewMessage()

	m.SetHeader("From", "<"+config[2]+">")
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(config[0], port, config[2], config[3])

	err := d.DialAndSend(m)

	fmt.Fprintln(gin.DefaultWriter, time.Now().Format("2006-01-02 15:04:05"))

	return err
}
