package zksm

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/cloudflare/bn256"
)

// E1QAP provided a pedersen commitment, generated a proof and verifies it.
func E1SM() bool {

	var err error

	var order = bn256.Order

	// Commitment

	var gamma *big.Int
	if gamma, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("error generating field element %v \n", err)
	}

	var delta *big.Int
	if delta, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("error generating field element %v \n", err)
	}

	var h *bn256.G2
	if _, h, err = bn256.RandomG2(rand.Reader); err != nil {
		fmt.Printf("error generating group element %v \n", err)
	}

	var x *big.Int
	var g *bn256.G2

	if x, _, err = bn256.RandomG2(rand.Reader); err != nil {
		fmt.Printf("error generating group element %v \n", err)
	}

	if _, g, err = bn256.RandomG2(rand.Reader); err != nil {
		fmt.Printf("error generating group element %v \n", err)
	}

	var y = new(bn256.G2).ScalarMult(g, x)

	var C = new(bn256.G2).Add(
		new(bn256.G2).ScalarMult(g, delta),
		new(bn256.G2).ScalarMult(h, gamma),
	)

	var sigs = make(map[int64]*bn256.G1)

	{ // Trusted Setup

		var s = []int64{
			0, 11, 13, 14, 2,
			10, 15, 16, 17, 3,
			4, 18, 19, 21, 22,
			23, 25, 5, 6, 8,
			9, 26, 27, 28, 30,
		}

		var expo *big.Int

		var elem int64
		for _, elem = range s {
			expo = new(big.Int).ModInverse(new(big.Int).Add(big.NewInt(elem), x), order)
			sigs[elem] = new(bn256.G1).ScalarMult(g, expo)
		}
	}

	// Prover

	var tau *big.Int
	if tau, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("error generating field element %v \n", err)
	}

	var V = new(bn256.G1).ScalarMult(sigs[15], tau)

	var s *big.Int
	if s, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("error generating field element %v \n", err)
	}

	var t *big.Int
	if t, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("error generating field element %v \n", err)
	}

	var m *big.Int
	if m, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("error generating field element %v \n", err)
	}

	var gRaw interface{}
	gRaw = g.Marshal()

	var g0 *bn256.G1
	g0, _ = gRaw.(*bn256.G1)

	var a = new(bn256.GT).Add(
		new(bn256.GT).ScalarMult(bn256.Pair(V, g), new(big.Int).Sub(order, s)),
		new(bn256.GT).ScalarMult(bn256.Pair(g0, g), t),
	)

	var D = new(bn256.GT).Add(
		new(bn256.GT).ScalarMult(s, g),
		new(bn256.GT).ScalarMult(m, h),
	)

	// Verifier

	var c *big.Int
	if c, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("error generating field element %v \n", err)
	}

	var zTau = new(big.Int).Sub(t, new(big.Int).Mul(tau, c))
	var zDelta = new(big.Int).Sub(s, new(big.Int).Mul(delta, c))
	var zGamma = new(big.Int).Sub(m, new(big.Int).Mul(gamma, c))

	var left = new(bn256.GT).Add(
		new(bn256.GT).ScalarMult(c, C),
		new(bn256.GT).Add(
			new(bn256.GT).ScalarMult(zGamma, h),
			new(bn256.GT).ScalarMult(zDelta, g),
		),
	)

}
