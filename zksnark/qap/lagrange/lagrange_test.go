package lagrange

import (
	"fmt"
	"math/big"

	"testing"
)

func TestBasisPolynomial1(t *testing.T) {

	var xCoords []*big.Int
	var base func(*big.Int) *big.Int

	// xCoords = *new([]*big.Int)
	// xCoords = make([]*big.Int, 2)

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))

	base = BasisPolynomial(0, xCoords...)

	fmt.Printf(" = l_{1}(2) = %d \n", base(big.NewInt(2))) // Should be 1
	fmt.Printf(" = l_{1}(4) = %d \n", base(big.NewInt(4))) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))

	base = BasisPolynomial(1, xCoords...)

	fmt.Printf(" = l_{1}(2) = %d \n", base(big.NewInt(2))) // Should be 0
	fmt.Printf(" = l_{1}(4) = %d \n", base(big.NewInt(4))) // Should be 1
}

func TestBasisPolynomial2(t *testing.T) {

	var xCoords []*big.Int
	// var base func(*big.Int) *big.Int

	var l []func(*big.Int) *big.Int

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))

	l = append(l, BasisPolynomial(0, xCoords...))

	fmt.Printf(" - l_{0}(xCoords[0] == 2) = %d \n", l[0](big.NewInt(2))) // Should be 1
	fmt.Printf(" - l_{0}(xCoords[1] == 4) = %d \n", l[0](big.NewInt(4))) // Should be 0
	fmt.Printf(" - l_{0}(xCoords[2] == 6) = %d \n", l[0](big.NewInt(6))) // Should be 0

	// base = BasisPolynomial(0, xCoords...)

	// fmt.Printf(" = l_{1}(2) = %d \n", base(big.NewInt(2))) // Should be 1
	// fmt.Printf(" = l_{1}(4) = %d \n", base(big.NewInt(4))) // Should be 0
	// fmt.Printf(" = l_{1}(6) = %d \n", base(big.NewInt(6))) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))

	l = append(l, BasisPolynomial(1, xCoords...))

	fmt.Printf(" - l_{1}(xCoords[0] == 2) = %d \n", l[1](big.NewInt(2))) // Should be 0
	fmt.Printf(" - l_{1}(xCoords[1] == 4) = %d \n", l[1](big.NewInt(4))) // Should be 1
	fmt.Printf(" - l_{1}(xCoords[2] == 6) = %d \n", l[1](big.NewInt(6))) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))

	l = append(l, BasisPolynomial(2, xCoords...))

	fmt.Printf(" - l_{1}(xCoords[0] == 2) = %d \n", l[2](big.NewInt(2))) // Should be 0
	fmt.Printf(" - l_{1}(xCoords[1] == 4) = %d \n", l[2](big.NewInt(4))) // Should be 0
	fmt.Printf(" - l_{1}(xCoords[2] == 6) = %d \n", l[2](big.NewInt(6))) // Should be 1

	fmt.Println(new(big.Int).ModInverse(big.NewInt(3), big.NewInt(7)))
	var a = new(big.Int).Mul(big.NewInt(2), new(big.Int).ModInverse(big.NewInt(3), big.NewInt(7))) // 2 / 3
	// var b = a.Int64() % big.NewInt(7)

	fmt.Println(new(big.Int).Mod(a, big.NewInt(7)))
	// fmt.Println(new(big.Int).Mul(big.NewInt(2), new(big.Int).ModInverse(big.NewInt(3), big.NewInt(7))) % big.NewInt(7))
}

