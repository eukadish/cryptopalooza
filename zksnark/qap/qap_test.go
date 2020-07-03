package qap

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"

	"testing"

	"github.com/cloudflare/bn256"
)

func TestBasisPolynomial1(t *testing.T) {

	var order = bn256.Order.Set(big.NewInt(11))

	var xCoords []*big.Int
	var base func(*big.Int) *big.Int

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))

	base = BasisPolynomial(order, 0, xCoords...)

	fmt.Println(new(big.Int).Mod(base(big.NewInt(2)), bn256.Order))
	fmt.Println(new(big.Int).Mod(base(big.NewInt(4)), bn256.Order))

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))

	base = BasisPolynomial(order, 1, xCoords...)

	fmt.Println(new(big.Int).Mod(base(big.NewInt(2)), bn256.Order))
	fmt.Println(new(big.Int).Mod(base(big.NewInt(4)), bn256.Order))
}

func TestBasisPolynomial2(t *testing.T) {

	var order = bn256.Order.Set(big.NewInt(11))

	var xCoords []*big.Int
	// var base func(*big.Int) *big.Int

	var l []func(*big.Int) *big.Int

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))

	l = append(l, BasisPolynomial(order, 0, xCoords...))

	fmt.Printf(" - l_{0}(xCoords[0] == 2) = %d \n", l[0](big.NewInt(2))) // Should be 1
	fmt.Printf(" - l_{0}(xCoords[1] == 4) = %d \n", l[0](big.NewInt(4))) // Should be 0
	fmt.Printf(" - l_{0}(xCoords[2] == 6) = %d \n", l[0](big.NewInt(6))) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))

	l = append(l, BasisPolynomial(order, 1, xCoords...))

	fmt.Printf(" - l_{1}(xCoords[0] == 2) = %d \n", l[1](big.NewInt(2))) // Should be 0
	fmt.Printf(" - l_{1}(xCoords[1] == 4) = %d \n", l[1](big.NewInt(4))) // Should be 1
	fmt.Printf(" - l_{1}(xCoords[2] == 6) = %d \n", l[1](big.NewInt(6))) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))

	l = append(l, BasisPolynomial(order, 2, xCoords...))

	fmt.Printf(" - l_{2}(xCoords[0] == 2) = %d \n", l[2](big.NewInt(2))) // Should be 0
	fmt.Printf(" - l_{2}(xCoords[1] == 4) = %d \n", l[2](big.NewInt(4))) // Should be 0
	fmt.Printf(" - l_{2}(xCoords[2] == 6) = %d \n", l[2](big.NewInt(6))) // Should be 1

	fmt.Println(new(big.Int).ModInverse(big.NewInt(3), big.NewInt(7)))
	var a = new(big.Int).Mul(big.NewInt(2), new(big.Int).ModInverse(big.NewInt(3), big.NewInt(7))) // 2 / 3
	// var b = a.Int64() % big.NewInt(7)

	fmt.Println(new(big.Int).Mod(a, big.NewInt(7)))
	// fmt.Println(new(big.Int).Mul(big.NewInt(2), new(big.Int).ModInverse(big.NewInt(3), big.NewInt(7))) % big.NewInt(7))
}

func TestBasisPolynomial3(t *testing.T) {

	var order = bn256.Order.Set(big.NewInt(11))

	var xCoords []*big.Int
	var l []func(*big.Int) *big.Int

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(8))

	l = append(l, BasisPolynomial(order, 0, xCoords...))

	fmt.Printf(" - l_{0}(xCoords[0] == 2) = %d \n", l[0](big.NewInt(2))) // Should be 1
	fmt.Printf(" - l_{0}(xCoords[1] == 4) = %d \n", l[0](big.NewInt(4))) // Should be 0
	fmt.Printf(" - l_{0}(xCoords[2] == 6) = %d \n", l[0](big.NewInt(6))) // Should be 0
	fmt.Printf(" - l_{0}(xCoords[2] == 8) = %d \n", l[0](big.NewInt(6))) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(8))

	l = append(l, BasisPolynomial(order, 1, xCoords...))

	fmt.Printf(" - l_{1}(xCoords[0] == 2) = %d \n", l[1](big.NewInt(2))) // Should be 0
	fmt.Printf(" - l_{1}(xCoords[1] == 4) = %d \n", l[1](big.NewInt(4))) // Should be 1
	fmt.Printf(" - l_{1}(xCoords[2] == 6) = %d \n", l[1](big.NewInt(6))) // Should be 0
	fmt.Printf(" - l_{1}(xCoords[2] == 8) = %d \n", l[1](big.NewInt(6))) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(8))

	l = append(l, BasisPolynomial(order, 2, xCoords...))

	fmt.Printf(" - l_{2}(xCoords[0] == 2) = %d \n", l[2](big.NewInt(2))) // Should be 0
	fmt.Printf(" - l_{2}(xCoords[1] == 4) = %d \n", l[2](big.NewInt(4))) // Should be 0
	fmt.Printf(" - l_{2}(xCoords[2] == 6) = %d \n", l[2](big.NewInt(6))) // Should be 1
	fmt.Printf(" - l_{2}(xCoords[2] == 8) = %d \n", l[2](big.NewInt(8))) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(8))

	l = append(l, BasisPolynomial(order, 3, xCoords...))

	fmt.Printf(" - l_{3}(xCoords[0] == 2) = %d \n", l[3](big.NewInt(2))) // Should be 0
	fmt.Printf(" - l_{3}(xCoords[1] == 4) = %d \n", l[3](big.NewInt(4))) // Should be 0
	fmt.Printf(" - l_{3}(xCoords[2] == 6) = %d \n", l[3](big.NewInt(6))) // Should be 1
	fmt.Printf(" - l_{3}(xCoords[2] == 8) = %d \n", l[3](big.NewInt(8))) // Should be 0

	fmt.Printf(" - l_{3}(xCoords[2] == 6) = %d \n", new(big.Int).Mod(l[3](big.NewInt(6)), order)) // Should be 1

	var err error

	var g1 *bn256.G1
	if _, g1, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var g2 = new(bn256.G1).ScalarMult(g1, l[3](big.NewInt(6)))

	fmt.Printf(" Is Equal        %t \n", bytes.Equal(g1.Marshal(), g2.Marshal()))
}

