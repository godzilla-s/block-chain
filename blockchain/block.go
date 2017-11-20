package main 

import (
	"encoding/gob"
	"bytes"
	"crypto/sha256"
	"time"
	_ "strconv"
	"log"
)

// 区块 
// 添加Nonce, 对工作量证明验证时用到 
// 添加交易 
type Block struct {
	PrevBlockHash 		[]byte  // 上一个区块Hash
	Hash 				[]byte  // 当前Hash
	// Data 				[]byte 
	Transactions 		[]*Transaction  // 交易
	Timestamp 			int64
	Nonce 				int 
}

// 新建一个区块 
// 默认Nonce为0 
//func NewBlock(data string, prevBlockHash []byte) *Block {
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block {
		PrevBlockHash: 	prevBlockHash,  // 前一区块的Hash
		Hash:		   	[]byte{},		// 当前区块的Hash
		// Data:			[]byte(data),	// 数据
		Transactions:	transactions,
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
/*
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}
*/

// 将Block序列转化为字节数据
func (b *Block) Serialize() []byte {
	var result bytes.Buffer 
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)

	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// 将字节数组反序列化为一个Block 
func DeserialBlock(d []byte) *Block {
	var block Block 
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

// 创世区块 
func NewGenesisBlock(coinbase *Transaction) *Block {
	// return NewBlock("Genesis Block", []byte{}) //version 1
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

// 计算区块里所有交易Hash
func (b *Block) HashTransactions() []byte {
	var txHashes  	[][]byte
	var txHash 		[32]byte 

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}

	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

