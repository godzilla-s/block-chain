package main

import (
	"os"
	"fmt"
	"log"

	// 引入数据库 
	_ "github.com/syndtr/goleveldb/leveldb"
	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

// BlockChain 
type BlockChain struct {
	// blocks []*Block 
	tip 	[]byte 
	// db 		*leveldb.DB
	db 		*bolt.DB
}

type BlockChainIterator struct {
	currentHash		[]byte 
	//db 				*leveldb
	db 				*bolt.DB
}

func (bc *BlockChain) Iterator() *BlockChainIterator {
	bci := &BlockChainIterator{bc.tip, bc.db}

	return bci
}

func dbExists() bool {
	return true 
}

// 创建一个含创世区块的链
func NewBlockChain(address string) *BlockChain {
	if dbExists() == false {
		fmt.Println("No existing blockchain found, please create firt")
		os.Exit(1)
	}

	db, err := bolt.Open("", 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	var tip []byte

	err = db.Update(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte(blocksBucket))
		tip = b.Get([]byte("l"))

		return nil 
	})

	if err != nil {
		log.Panic(err)
	}

	bc := BlockChain{tip:tip, db: db}
	return &bc
}

func CreateBlockChain(address string) *BlockChain {
	// TODO 
	if dbExists() == false {
		fmt.Println("No existing blockchain found, please create firt")
		os.Exit(1)
	}

	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	var tip []byte 

	err = db.Update(func(tx *bolt.Tx) error {
		cbtx := NewCoinbaseTx(address, genesisCoinbaseData)
		genesis := NewGenesisBlock(cbtx)

		b, err := tx.CreateBucket([]byte(blocksBucket))
		if err != nil {
			log.Panic(err)
		}

		err = b.Put(genesis.Hash, genesis.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), genesis.Hash)
		if err != nil {
			log.Panic(err)
		}

		tip = genesis.Hash
		return nil 
	})
	if err != nil {
		log.Panic(err)
	}

	bc := BlockChain{tip: tip, }

	return &bc
}

// 查找为花费的交易
func (bc *BlockChain) FindUnspentTransactions(address string) []Transaction {
	var unspentTx []Transaction 
	spentTXOs := make(map[string][]int)

	for {
		
	}
}

/*
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1] 
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
*/