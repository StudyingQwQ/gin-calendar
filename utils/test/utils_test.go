package test

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	mailAccount := viper.GetString("email.account")
	mailPassword := viper.GetString("email.password")
	fmt.Printf("mailAccount: %s, mailPassword: %s\n", mailAccount, mailPassword)
	e.From = "Test <13760831277@163.com>"
	e.To = []string{"2945484861@qq.com"}
	// e.Bcc = []string{"test_bcc@example.com"}
	// e.Cc = []string{"test_cc@example.com"}
	e.Subject = "验证码发送测试"
	// e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("您的验证码为<h1>123456</h1>")
	// e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"))
	err := e.SendWithStartTLS("smtp.163.com:25", smtp.PlainAuth("", "13760831277@163.com", "PAPQMFLGPZIIEPBJ", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
