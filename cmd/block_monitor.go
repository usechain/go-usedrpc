package main

import (
	"github.com/usechain/go-usedrpc"
	"os/exec"
	"time"
)

func main() {
	for {
		c := usedrpc.NewUseRPC("http://127.0.0.1:8848")
		lastBlockNumberBefore, err := c.UseBlockNumber()
		if err != nil {
			sendEmail("./mail_Browser.sh")
		}
		time.Sleep(10 * time.Minute)
		lastBlockNumberNow, err := c.UseBlockNumber()
		if err != nil {
			sendEmail("./mail_Browser.sh")
		}
		if lastBlockNumberBefore == lastBlockNumberNow {
			sendEmail("./mail.sh")
			return
		}
	}
}

func sendEmail(sh string) {
	cmd := exec.Command(sh)
	cmd.Run()
}
