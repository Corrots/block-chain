package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	hashFunc("test1")
	hashFunc("test1")
	hashFunc("test")
}

func hashFunc(toHash string) string {
	bytes := sha256.Sum256([]byte(toHash))
	hashStr := hex.EncodeToString(bytes[:])
	fmt.Printf("%s -> %s\n", toHash, hashStr)
	return hashStr
}