package main

import (
	"github.com/usechain/go-usedrpc"
	"os/exec"
	"time"
)

func main() {
	for {
		c := usedrpc.NewUseRPC("http://127.0.0.1:8848")
		lastBlockNumberBefore, _ := c.UseBlockNumber()
		time.Sleep(10 * time.Minute)
		lastBlockNumberNow, _ := c.UseBlockNumber()
		if lastBlockNumberBefore == lastBlockNumberNow {
			sendEmail()
			return
		}
	}
}

func sendEmail() {
	cmd := exec.Command("./mail.sh")

	err := cmd.Run()
	if err != nil {
		//fmt.Println("Execute Command failed:" + err.Error())
		return
	}
	//fmt.Println("Execute Command finished.")
}
