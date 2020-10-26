package sm

import (
	"bytes"
	"math/big"
	"math/rand"
	"time"
)

// HPrime is the prime represntative for all members in the set.
var HPrime = map[string]int{
	"B": 3, "E": 17, "K": 31, "M": 53,
}

// E1ACCUM computes an RSA accumulator and evaluates one of the elements as a
// check of membership.
func E1ACCUM(order *big.Int) bool {

	var p, q = big.NewInt(11), big.NewInt(17) // The two smallest strong primes

	var N = new(big.Int).Mul(p, q)
	var g = new(big.Int).Exp( // Quadratic residue of order the RSA modulus N
		new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), N), big.NewInt(2), N,
	)

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
