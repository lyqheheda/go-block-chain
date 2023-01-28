package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"time"
)

type Block struct {
	Timestamp int64
	PrevHash  []byte
	Hash      []byte
	Data      []byte
}

type BlockChain struct {
	Blocks []*Block
}

func (b *Block) SetHash() {
	information := bytes.Join([][]byte{ToHexInt(b.Timestamp), b.PrevHash, b.Data}, []byte{})
	hash := sha256.Sum256(information)
	b.Hash = hash[:]
}

func ToHexInt(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func CreateBlock(prevHash []byte, data []byte) *Block {
	block := Block{time.Now().Unix(), prevHash, []byte{}, data}
	block.SetHash()
	return &block
}

func GenesisBlock() *Block {
	block := CreateBlock([]byte{}, []byte("Hello, blockchain!"))
	return block
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	bc.Blocks = append(bc.Blocks, CreateBlock(prevBlock.Hash, []byte(data)))
}

func CreateBlockChain() *BlockChain {
	blockchain := BlockChain{}
	blockchain.Blocks = append(blockchain.Blocks, GenesisBlock())
	return &blockchain
}

func main() {
	blockchain := CreateBlockChain()
	time.Sleep(time.Second)
	blockchain.AddBlock("After genesis, I have something to say.")
	time.Sleep(time.Second)
	blockchain.AddBlock("LinBenWei is awesome!")
	time.Sleep(time.Second)
	blockchain.AddBlock("I can't wait to follow his github!")
	time.Sleep(time.Second)

	for _, block := range blockchain.Blocks {
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
	}
}
