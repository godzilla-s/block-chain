package main 

import (
	"bytes"
	"crypto/sha256"
	"time"
	"strconv"
)

// 区块 
type Block struct {
	PrevBlockHash 		[]byte  // 上一个区块Hash
	Hash 				[]byte
	Data 				[]byte 
	Timestamp 			int64
}


func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block {
		PrevBlockHash: 	prevBlockHash,
		Hash:		   	[]byte{},
		Data:			[]byte{},
		Timestamp:		time.Now().Unix(),
	}

	block.SetHash()

	return block 
}

func (b *Block) SetHash() {
	
}