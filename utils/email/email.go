package email

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"math/rand"
	"net/smtp"
)

func SendMail(ToEmail string) (*string, error) {
	e := email.NewEmail()

	mailUserName := viper.GetString("emailCode.seedEmail")    //邮箱账号
	mailPassword := viper.GetString("emailCode.mailPassword") //邮箱授权码
	Subject := viper.GetString("emailCode.seedSubject")       //发送的主题
	addr := viper.GetString("emailCode.addr")
	host := viper.GetString("emailCode.host")
	serverName := viper.GetString("emailCode.Server")
	code := RandCode() //发送的验证码
	e.From = viper.GetString("emailCode.eFrom")

	e.To = []string{ToEmail}
	e.Subject = Subject
	e.HTML = []byte("<p>你好！</p>你的验证码为：<h1>" + code + "</h1>" + "<p>***该验证码5分钟内有效***</p>")
	err := e.SendWithTLS(addr, smtp.PlainAuth("", mailUserName, mailPassword, host),
		&tls.Config{InsecureSkipVerify: true, ServerName: serverName})
	if err != nil {
		return nil, err
	}
	return &code, nil
}

func RandCode() string {
	s := "1234567890"
	code := ""
	for i := 0; i < 6; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}
