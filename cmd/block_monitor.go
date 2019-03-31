package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func main() {
	/*for {
		c := usedrpc.NewUseRPC("http://127.0.0.1:8848")
		lastBlockNumberBefore, _ := c.UseBlockNumber()
		time.Sleep(10 * time.Minute)
		lastBlockNumberNow, _ := c.UseBlockNumber()
		if lastBlockNumberBefore == lastBlockNumberNow {
			sendEmail()
			return
		}
	}*/
	sendEmail()
}

func sendEmail() {
	auth := smtp.PlainAuth("", "usechain_test@163.com", "use123456", "smtp.163.com:25")
	to := []string{"zhouke@usechain.net"}
	nickname := "usechain_test@163.com"
	user := "usechain_test@163.com"
	subject := "usechain report"
	content_type := "Content-Type: text/plain; charset=UTF-8"
	body := "区块高度10分钟未变化"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.163.com:25", auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
}
