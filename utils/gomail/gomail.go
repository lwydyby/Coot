package gomail

import (
	"strconv"
	"wshell-server/src/gopkg.in/gomail.v2"
)

//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
var mailConn = map[string]string{
	"user": "jieggi@126.com",
	"pass": "wangyi126",
	"host": "smtp.126.com",
	"port": "465",
}

//邮件主题为"Hello"
//subject := "验证码"
//// 邮件正文
//body := "您的wshell注册验证码为:【"+code+"】"
//if typestr=="resetpwd"{
//body="您正在修改wshell的账号密码，验证码为:【"+code+"】"
//}
//errEmail:=SendMail(mailTo, subject, body)

func SendMail(mailTo []string, subject string, body string) error {

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	m := gomail.NewMessage()
	m.SetHeader("From", "<"+mailConn["user"]+">") //这种方式可以添加别名，即， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                  //发送给多个用户
	m.SetHeader("Subject", subject)               //设置邮件主题
	m.SetBody("text/html", body)                  //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}
