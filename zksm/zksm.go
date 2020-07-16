package zksm

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/cloudflare/bn256"
)

// E1SM provides a pedersen commitment, generated a zero-knowledge proof of set
// membership, and verifies it.
func E1SM() bool {

	var err error

	var order = bn256.Order

	// Commitment

	var gamma *big.Int
	if gamma, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("error generating field element %v \n", err)
	}

	var h *bn256.G1
	if _, h, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("error generating group element %v \n", err)
	}

	var x *big.Int
	if x, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("error generating group element %v \n", err)
	}

	var g1 *bn256.G1
	if _, g1, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("error generating group element %v \n", err)
	}

	var g2 *bn256.G2
	if _, g2, err = bn256.RandomG2(rand.Reader); err != nil {
		fmt.Printf("error generating group element %v \n", err)
	}

	var y = new(bn256.G2).ScalarMult(g2, x)

	var delta int64 = 15
	var C = new(bn256.G1).Add(
		new(bn256.G1).ScalarMult(g1, big.NewInt(delta)),
		new(bn256.G1).ScalarMult(h, gamma),
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
			sigs[elem] = new(bn256.G1).ScalarMult(g1, expo)
		}
	}

	// For building applications this data is ideal for submitting to an
	// immutable log such as a blockchain. This is because the validity of the
	// proofs rely on a sequence of data transfer where the commitment needs
	// to be supplied before the proof is generated.

	// Also, the setup step can be considered a configuration for the
	//  distributed system, and with a blockchain can be locked into a
	//  particular time slot. This would analogous to how the Paxos
	//  specification describes system configuration:
	//  https://youtu.be/JEpsBg0AO6o?t=3521

	// Prover

	var tau *big.Int
	if tau, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("error generating field element %v \n", err)
	}

	var V = new(bn256.G1).ScalarMult(sigs[delta], tau)

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

	var a = new(bn256.GT).Add(
		new(bn256.GT).ScalarMult(bn256.Pair(V, g2), new(big.Int).Sub(order, s)),
		new(bn256.GT).ScalarMult(bn256.Pair(g1, g2), t),
	)

	var D = new(bn256.G1).Add(
		new(bn256.G1).ScalarMult(g1, s),
		new(bn256.G1).ScalarMult(h, m),
	)

	// Verifier

	var c *big.Int
	if c, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("error generating field element %v \n", err)
	}

	// var zTau = new(big.Int).Sub(t, new(big.Int).Mul(tau, c))
	var zTau = new(big.Int).Add(t, new(big.Int).Sub(order, new(big.Int).Mul(tau, c)))
	// var zGamma = new(big.Int).Sub(m, new(big.Int).Mul(gamma, c))
	var zGamma = new(big.Int).Add(m, new(big.Int).Sub(order, new(big.Int).Mul(gamma, c)))
	// var zDelta = new(big.Int).Sub(s, new(big.Int).Mul(big.NewInt(delta), c))
	var zDelta = new(big.Int).Add(s, new(big.Int).Sub(order, new(big.Int).Mul(big.NewInt(delta), c)))

	c = new(big.Int).Mod(c, order)
	zGamma = new(big.Int).Mod(zGamma, order)
	zDelta = new(big.Int).Mod(zDelta, order)

	var left = new(bn256.G1).Add(
		new(bn256.G1).ScalarMult(C, c),
		new(bn256.G1).Add(
			new(bn256.G1).ScalarMult(h, zGamma),
			new(bn256.G1).ScalarMult(g1, zDelta),
		),
	)

	// fmt.Println(
	// 	new(big.Int).Mod(
	// 		new(big.Int).Sub(t,
	// 			new(big.Int).Mul(
	// 				new(big.Int).Mul(s, tau),
	// 				new(big.Int).ModInverse(new(big.Int).Add(x, big.NewInt(delta)), order),
	// 			),
	// 		),
	// 		order,
	// 	),
	// )

	// fmt.Println(
	// 	new(big.Int).Mod(
	// 		new(big.Int).Add(
	// 			new(big.Int).Sub(
	// 				new(big.Int).Mul(c,
	// 					new(big.Int).Mul(tau,
	// 						new(big.Int).Mul(x, new(big.Int).ModInverse(new(big.Int).Add(x, big.NewInt(delta)), order)),
	// 					),
	// 				),
	// 				new(big.Int).Mul(tau,
	// 					new(big.Int).Mul(zDelta, new(big.Int).ModInverse(new(big.Int).Add(x, big.NewInt(delta)), order)),
	// 				),
	// 			),
	// 			zTau,
	// 		),
	// 		order,
	// 	),
	// )

	c = new(big.Int).Mod(c, order)
	zDelta = new(big.Int).Mod(zDelta, order)
	zTau = new(big.Int).Mod(zTau, order)

	var right = new(bn256.GT).Add(
		new(bn256.GT).ScalarMult(bn256.Pair(V, y), c),
		new(bn256.GT).Add(
			new(bn256.GT).ScalarMult(bn256.Pair(V, g2), new(big.Int).Sub(order, zDelta)),
			new(bn256.GT).ScalarMult(bn256.Pair(g1, g2), zTau),
		),
	)

	return bytes.Equal(D.Marshal(), left.Marshal()) && bytes.Equal(a.Marshal(), right.Marshal())
}
