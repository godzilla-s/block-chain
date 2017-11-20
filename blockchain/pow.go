package main 

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

// 难度值， hash前n为0 
const targetBits = 24 

const maxNonce = math.MaxInt64 

type PoW struct {
	block  *Block 
	target *big.Int 
}

// 创建PoW 
func NewPoW(b *Block) *PoW {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))  // 左移位targetBits

	pow := &PoW{b, target}

	return pow 
}

// 用到的有效数据
func (pow *PoW) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			// pow.block.Data,
			pow.block.Hash,
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// PoW 算法
func (pow *PoW) Run() (int, []byte) {
	var hashInt big.Int 
	var hash [32]byte 
	nonce := 0

	// fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)

		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("\r%x", hash)
			break
		} else {
			nonce++ 
		}
	}
	fmt.Printf("\n\n")

	return nonce, hash[:]
}

func (pow *PoW) Validate() bool {
	var hashInt big.Int 
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}