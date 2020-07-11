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

	var left, right *bn256.GT

	var order = bn256.Order
	// var order = bn256.Order.Set(big.NewInt(11))
	// var order = bn256.Order.Set(big.NewInt(983))

	var g1 *bn256.G1
	if _, g1, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var g2 *bn256.G2
	if _, g2, err = bn256.RandomG2(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	left = new(bn256.GT).Add(
		bn256.Pair(
			new(bn256.G1).ScalarMult(g1, big.NewInt(10)),
			new(bn256.G2).ScalarMult(g2, big.NewInt(9)),
		),
		bn256.Pair(
			new(bn256.G1).ScalarMult(g1, new(big.Int).Sub(order, big.NewInt(1))),
			new(bn256.G2).ScalarMult(g2, big.NewInt(90)),
		),
	)

	right = bn256.Pair(
		g1,
		new(bn256.G2).ScalarMult(g2, big.NewInt(0)),
	)

	// 10 * 9 + (-1) * 90 == 0
	fmt.Println(bytes.Equal(left.Marshal(), right.Marshal()))

	left = new(bn256.GT).Add(
		bn256.Pair(
			new(bn256.G1).ScalarMult(g1, big.NewInt(3)),
			new(bn256.G2).ScalarMult(g2, big.NewInt(7)),
		),
		bn256.Pair(
			g1,
			// new(bn256.G2).ScalarMult(g2, big.NewInt(-16)),
			new(bn256.G2).ScalarMult(g2, new(big.Int).Sub(order, big.NewInt(16))),
		),
	)

	right = bn256.Pair(
		g1,
		new(bn256.G2).ScalarMult(g2, big.NewInt(5)),
	)

	// 3 * 7 - 16 = 5
	fmt.Println(bytes.Equal(left.Marshal(), right.Marshal()))

	left = bn256.Pair(
		new(bn256.G1).ScalarMult(g1, big.NewInt(2)),
		new(bn256.G2).ScalarMult(g2, big.NewInt(5)),
	)

	right = bn256.Pair(
		g1,
		new(bn256.G2).ScalarMult(g2, big.NewInt(10)),
	)

	// 2 * 5 == 10
	fmt.Println(bytes.Equal(left.Marshal(), right.Marshal()))

	left = new(bn256.GT).Add(
		bn256.Pair(
			new(bn256.G1).ScalarMult(g1, big.NewInt(4)),
			new(bn256.G2).ScalarMult(g2, big.NewInt(3)),
		),
		bn256.Pair(g1, new(bn256.G2).ScalarMult(g2, new(big.Int).Sub(order, big.NewInt(2)))),
	)

	right = bn256.Pair(
		new(bn256.G1).ScalarMult(g1, big.NewInt(2)),
		new(bn256.G2).ScalarMult(g2, big.NewInt(5)),
	)

	// 4 * 3 - 2 = 10 = 2 * 5
	fmt.Println(bytes.Equal(left.Marshal(), right.Marshal()))

	left = new(bn256.GT).Add(
		bn256.Pair(
			new(bn256.G1).ScalarMult(g1, big.NewInt(2)),
			new(bn256.G2).ScalarMult(g2, big.NewInt(3)),
		),
		new(bn256.GT).Neg(
			bn256.Pair(
				g1,
				new(bn256.G2).ScalarMult(g2, big.NewInt(6)),
			),
		),
	)

	right = bn256.Pair(
		g1,
		new(bn256.G2).ScalarMult(g2, big.NewInt(0)),
	)

	// 2 * 3 - 6 = 0
	fmt.Println(bytes.Equal(left.Marshal(), right.Marshal()))
}
