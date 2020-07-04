package qap

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"

	"testing"

	"github.com/cloudflare/bn256"
)

func TestBasisPolynomial(t *testing.T) {

	var order = bn256.Order.Set(big.NewInt(23))

	var xCoords []*big.Int
	var l []func(*big.Int) *big.Int

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(8))

	l = append(l, BasisPolynomial(order, 0, xCoords...))

	fmt.Printf(" - l_{0}(xCoords[0] == 2) = %d \n", new(big.Int).Mod(l[0](big.NewInt(2)), order)) // Should be 1
	fmt.Printf(" - l_{0}(xCoords[1] == 4) = %d \n", new(big.Int).Mod(l[0](big.NewInt(4)), order)) // Should be 0
	fmt.Printf(" - l_{0}(xCoords[2] == 6) = %d \n", new(big.Int).Mod(l[0](big.NewInt(6)), order)) // Should be 0
	fmt.Printf(" - l_{0}(xCoords[3] == 8) = %d \n", new(big.Int).Mod(l[0](big.NewInt(8)), order)) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(8))

	l = append(l, BasisPolynomial(order, 1, xCoords...))

	fmt.Printf(" - l_{1}(xCoords[0] == 2) = %d \n", new(big.Int).Mod(l[1](big.NewInt(2)), order)) // Should be 0
	fmt.Printf(" - l_{1}(xCoords[1] == 4) = %d \n", new(big.Int).Mod(l[1](big.NewInt(4)), order)) // Should be 1
	fmt.Printf(" - l_{1}(xCoords[2] == 6) = %d \n", new(big.Int).Mod(l[1](big.NewInt(6)), order)) // Should be 0
	fmt.Printf(" - l_{1}(xCoords[3] == 8) = %d \n", new(big.Int).Mod(l[1](big.NewInt(8)), order)) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(8))

	l = append(l, BasisPolynomial(order, 2, xCoords...))

	fmt.Printf(" - l_{2}(xCoords[0] == 2) = %d \n", new(big.Int).Mod(l[2](big.NewInt(2)), order)) // Should be 0
	fmt.Printf(" - l_{2}(xCoords[1] == 4) = %d \n", new(big.Int).Mod(l[2](big.NewInt(4)), order)) // Should be 0
	fmt.Printf(" - l_{2}(xCoords[2] == 6) = %d \n", new(big.Int).Mod(l[2](big.NewInt(6)), order)) // Should be 1
	fmt.Printf(" - l_{2}(xCoords[3] == 8) = %d \n", new(big.Int).Mod(l[2](big.NewInt(8)), order)) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(4))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(8))

	l = append(l, BasisPolynomial(order, 3, xCoords...))

	fmt.Printf(" - l_{3}(xCoords[0] == 2) = %d \n", new(big.Int).Mod(l[3](big.NewInt(2)), order)) // Should be 0
	fmt.Printf(" - l_{3}(xCoords[1] == 4) = %d \n", new(big.Int).Mod(l[3](big.NewInt(4)), order)) // Should be 0
	fmt.Printf(" - l_{3}(xCoords[2] == 6) = %d \n", new(big.Int).Mod(l[3](big.NewInt(6)), order)) // Should be 0
	fmt.Printf(" - l_{3}(xCoords[3] == 8) = %d \n", new(big.Int).Mod(l[3](big.NewInt(8)), order)) // Should be 1
}

