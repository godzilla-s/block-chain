package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockChain()  // 创建一个区块链

	bc.AddBlock("Send 1 BTC to Thomas")
	bc.AddBlock("Send 1.5 BTC To Calvin")

	for _, block := range bc.blocks {
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewPoW(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}