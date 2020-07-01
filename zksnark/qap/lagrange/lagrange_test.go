package lagrange

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"

	"testing"

	"github.com/cloudflare/bn256"
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

	// fmt.Printf(" = l_{1}(2) = %d \n", base(big.NewInt(2))) // Should be 1
	// fmt.Printf(" = l_{1}(4) = %d \n", base(big.NewInt(4))) // Should be 0

	fmt.Println(new(big.Int).Mod(base(big.NewInt(2)), bn256.Order))
	fmt.Println(new(big.Int).Mod(base(big.NewInt(4)), bn256.Order))

	// fmt.Println(new(big.Int).Mod(base(big.NewInt(2)), big.NewInt(7)))
	// fmt.Println(new(big.Int).Mod(base(big.NewInt(4)), big.NewInt(7)))

	fmt.Println(" = = = = = = = = = = = = = = = = = = = = = = ")

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))

	base = BasisPolynomial(1, xCoords...)

	// fmt.Printf(" = l_{1}(2) = %d \n", base(big.NewInt(2))) // Should be 0
	// fmt.Printf(" = l_{1}(4) = %d \n", base(big.NewInt(4))) // Should be 1

	fmt.Println(new(big.Int).Mod(base(big.NewInt(2)), bn256.Order))
	fmt.Println(new(big.Int).Mod(base(big.NewInt(4)), bn256.Order))

	// fmt.Println(new(big.Int).Mod(base(big.NewInt(2)), big.NewInt(7)))
	// fmt.Println(new(big.Int).Mod(base(big.NewInt(4)), big.NewInt(7)))

	fmt.Println(" = = = = = = = = = = = = = = = = = = = = = = ")
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

	fmt.Printf(" - l_{3}(xCoords[2] == 6) = %d \n", new(big.Int).Mod(l[3](big.NewInt(6)), big.NewInt(11))) // Should be 1

	var err error

	var g1 *bn256.G1
	if _, g1, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var g2 = new(bn256.G1).ScalarMult(g1, l[3](big.NewInt(6)))

	fmt.Printf(" Is Equal        %t \n", bytes.Equal(g1.Marshal(), g2.Marshal()))
}

