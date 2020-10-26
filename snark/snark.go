package snark

import (
	"fmt"
	"math/big"
	"math/rand"
)

const (
	PLUS = iota
	MINUS
)

type BinaryTree struct {
	left  *BinaryTree
	right *BinaryTree

	wire int
	op   string
}

func E1QAP(order *big.Int) bool {

	return true
}

func E1SNARK(order *big.Int) bool {

	// Seeding with the same value results in the same random sequence each run.
	// For different numbers, seed with a different value, such as
	// time.Now().UnixNano(), which yields a constantly-changing number.

	rand.Seed(86)

	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))

	// var g = big.NewInt()

	return true
}
