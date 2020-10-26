package sm

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/cloudflare/bn256"
)

// HPrime is the prime represntative for all members in the set.
var HPrime = map[string]int{
	"B": 3, "E": 17, "K": 31, "M": 53,
}

// E1ACCUM computes an RSA accumulator and evaluates one of the elements as a
// check of membership.
func E1ACCUM(order *big.Int) bool {

	var err error

	var p, q = big.NewInt(11), big.NewInt(17) // The two smallest strong primes

	var N = new(big.Int).Mul(p, q)

	var x *big.Int
	if x, err = rand.Int(rand.Reader, N); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	// Quadratic residue of order the RSA modulus N
	var g = new(big.Int).Exp(x, big.NewInt(2), N)

	var A17 = new(big.Int).Exp(
		g,
		new(big.Int).Mul(
			new(big.Int).Mul(
				big.NewInt(1),
				big.NewInt(3),
			),
			new(big.Int).Mul(
				big.NewInt(31),
				big.NewInt(53),
			),
		),
		N,
	)

	var left = new(big.Int).Exp(
		A17, big.NewInt(17), N,
	)

	var right = new(big.Int).Exp(
		g,
		new(big.Int).Mul(
			new(big.Int).Mul(
				big.NewInt(3),
				big.NewInt(17),
			),
			new(big.Int).Mul(
				big.NewInt(31),
				big.NewInt(53),
			),
		),
		N,
	)

	return bytes.Equal(left.Bytes(), right.Bytes())
}

// E2ACCUM computes a bilinear-map accumulator and evaluates one of the elements
//  as a check of membership.
func E2ACCUM(order *big.Int) bool {

	var err error

	var s *big.Int
	if s, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var g1 *bn256.G1
	if _, g1, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var g2 *bn256.G2
	if _, g2, err = bn256.RandomG2(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var A17 = new(bn256.G2).ScalarMult(
		g2,
		new(big.Int).Mul(
			new(big.Int).Mul(
				new(big.Int).Add(s, big.NewInt(3)),
				new(big.Int).Add(big.NewInt(0), big.NewInt(1)),
			),
			new(big.Int).Mul(
				new(big.Int).Add(s, big.NewInt(31)),
				new(big.Int).Add(s, big.NewInt(53)),
			),
		),
	)

	var left = bn256.Pair(
		new(bn256.G1).ScalarMult(
			g1,
			new(big.Int).Add(s, big.NewInt(17)),
		),
		A17,
	)

	var right = bn256.Pair(
		g1,
		new(bn256.G2).ScalarMult(
			g2,
			new(big.Int).Mul(
				new(big.Int).Mul(
					new(big.Int).Add(s, big.NewInt(3)),
					new(big.Int).Add(s, big.NewInt(17)),
				),
				new(big.Int).Mul(
					new(big.Int).Add(s, big.NewInt(31)),
					new(big.Int).Add(s, big.NewInt(53)),
				),
			),
		),
	)

	return bytes.Equal(left.Marshal(), right.Marshal())
}