func TestExponentiation1(t *testing.T) {

	var order = bn256.Order.Set(big.NewInt(11))

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

	l = append(l, BasisPolynomial(order, 0, xCoords...))

	fmt.Println(new(big.Int).Mod(l[0](big.NewInt(3)), bn256.Order))
	fmt.Println(new(big.Int).Mod(l[0](big.NewInt(4)), bn256.Order))
	fmt.Println(new(big.Int).Mod(l[0](big.NewInt(6)), bn256.Order))
	fmt.Println(new(big.Int).Mod(l[0](big.NewInt(9)), bn256.Order))

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(3))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(9))

	l = append(l, BasisPolynomial(order, 2, xCoords...))

	fmt.Println(new(big.Int).Mod(l[1](big.NewInt(3)), bn256.Order))
	fmt.Println(new(big.Int).Mod(l[1](big.NewInt(4)), bn256.Order))
	fmt.Println(new(big.Int).Mod(l[1](big.NewInt(6)), bn256.Order))
	fmt.Println(new(big.Int).Mod(l[1](big.NewInt(9)), bn256.Order))

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(3))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(9))

	l = append(l, BasisPolynomial(order, 3, xCoords...))

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

func TestBN256Negation(t *testing.T) {

	var xCoords []*big.Int
	// var l []func(*big.Int) *big.Int

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(8))

	// l = append(l, BasisPolynomial(0, xCoords...))

	// fmt.Printf(" - l_{0}(xCoords[0] == 2) = %d \n", l[0](big.NewInt(2))) // Should be 1
	// fmt.Printf(" - l_{0}(xCoords[1] == 4) = %d \n", l[0](big.NewInt(4))) // Should be 0
	// fmt.Printf(" - l_{0}(xCoords[2] == 6) = %d \n", l[0](big.NewInt(6))) // Should be 0
	// fmt.Printf(" - l_{0}(xCoords[2] == 8) = %d \n", l[0](big.NewInt(6))) // Should be 0

	// var empty [][]byte

	// // An empty slice should fail
	// _, err := Transpose(empty)
	// if err == nil {
	// 	t.Fail()
	// }

	// blocks := make([][]byte, 2)
	// for i := range blocks {
	// 	blocks[i] = []byte{1, 2}
	// }

	// transposed, err := Transpose(blocks)
	// if err != nil {
	// 	t.Fail()
	// }
	// if len(transposed) != len(blocks[0]) {
	// 	t.Fail()
	// }
	// if len(transposed[0]) != len(blocks) {
	// 	t.Fail()
	// }
	// if transposed[0][0] != 1 {
	// 	t.Fail()
	// }
	// if transposed[1][1] != 2 {
	// 	t.Fail()
	// }
}

func TestBN256Pairing(*testing.T) {

	var err error

	var term1, term2, term3, h, t *big.Int

	var eV *bn256.G1
	var eY *bn256.G2
	var eW *bn256.G2
	var eT *bn256.G1
	var eH *bn256.G2

	var res1 *bn256.GT
	var res2 *bn256.GT

	var order = bn256.Order.Set(big.NewInt(11))
	fmt.Printf(" order: %d \n", order)

	term1 = big.NewInt(4)
	term2 = big.NewInt(3)
	term3 = big.NewInt(2)

	t = big.NewInt(2)
	h = big.NewInt(5)

	var g1 *bn256.G1
	if _, g1, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var g2 *bn256.G2
	if _, g2, err = bn256.RandomG2(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	eV = new(bn256.G1).ScalarMult(g1, term1)
	eW = new(bn256.G2).ScalarMult(g2, term2)
	eY = new(bn256.G2).ScalarMult(g2, term3)

	eT = new(bn256.G1).ScalarMult(g1, t)
	eH = new(bn256.G2).ScalarMult(g2, h)

	res1 = new(bn256.GT).Add(bn256.Pair(eV, eW), new(bn256.GT).Neg(bn256.Pair(g1, eY)))
	res2 = bn256.Pair(eT, eH)

	// Clearly: 4 * 3 - 2 = 10 = 2 * 5 (mod 11)

	fmt.Println(bytes.Equal(res1.Marshal(), res2.Marshal()))
}