func TestInterpolation(t *testing.T) {

	var order = bn256.Order.Set(big.NewInt(11))

	var xCoords []*big.Int
	var l []func(*big.Int) *big.Int

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(5))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(9))

	l = append(l, BasisPolynomial(order, 0, xCoords...))

	fmt.Printf(" - l_{0}(xCoords[0] == 2) = %d \n", new(big.Int).Mod(l[0](big.NewInt(2)), order)) // Should be 1
	fmt.Printf(" - l_{0}(xCoords[1] == 5) = %d \n", new(big.Int).Mod(l[0](big.NewInt(5)), order)) // Should be 0
	fmt.Printf(" - l_{0}(xCoords[2] == 6) = %d \n", new(big.Int).Mod(l[0](big.NewInt(6)), order)) // Should be 0
	fmt.Printf(" - l_{0}(xCoords[2] == 9) = %d \n", new(big.Int).Mod(l[0](big.NewInt(9)), order)) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(5))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(9))

	l = append(l, BasisPolynomial(order, 1, xCoords...))

	fmt.Printf(" - l_{1}(xCoords[0] == 2) = %d \n", new(big.Int).Mod(l[0](big.NewInt(2)), order)) // Should be 0
	fmt.Printf(" - l_{1}(xCoords[1] == 5) = %d \n", new(big.Int).Mod(l[0](big.NewInt(5)), order)) // Should be 1
	fmt.Printf(" - l_{1}(xCoords[2] == 6) = %d \n", new(big.Int).Mod(l[0](big.NewInt(6)), order)) // Should be 0
	fmt.Printf(" - l_{1}(xCoords[2] == 9) = %d \n", new(big.Int).Mod(l[0](big.NewInt(9)), order)) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(5))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(9))

	l = append(l, BasisPolynomial(order, 2, xCoords...))

	fmt.Printf(" - l_{2}(xCoords[0] == 2) = %d \n", new(big.Int).Mod(l[0](big.NewInt(2)), order)) // Should be 0
	fmt.Printf(" - l_{2}(xCoords[1] == 5) = %d \n", new(big.Int).Mod(l[0](big.NewInt(5)), order)) // Should be 0
	fmt.Printf(" - l_{2}(xCoords[2] == 6) = %d \n", new(big.Int).Mod(l[0](big.NewInt(6)), order)) // Should be 1
	fmt.Printf(" - l_{2}(xCoords[2] == 9) = %d \n", new(big.Int).Mod(l[0](big.NewInt(9)), order)) // Should be 0

	xCoords = []*big.Int{}

	xCoords = append(xCoords, big.NewInt(2))
	xCoords = append(xCoords, big.NewInt(5))
	xCoords = append(xCoords, big.NewInt(6))
	xCoords = append(xCoords, big.NewInt(9))

	l = append(l, BasisPolynomial(order, 3, xCoords...))

	fmt.Printf(" - l_{3}(xCoords[0] == 2) = %d \n", new(big.Int).Mod(l[0](big.NewInt(2)), order)) // Should be 0
	fmt.Printf(" - l_{3}(xCoords[1] == 5) = %d \n", new(big.Int).Mod(l[0](big.NewInt(5)), order)) // Should be 0
	fmt.Printf(" - l_{3}(xCoords[2] == 6) = %d \n", new(big.Int).Mod(l[0](big.NewInt(6)), order)) // Should be 0
	fmt.Printf(" - l_{3}(xCoords[2] == 9) = %d \n", new(big.Int).Mod(l[0](big.NewInt(9)), order)) // Should be 1

	var eval = Interpolate(big.NewInt(5), []int64{2, 7, 1, 3}, l...)

	fmt.Printf(" - eval = %d \n", new(big.Int).Mod(eval, order)) // Should be 7
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

	fmt.Println(bytes.Equal(res1.Marshal(), res2.Marshal())) // Should be true

	term1 = big.NewInt(10)
	term2 = big.NewInt(9)
	term3 = big.NewInt(1)

	t = big.NewInt(4)
	h = big.NewInt(3)

	eV = new(bn256.G1).ScalarMult(g1, term1)
	// eV = new(bn256.G1).ScalarMult(g1, new(big.Int).Mod(term1, order))

	eW = new(bn256.G2).ScalarMult(g2, term2)
	// eW = new(bn256.G2).ScalarMult(g2, new(big.Int).Mod(term2, bn256.Order))

	eY = new(bn256.G2).ScalarMult(g2, term3)
	// eY = new(bn256.G2).Neg(new(bn256.G2).ScalarMult(g2, term3))
	// eY = new(bn256.G2).ScalarMult(g2, new(big.Int).Neg(term3))
	// eY = new(bn256.G2).ScalarMult(g2, new(big.Int).Mod(new(big.Int).Neg(term3), bn256.Order))

	eT = new(bn256.G1).ScalarMult(g1, t)
	// eT = new(bn256.G1).ScalarMult(g1, new(big.Int).Mod(t, bn256.Order))

	eH = new(bn256.G2).ScalarMult(g2, h)
	// eH = new(bn256.G2).ScalarMult(g2, new(big.Int).Mod(h, bn256.Order))

	res1 = new(bn256.GT).Add(bn256.Pair(eV, eW), new(bn256.GT).Neg(bn256.Pair(g1, eY)))
	// res1 = new(bn256.GT).Add(bn256.Pair(eV, eW), bn256.Pair(g1, eY))

	res2 = bn256.Pair(eT, eH)

	// Clearly: 10 * 9 - 1 = 89 = > 1 < = 4 * 3 = 12 (mod 11)

	fmt.Println(bytes.Equal(res1.Marshal(), res2.Marshal())) // Should be true
}