func TestExponentiation1(t *testing.T) {

	var xCoords []*big.Int
	var l []func(*big.Int) *big.Int

	// xCoords = []*big.Int{}

	// xCoords = append(xCoords, big.NewInt(2))
	// xCoords = append(xCoords, big.NewInt(4))
	// xCoords = append(xCoords, big.NewInt(6))
	// xCoords = append(xCoords, big.NewInt(8))

	// l = append(l, BasisPolynomial(0, xCoords...))

	// fmt.Printf(" - l_{0}(xCoords[0] == 2) = %d \n", l[0](big.NewInt(2))) // Should be 1
	// fmt.Printf(" - l_{0}(xCoords[1] == 4) = %d \n", l[0](big.NewInt(4))) // Should be 0
	// fmt.Printf(" - l_{0}(xCoords[2] == 6) = %d \n", l[0](big.NewInt(6))) // Should be 0
	// fmt.Printf(" - l_{0}(xCoords[2] == 8) = %d \n", l[0](big.NewInt(6))) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(3))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(9))

	l = append(l, BasisPolynomial(0, xCoords...))

	fmt.Println(new(big.Int).Mod(l[0](big.NewInt(3)), bn256.Order))
	fmt.Println(new(big.Int).Mod(l[0](big.NewInt(4)), bn256.Order))
	fmt.Println(new(big.Int).Mod(l[0](big.NewInt(6)), bn256.Order))
	fmt.Println(new(big.Int).Mod(l[0](big.NewInt(9)), bn256.Order))

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(3))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(9))

	l = append(l, BasisPolynomial(2, xCoords...))

	fmt.Println(new(big.Int).Mod(l[1](big.NewInt(3)), bn256.Order))
	fmt.Println(new(big.Int).Mod(l[1](big.NewInt(4)), bn256.Order))
	fmt.Println(new(big.Int).Mod(l[1](big.NewInt(6)), bn256.Order))
	fmt.Println(new(big.Int).Mod(l[1](big.NewInt(9)), bn256.Order))

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(3))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(9))

	l = append(l, BasisPolynomial(3, xCoords...))

	fmt.Println(new(big.Int).Mod(l[2](big.NewInt(3)), bn256.Order))
	fmt.Println(new(big.Int).Mod(l[2](big.NewInt(4)), bn256.Order))
	fmt.Println(new(big.Int).Mod(l[2](big.NewInt(6)), bn256.Order))
	fmt.Println(new(big.Int).Mod(l[2](big.NewInt(9)), bn256.Order))

	var check *big.Int

	check = Interpolate(
		big.NewInt(3), []int64{3, 1, 1}, l...,
	)

	fmt.Println(new(big.Int).Mod(check, bn256.Order))

	check = Interpolate(
		big.NewInt(6), []int64{3, 1, 1}, l...,
	)

	fmt.Println(new(big.Int).Mod(check, bn256.Order))

	check = Interpolate(
		big.NewInt(9), []int64{3, 1, 1}, l...,
	)

	fmt.Println(new(big.Int).Mod(check, bn256.Order))

	// l = append(l, BasisPolynomial(0, xCoords...))

	// fmt.Printf(" - l_{0}(xCoords[0] == 2) = %d \n", new(big.Int).Mod(l[0](big.NewInt(2)), big.NewInt(11))) // Should be 2

	// xCoords = []*big.Int{}

	// xCoords = append(xCoords, big.NewInt(3))
	// xCoords = append(xCoords, big.NewInt(4))
	// // xCoords = append(xCoords, big.NewInt(6))
	// // xCoords = append(xCoords, big.NewInt(9))

	// l = append(l, BasisPolynomial(1, xCoords...))

	// fmt.Printf(" - l_{0}(xCoords[0] == 2) = %d \n", new(big.Int).Mod(l[1](big.NewInt(2)), big.NewInt(11))) // Should be 2

	// var check = new(big.Int).Mul(l[0](big.NewInt(2)), l[1](big.NewInt(2)))

	// fmt.Printf(" = = C H E C K %d \n", new(big.Int).Mod(check, big.NewInt(11)))

	// xCoords = []*big.Int{}

	// xCoords = append(xCoords, big.NewInt(3))
	// xCoords = append(xCoords, big.NewInt(4))
	// xCoords = append(xCoords, big.NewInt(6))
	// xCoords = append(xCoords, big.NewInt(9))

	// check = Interpolate(
	// 	s, []int64{3, 1, 1},
	// 	lagrange.BasisPolynomial(0, []*big.Int{r, r1, r2, s1, s2}...),
	// 	lagrange.BasisPolynomial(3, []*big.Int{r, r1, r2, s1, s2}...),
	// 	lagrange.BasisPolynomial(4, []*big.Int{r, r1, r2, s1, s2}...),
	// ),

	// fmt.Printf(" = = C H E C K %d \n", new(big.Int).Mod(check, big.NewInt(11)))

	// var g1 *bn256.G1
	// if _, g1, err = bn256.RandomG1(rand.Reader); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	// v[0] = new(bn256.G1).ScalarMult(g1, l[1])

	// r1 = big.NewInt(3)
	// r2 = big.NewInt(4)
	// s1 = big.NewInt(6)
	// s2 = big.NewInt(9)

	// s = big.NewInt(2)

	// var v [3]*bn256.G1
	// var leftG []*big.Int

	// leftG = append(
	// 	leftG,
	// 	lagrange.Interpolate(
	// 		s, []int64{3, 1, 1},
	// 		lagrange.BasisPolynomial(0, []*big.Int{r, r1, r2, s1, s2}...),
	// 		lagrange.BasisPolynomial(3, []*big.Int{r, r1, r2, s1, s2}...),
	// 		lagrange.BasisPolynomial(4, []*big.Int{r, r1, r2, s1, s2}...),
	// 	),
	// )
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
