package minimin

import (
	"log"
	"strings"

	"gopkg.in/gomail.v2"
)

// MailboxConf 邮箱配置
type MailboxConf struct {
	Sender    string `json:"sender"`
	SMTPPort  int
	SPassword string
	SMTPAddr  string
	Title     string `json:"title"`     // 邮件标题
	Body      string `json:"body"`      // 邮件内容
	Addressee string `json:"addressee"` // 收件人列表
}

func SendEmail(mailConf MailboxConf) (err error) {
	m := gomail.NewMessage()
	var addresseeList = strings.Split(mailConf.Addressee, ";")
	m.SetHeader(`From`, mailConf.Sender)
	m.SetHeader(`To`, addresseeList...)
	m.SetHeader(`Subject`, mailConf.Title)
	m.SetBody(`text/html`, mailConf.Body)
	err = gomail.NewDialer(
		mailConf.SMTPAddr, mailConf.SMTPPort,
		mailConf.Sender, mailConf.SPassword,
	).DialAndSend(m)
	if err != nil {
		log.Fatalf("Send Email Fail, %s", err.Error())
	} else {
		log.Println("Send Email Success")
	}
	return
}
