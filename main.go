package main

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

}

func CreateBlock(prevHash []byte, data []byte) *Block {

}

func GenesisBlock() *Block {

}

func (bc *BlockChain) AddBlock(data string) {

}

func CreateBlockChain() *BlockChain {

}

func ToHexInt(num int64) []byte {

}

func main() {

}
