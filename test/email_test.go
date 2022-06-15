package test

import (
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <604862834@qq.com>"
	e.To = []string{"604862834@qq.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码是<b>123123</b>")
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "604862834@qq.com", "whagfynxylvfbfgh", "smtp.qq.com"))
	if err != nil {
		t.Fatal(err)
	}
}
