package main

import (
	"fmt"
	"github.com/usechain/go-usedrpc"
	"os"
	"strconv"
)

func main() {
	start, _ := strconv.Atoi(os.Args[1])
	end, _ := strconv.Atoi(os.Args[2])
	var slot = end - start
	c := usedrpc.NewUseRPC("http://127.0.0.1:8848")
	blockStart, _ := c.UseGetBlockByNumber(start, false)
	blockNow, _ := c.UseGetBlockByNumber(end, false)
	timeslot := blockNow.Timestamp - blockStart.Timestamp
	var totalTxNum = 0
	var totalSize = 0
	var maxSize = 0
	var maxslot = 0
	var minerNum = make(map[string]int)

	for i := start; i < end; i++ {
		num, _ := c.UseGetBlockTransactionCountByNumber(i)
		totalTxNum += num
		blockbefore, _ := c.UseGetBlockByNumber(i, false)
		block, _ := c.UseGetBlockByNumber(i+1, false)
		totalSize += block.Size
		if maxSize < block.Size {
			maxSize = block.Size
		}
		tempslot := block.Timestamp - blockbefore.Timestamp
		if maxslot < tempslot {
			maxslot = tempslot
		}

		minerNum[blockbefore.Miner]++
	}
	fmt.Println("总交易数量：", totalTxNum)
	fmt.Println("总耗时(s)：", timeslot)
	fmt.Println("tps：", totalTxNum/timeslot)
	fmt.Println("平均出块延时(s)：", timeslot/slot)
	fmt.Println("最大出块延时(s)：", maxslot)
	fmt.Println("平均区块大小(byte)：", totalSize/slot)
	fmt.Println("最大区块大小(byte)：", maxSize)
	fmt.Println("平均每个区块包含交易数量：", totalTxNum/slot)
	for minernum := range minerNum {
		fmt.Println(minernum, "在此期间共产出", minerNum[minernum], "个区块")
	}
}
