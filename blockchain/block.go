package main 

import (
	"bytes"
	"crypto/sha256"
	"time"
	"strconv"
)

// 区块 
// 添加Nonce, 对工作量证明验证时用到 
type Block struct {
	PrevBlockHash 		[]byte  // 上一个区块Hash
	Hash 				[]byte
	Data 				[]byte 
	Timestamp 			int64
	Nonce 				int 
}

// 新建一个区块 
// 默认Nonce为0 
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block {
		PrevBlockHash: 	prevBlockHash,  // 前一区块的Hash
		Hash:		   	[]byte{},		// 当前区块的Hash
		Data:			[]byte(data),	// 数据
		Timestamp:		time.Now().Unix(),  // 时间戳
		Nonce: 			0,
	}

	// 加入工作量证明 
	pow := NewPoW(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block 
}

// 设置当前区块Hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// 创世区块 
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

type BlockChain struct {
	blocks []*Block 
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1] 
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}