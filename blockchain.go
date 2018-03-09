package main

import  (
	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

type Blockchain struct {
	tip []byte
	db *bolt.DB
}

type BlockchainIterator struct {
	currentHash []byte
	db *bolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.tip, bc.db}
}

func (i *BlockchainIterator) Next() *Block {
	var block *Block

	i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	i.currentHash = block.PrevBlockHash
	return block
}

func (bc *Blockchain) AddNewBlock(data string) {
	 var lastHash []byte

	bc.db.View(func(tx *bolt.Tx) error {
	 	b := tx.Bucket([]byte(blocksBucket))
	 	lastHash = b.Get([]byte("1"))

	 	return nil
	 })

	 newBlock := NewBlock(data, lastHash)

	 bc.db.Update(func(tx *bolt.Tx) error {
	 	b := tx.Bucket([]byte(blocksBucket))
	 	b.Put(newBlock.Hash, newBlock.Serialize())
	 	b.Put([]byte("1"), newBlock.Hash)
	 	bc.tip = newBlock.Hash

	 	return nil
	 })
}

func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)

	err = db.Update(func(tx *bolt.Tx) error  {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("1"), genesis.Hash)
			tip = genesis.Hash
			if (err == nil) {

			}
		} else {
			tip = b.Get([]byte("1"))
		}

		if err == nil {
		}
		return nil
	})

	bc := Blockchain{tip, db}
	return &bc
}