package main

import (
	"encoding/hex"
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

// 迭代查找
func (bc *BlockChain) Iterator() *BlockChainIterator {
	bci := &BlockChainIterator{bc.tip, bc.db}

	return bci
}

// 查找下一个区块 
func (bci *BlockChainIterator) Next() *Block {
	var block *Block 

	err := bci.db.View(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(bci.currentHash)
		block = DeserialBlock(encodedBlock)
		return nil 
	})
	if err != nil {
		log.Panic(err)
	}

	bci.currentHash = block.PrevBlockHash

	return block
}

func dbExists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}
	return true 
}

// 创建一个含创世区块的链
func NewBlockChain(address string) *BlockChain {
	if dbExists() == false {
		fmt.Println("No existing blockchain found, please create firt")
		os.Exit(1)
	}

	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	var tip []byte
	// 
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

// 创建一个新的区块链数据库 
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

	bc := BlockChain{tip, db}

	return &bc
}

// 查找为花费的交易
func (bc *BlockChain) FindUnspentTransactions(address string) []Transaction {
	var unspentTxs []Transaction 
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()

	for {
		block := bci.Next()

		for _, tx := range block.Transactions {
			txId := hex.EncodeToString(tx.ID)

		Outputs:
			for outIdx, out := range tx.Vout {
				if spentTXOs[txId] != nil {
					for _, spentOut := range spentTXOs[txId] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}

				if out.CanBeUnlockWith(address) {
					unspentTxs = append(unspentTxs, *tx)
				}
			} 

			if tx.IsCoinbase() == false {
				for _, in := range tx.Vin {
					if in.CanUnlockOutputWith(address) {
						inTxId := hex.EncodeToString(in.TxId)
						spentTXOs[inTxId] = append(spentTXOs[inTxId], in.Vout)
					}
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return unspentTxs
}

/*
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1] 
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
*/

// 查找未消费的输出
func (bc *BlockChain) FindUTXO(address string) []TxOutput {
	var UTXOs []TxOutput 

	unspentTransactions := bc.FindUnspentTransactions(address)

	for _, tx := range unspentTransactions {
		for _, out := range tx.Vout {
			if out.CanBeUnlockWith(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

func (bc *BlockChain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)
	unspentTXs := bc.FindUnspentTransactions(address)
	accumulated := 0

Work:
	for _, tx := range unspentTXs {
		txId := hex.EncodeToString(tx.ID)
		for outIdx, out := range tx.Vout {
			if out.CanBeUnlockWith(address) && accumulated < amount {
				accumulated += out.Value 
				unspentOutputs[txId] = append(unspentOutputs[txId], outIdx)

				if accumulated >= amount {
					break Work
				}
			}
		}
	}

	return accumulated, unspentOutputs
}