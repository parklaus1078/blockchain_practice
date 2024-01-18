package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("hello, world")
}

type Block struct { // Definition of a Block
	data         map[string]interface{} // data will be map of strings
	hash         string                 // hash is an identifier generated using cryptography : is string
	previousHash string                 // previousHash is the cruptographic hash of the last block in the blockchain. It links the blocks in the chain and improves the chain's security- : value is string
	timestamp    time.Time              // timestamp is in type Time of time package
	pow          int                    // pow(proof of work) is the amount of effort taken to derive the current block's hash : is integer
}

type Blockchain struct { // Definition of Blockchain
	genesisBlock Block   // genesisBlock is the very first block added to the blockchain : is Block
	chain        []Block // chain is the list of Blocks defined previously
	difficulty   int     // difficulty is the minimum effort miners must undertake to mine : is integer
}

func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}
