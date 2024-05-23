package utils

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

var myKey = []byte("gin-calendar-key")

// 生成 token
func GenerateToken(identity, email string) (string, error) {
	fmt.Printf("%s", string(myKey))
	UserClaim := &UserClaims{
		Identity:         identity,
		Email:            email,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析 token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

// MD5加密
func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// 生成唯一码
func GetUUID() string {
	u := uuid.New()
	return fmt.Sprintf("%x", u)
}

// 发送邮件
func MailReminder(mail, content string) error {
	mailAccount := "13760831277@163.com"
	mailPassword := "PAPQMFLGPZIIEPBJ"
	e := email.NewEmail()
	e.From = "Test <" + mailAccount + ">"
	e.To = []string{mail}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码为<h1>" + content + "</h1>")
	err := e.SendWithStartTLS("smtp.163.com:25", smtp.PlainAuth("", mailAccount, mailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	return nil
}
