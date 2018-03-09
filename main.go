package main

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()

	bc.AddNewBlock("First transaction")
	bc.AddNewBlock("Second transaction")
	bc.AddNewBlock("Third transaction")

	for _, block := range bc.blocks {
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Prev hash : %x\n", block.PrevBlockHash)
		fmt.Printf("Current hash : %x\n", block.Hash)

		fmt.Println()
	}
}
