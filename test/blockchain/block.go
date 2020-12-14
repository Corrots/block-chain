package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	// 区块链编号
	Index int64 `json:"index"`
	// 区块时间戳
	Timestamp int64 `json:"timestamp"`
	// 父区块哈希
	PrevHash string `json:"prev_hash"`
	// 当前区块哈希
	Hash string `json:"hash"`
	// 区块数据
	Data string `json:"data"`
}

func calculateHash(b *Block) string {
	data := string(b.Index) + string(b.Timestamp) + b.PrevHash + b.Data
	bytes := sha256.Sum256([]byte(data))
	return hex.EncodeToString(bytes[:])
}

func GenerateBlock(prevBlock *Block, data string) *Block {
	return &Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().Unix(),
		PrevHash:  prevBlock.Hash,
		Hash:      calculateHash(prevBlock),
		Data:      data,
	}
}

// 生成创世区块(index=0)
func GenerateGenesisBlock() *Block {
	dummy := &Block{Index: -1, Hash: ""}
	return GenerateBlock(dummy, "Genesis Block")
}
