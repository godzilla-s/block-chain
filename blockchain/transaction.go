package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"bytes"
	"crypto/sha256"
)

const subsidy = 10

type Transaction struct {
	ID 		[]byte 		// 即Hash
	Vin 	[]TxInput  	// 输入
	Vout 	[]TxOutput  // 输出 
}

// 输入
type TxInput struct {
	TxId 		[]byte 	 // 之前交易的ID
	Vout 		int 	 // 索引
	ScriptSig 	string   // 签名 
}

// 输出
type TxOutput struct {
	Value 		int 	// 存储的币
	ScriptPubKey 	string  // 对输出进行锁定
}

// coinbase交易只有输出，没有输入 
func NewCoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txIn := TxInput{[]byte{}, -1, data}
	txOut := TxOutput{subsidy, to}
	tx := Transaction{nil, []TxInput{txIn}, []TxOutput{txOut}}
	tx.SetID()

	return &tx
}

func (tx Transaction) SetID() {
	var encoded  bytes.Buffer 
	var hash 	 [32]byte 

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].TxId) == 0 && tx.Vin[0].Vout == -1
}

func (in *TxInput) CanUnlockOutputWith(unlockData string) bool {
	return in.ScriptSig == unlockData
}

func (out *TxOutput) CanBeUnlockWith(unlockData string) bool {
	return out.ScriptPubKey == unlockData
}

func NewUTXOTransaction(from, to string, amount int, bc *BlockChain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	// acc, validOutputs := bc.FindSpeedableOutputs(from, amount)

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx 
}