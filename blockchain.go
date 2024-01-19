package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

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

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}

	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

func (b *Blockchain) addBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}

	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}

	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

func (b Blockchain) isValid() bool { // Checking validity of all blocks in the chain
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}

	return true
}

func main() {
	blockchain := CreateBlockchain(2)

	blockchain.addBlock("Person1", "Receiver3", 10)
	blockchain.addBlock("Person2", "Receiver2", 23)
	blockchain.addBlock("Receiver3", "Receiver5", 15)
	blockchain.addBlock("Receiver2", "Receiver5", 10)

	fmt.Println(blockchain.isValid())
}