func TestBasisPolynomial3(t *testing.T) {

	var xCoords []*big.Int
	var l []func(*big.Int) *big.Int

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(8))

	l = append(l, BasisPolynomial(0, xCoords...))

	fmt.Printf(" - l_{0}(xCoords[0] == 2) = %d \n", l[0](big.NewInt(2))) // Should be 1
	fmt.Printf(" - l_{0}(xCoords[1] == 4) = %d \n", l[0](big.NewInt(4))) // Should be 0
	fmt.Printf(" - l_{0}(xCoords[2] == 6) = %d \n", l[0](big.NewInt(6))) // Should be 0
	fmt.Printf(" - l_{0}(xCoords[2] == 8) = %d \n", l[0](big.NewInt(6))) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(8))

	l = append(l, BasisPolynomial(1, xCoords...))

	fmt.Printf(" - l_{1}(xCoords[0] == 2) = %d \n", l[1](big.NewInt(2))) // Should be 0
	fmt.Printf(" - l_{1}(xCoords[1] == 4) = %d \n", l[1](big.NewInt(4))) // Should be 1
	fmt.Printf(" - l_{1}(xCoords[2] == 6) = %d \n", l[1](big.NewInt(6))) // Should be 0
	fmt.Printf(" - l_{1}(xCoords[2] == 8) = %d \n", l[1](big.NewInt(6))) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(8))

	l = append(l, BasisPolynomial(2, xCoords...))

	fmt.Printf(" - l_{2}(xCoords[0] == 2) = %d \n", l[2](big.NewInt(2))) // Should be 0
	fmt.Printf(" - l_{2}(xCoords[1] == 4) = %d \n", l[2](big.NewInt(4))) // Should be 0
	fmt.Printf(" - l_{2}(xCoords[2] == 6) = %d \n", l[2](big.NewInt(6))) // Should be 1
	fmt.Printf(" - l_{2}(xCoords[2] == 8) = %d \n", l[2](big.NewInt(8))) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(8))

	l = append(l, BasisPolynomial(3, xCoords...))

	fmt.Printf(" - l_{3}(xCoords[0] == 2) = %d \n", l[3](big.NewInt(2))) // Should be 0
	fmt.Printf(" - l_{3}(xCoords[1] == 4) = %d \n", l[3](big.NewInt(4))) // Should be 0
	fmt.Printf(" - l_{3}(xCoords[2] == 6) = %d \n", l[3](big.NewInt(6))) // Should be 1
	fmt.Printf(" - l_{3}(xCoords[2] == 8) = %d \n", l[3](big.NewInt(8))) // Should be 0
}

// func TestBasisPolynomialEvals(t *testing.T) {

// 	var xCoords []*big.Int
// 	var l []func(*big.Int) *big.Int

// 	xCoords = []*big.Int{}

// 	xCoords = append(xCoords, big.NewInt(2))
// 	xCoords = append(xCoords, big.NewInt(4))
// 	xCoords = append(xCoords, big.NewInt(6))
// 	xCoords = append(xCoords, big.NewInt(8))

// 	l = append(l, BasisPolynomial(0, xCoords...))

// 	fmt.Printf(" - l_{0}(xCoords[0] == 2) = %d \n", l[0](big.NewInt(2))) // Should be 1
// 	fmt.Printf(" - l_{0}(xCoords[1] == 4) = %d \n", l[0](big.NewInt(4))) // Should be 0
// 	fmt.Printf(" - l_{0}(xCoords[2] == 6) = %d \n", l[0](big.NewInt(6))) // Should be 0
// 	fmt.Printf(" - l_{0}(xCoords[2] == 8) = %d \n", l[0](big.NewInt(6))) // Should be 0

// 	// var empty [][]byte

// 	// // An empty slice should fail
// 	// _, err := Transpose(empty)
// 	// if err == nil {
// 	// 	t.Fail()
// 	// }

// 	// blocks := make([][]byte, 2)
// 	// for i := range blocks {
// 	// 	blocks[i] = []byte{1, 2}
// 	// }

// 	// transposed, err := Transpose(blocks)
// 	// if err != nil {
// 	// 	t.Fail()
// 	// }
// 	// if len(transposed) != len(blocks[0]) {
// 	// 	t.Fail()
// 	// }
// 	// if len(transposed[0]) != len(blocks) {
// 	// 	t.Fail()
// 	// }
// 	// if transposed[0][0] != 1 {
// 	// 	t.Fail()
// 	// }
// 	// if transposed[1][1] != 2 {
// 	// 	t.Fail()
// 	// }
// }
