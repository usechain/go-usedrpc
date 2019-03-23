package main

import (
	"fmt"
	"github.com/usechain/go-usedrpc"
	"math/big"
	"math/rand"
	"time"
)

var letters = []rune("0123456789abcdef")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	c := usedrpc.NewUseRPC("http://127.0.0.1:8848")
	accounts, _ := c.UseAccounts()
	for {
		tx := usedrpc.T{
			From:     accounts[0],
			To:       "0x" + randSeq(40),
			Value:    big.NewInt(int64(rand.Intn(10000000000000000))),
			Data:     "",
			GasPrice: big.NewInt(40000000000), //10^10
		}
		flag, err := c.UseSendTransaction(tx)
		fmt.Println("The tx hash:", err, flag)
		time.Sleep(50 * time.Millisecond)
	}
}
