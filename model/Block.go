package model

type Block struct {
	Index int
	Timestamp int32
	Transaction Transaction
	Proof int
	previousHash string
}
