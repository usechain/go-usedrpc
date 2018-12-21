package main

import (
	"fmt"
	"github.com/usechain/go-usedrpc"
	"math/big"
	"time"
)

func main() {
	c := usedrpc.NewUseRPC("http://127.0.0.1:8545")
	accounts, _ := c.UseAccounts()
	tx := usedrpc.T {
		From: accounts[1],
		To:   "0xfffffffffffffffffffffffffffffffff0000008",
		Value: big.NewInt(0),
		Data:  "",
		GasPrice: big.NewInt(1000000000000000000), //10^18
	}
	for {
		flag, err := c.UseSendTransaction(tx)
		fmt.Println("The tx hash:", err, flag)
		time.Sleep(50*time.Millisecond)
	}
}