func TestBN256Negation(*testing.T) {

	var err error

	var term1, term2, term3, term4 /* h, t */ *big.Int

	var eV *bn256.G1
	var eY *bn256.G2
	var eW *bn256.G2
	// var eT *bn256.G1
	// var eH *bn256.G2

	var res1 *bn256.GT
	var res2 *bn256.GT

	var order = bn256.Order.Set(big.NewInt(11))
	fmt.Printf(" order: %d \n", order)

	// term1 = big.NewInt(4)
	// term2 = big.NewInt(3)
	// term3 = big.NewInt(2)

	// t = big.NewInt(2)
	// h = big.NewInt(5)

	// var g1 *bn256.G1
	// if _, g1, err = bn256.RandomG1(rand.Reader); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	// var g2 *bn256.G2
	// if _, g2, err = bn256.RandomG2(rand.Reader); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	// eV = new(bn256.G1).ScalarMult(g1, term1)
	// eW = new(bn256.G2).ScalarMult(g2, term2)
	// eY = new(bn256.G2).ScalarMult(g2, term3)

	// eT = new(bn256.G1).ScalarMult(g1, t)
	// eH = new(bn256.G2).ScalarMult(g2, h)

	// res1 = new(bn256.GT).Add(bn256.Pair(eV, eW), new(bn256.GT).Neg(bn256.Pair(g1, eY)))
	// res2 = bn256.Pair(eT, eH)

	// // Clearly: 4 * 3 - 2 = 10 = 2 * 5 (mod 11)

	// fmt.Println(bytes.Equal(res1.Marshal(), res2.Marshal())) // Should be true

	// term1 = big.NewInt(10)
	// term2 = big.NewInt(9)
	// term3 = big.NewInt(1)

	// t = big.NewInt(4)
	// h = big.NewInt(3)

	// eV = new(bn256.G1).ScalarMult(g1, term1)
	// // eV = new(bn256.G1).ScalarMult(g1, new(big.Int).Mod(term1, order))

	// eW = new(bn256.G2).ScalarMult(g2, term2)
	// // eW = new(bn256.G2).ScalarMult(g2, new(big.Int).Mod(term2, bn256.Order))

	// eY = new(bn256.G2).ScalarMult(g2, term3)
	// // eY = new(bn256.G2).Neg(new(bn256.G2).ScalarMult(g2, term3))
	// // eY = new(bn256.G2).ScalarMult(g2, new(big.Int).Neg(term3))
	// // eY = new(bn256.G2).ScalarMult(g2, new(big.Int).Mod(new(big.Int).Neg(term3), bn256.Order))

	// eT = new(bn256.G1).ScalarMult(g1, t)
	// // eT = new(bn256.G1).ScalarMult(g1, new(big.Int).Mod(t, bn256.Order))

	// eH = new(bn256.G2).ScalarMult(g2, h)
	// // eH = new(bn256.G2).ScalarMult(g2, new(big.Int).Mod(h, bn256.Order))

	// res1 = new(bn256.GT).Add(bn256.Pair(eV, eW), new(bn256.GT).Neg(bn256.Pair(g1, eY)))
	// // res1 = new(bn256.GT).Add(bn256.Pair(eV, eW), bn256.Pair(g1, eY))

	// res2 = bn256.Pair(eT, eH)

	// // Clearly: 10 * 9 - 1 = 89 = > 1 < = 4 * 3 = 12 (mod 11)

	// fmt.Println(bytes.Equal(res1.Marshal(), res2.Marshal())) // Should be true

	term1 = big.NewInt(2)
	term2 = big.NewInt(3)
	term3 = big.NewInt(6)

	// t = big.NewInt(2)
	// h = big.NewInt(5)

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

	// eT = new(bn256.G1).ScalarMult(g1, t)
	// eH = new(bn256.G2).ScalarMult(g2, h)

	// TODO: use the negative int here

	res1 = new(bn256.GT).Add(bn256.Pair(eV, eW), new(bn256.GT).Neg(bn256.Pair(g1, eY)))

	fmt.Printf("e(E(v) + E(w)) + e(G1 + E(y)) = %v \n", res1)

	// res2 = bn256.Pair(eT, eH)

	term4 = big.NewInt(0)

	res2 = bn256.Pair(g1, new(bn256.G2).ScalarMult(g2, term4))

	fmt.Println(bytes.Equal(res1.Marshal(), res2.Marshal())) // Should be true
}

func TestBN256GroupOperator(*testing.T) {

}
