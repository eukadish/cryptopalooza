package sm

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/cloudflare/bn256"
	"gonum.org/v1/gonum/mat"
)

// HPrime is the prime representatives of all members of the set.
var HPrime = map[string]int{
	"B": 3, "E": 17, "K": 31, "M": 53,
}

// E1ACCUM computes an RSA accumulator and evaluates one of the elements as a
// check of membership.
func E1ACCUM(order *big.Int) bool {

	var err error

	// The two smallest strong primes
	var p, q = big.NewInt(11), big.NewInt(17)

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

// E3ACCUM computes an RSA accumulator and dynamically calculates the hash to
// prime functions for each element of the set.
func E3ACCUM(order *big.Int) bool {

	var err error

	var e = []*struct {
		val   *big.Int
		nonce int
	}{
		{val: big.NewInt(66)},
		{val: big.NewInt(69)},
		{val: big.NewInt(75)},
		{val: big.NewInt(77)},
	}

	// var r [4]*big.Int
	var r = make([]*big.Int, len(e))

	var index int
	var elem *struct {
		val   *big.Int
		nonce int
	}

	for index, elem = range e {

		var ct int

		var h = sha256.New()
		_, _ = h.Write(elem.val.Bytes())

		var v = new(big.Int)
		_ = v.SetBytes(h.Sum(nil))

		ct++

		for !v.ProbablyPrime(0) {

			h.Write(v.Bytes())
			v.SetBytes(h.Sum(nil))

			ct++
		}

		elem.nonce = ct
	}

	var p *big.Int
	if p, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var q *big.Int
	if q, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var N = new(big.Int).Mul(p, q)

	var x *big.Int
	if x, err = rand.Int(rand.Reader, N); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	// Quadratic residue of order the RSA modulus N
	var g = new(big.Int).Exp(x, big.NewInt(2), N)

	for index, elem = range e {

		var ct int

		var v = elem.val
		var n = elem.nonce

		var h = sha256.New()

		h.Write(v.Bytes())
		v.SetBytes(h.Sum(nil))

		ct++

		for ct < n {
			h.Write(v.Bytes())
			v.SetBytes(h.Sum(nil))

			ct++
		}

		r[index] = v
	}

	var A17 = new(big.Int).Exp(
		g,
		new(big.Int).Mul(
			new(big.Int).Mul(
				big.NewInt(1),
				r[0],
			),
			new(big.Int).Mul(
				r[2],
				r[3],
			),
		),
		N,
	)

	var left = new(big.Int).Exp(
		A17, r[1], N,
	)

	var right = new(big.Int).Exp(
		g,
		new(big.Int).Mul(
			new(big.Int).Mul(
				r[0],
				r[1],
			),
			new(big.Int).Mul(
				r[2],
				r[3],
			),
		),
		N,
	)

	return bytes.Equal(left.Bytes(), right.Bytes())
}

// E4ACCUM computes an RSA accumulator and dynamically calculates a subset for
// the two-universal family of functions for the set.
func E4ACCUM(order *big.Int) bool {

	var err error

	// *  3 = 0 0 0 0 0 0 1 1  * 66 - B = 0 1 0 0 0 0 1 0
	// * 17 = 0 0 0 1 0 0 0 1  * 69 - E = 0 1 0 0 0 1 0 1
	// * 31 = 0 0 0 1 1 1 1 1  * 75 - K = 0 1 0 0 1 0 1 1
	// * 53 = 0 0 1 1 0 1 0 1  * 77 - M = 0 1 0 0 1 1 0 1

	var e = [][]float64{
		{0, 1, 0, 0, 0, 0, 1, 0}, // 66
		{0, 1, 0, 0, 0, 1, 0, 1}, // 69
		{0, 1, 0, 0, 1, 0, 1, 1}, // 75
		{0, 1, 0, 0, 1, 1, 0, 1}, // 77
	}

	var r = [][]float64{
		{0, 0, 0, 0, 0, 0, 1, 1}, //  3
		{0, 0, 0, 1, 0, 0, 0, 1}, // 17
		{0, 0, 0, 1, 1, 1, 1, 1}, // 31
		{0, 0, 1, 1, 0, 1, 0, 1}, // 53
	}

	var h = [][]float64{
		make([]float64, 8*8),
		make([]float64, 8*8),
		make([]float64, 8*8),
		make([]float64, 8*8),
	}

	var index int
	var elem []float64

	// NOTE: Instead of using predetermined primes, the matrices can be fully
	// randomized and checked if mapping to some prime in order to derive the
	// two-universal functions.

	for index, elem = range e {

		// NOTE: The iteration ends when some two-universal map is found, but
		// this loop could be used to compute the entire family of
		// two-universal functions.

		var v = mat.NewVecDense(8, nil)
		for !mat.Equal(v, mat.NewVecDense(8, r[index])) {
			var item float64

			var i int
			for i, item = range r[index] {

				var k int
				for k = i * 8; k < (i+1)*8; k++ {

					if item == 0 {
						if elem[k%8] == 0 {

							var b *big.Int
							if b, err = rand.Int(rand.Reader, big.NewInt(2)); err != nil {
								fmt.Printf("parameter generation %v", err)
							}

							h[index][k] = float64(b.Int64())

							continue
						}

						h[index][k] = 0

						continue
					}

					var b *big.Int
					if b, err = rand.Int(rand.Reader, big.NewInt(2)); err != nil {
						fmt.Printf("parameter generation %v", err)
					}

					h[index][k] = float64(b.Int64())
				}
			}

			var b = mat.NewVecDense(8, elem)
			var a = mat.NewDense(8, 8, h[index])

			v.MulVec(a, b)
		}
	}

	// var hp []float64
	// for _, hp = range h {
	// 	fmt.Printf(
	// 		" \n A = \n    %v \n", mat.Formatted(mat.NewDense(8, 8, hp), mat.Prefix("    "), mat.Squeeze()),
	// 	)
	// }

	var p, q = big.NewInt(11), big.NewInt(17) // The two smallest strong primes

	var N = new(big.Int).Mul(p, q)

	var x *big.Int
	if x, err = rand.Int(rand.Reader, N); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	// Quadratic residue of order the RSA modulus N
	var g = new(big.Int).Exp(x, big.NewInt(2), N)

	var s = mat.NewVecDense(
		8, []float64{128, 64, 32, 16, 8, 4, 2, 1},
	)

	var t = []*mat.VecDense{
		mat.NewVecDense(8, nil),
		mat.NewVecDense(8, nil),
		mat.NewVecDense(8, nil),
		mat.NewVecDense(8, nil),
	}

	t[0].MulVec(mat.NewDense(8, 8, h[0]), mat.NewVecDense(8, e[0]))
	t[1].MulVec(mat.NewDense(8, 8, h[1]), mat.NewVecDense(8, e[1]))
	t[2].MulVec(mat.NewDense(8, 8, h[2]), mat.NewVecDense(8, e[2]))
	t[3].MulVec(mat.NewDense(8, 8, h[3]), mat.NewVecDense(8, e[3]))

	var A17 = new(big.Int).Exp(
		g,
		new(big.Int).Mul(
			new(big.Int).Mul(
				big.NewInt(int64(1)),
				big.NewInt(int64(mat.Dot(s, t[0]))),
			),
			new(big.Int).Mul(
				big.NewInt(int64(mat.Dot(s, t[2]))),
				big.NewInt(int64(mat.Dot(s, t[3]))),
			),
		),
		N,
	)

	var left = new(big.Int).Exp(
		A17, big.NewInt(int64(mat.Dot(s, t[1]))), N,
	)

	var right = new(big.Int).Exp(
		g,
		new(big.Int).Mul(
			new(big.Int).Mul(
				big.NewInt(int64(mat.Dot(s, mat.NewVecDense(8, r[0])))),
				big.NewInt(int64(mat.Dot(s, mat.NewVecDense(8, r[1])))),
			),
			new(big.Int).Mul(
				big.NewInt(int64(mat.Dot(s, mat.NewVecDense(8, r[2])))),
				big.NewInt(int64(mat.Dot(s, mat.NewVecDense(8, r[3])))),
			),
		),
		N,
	)

	return bytes.Equal(left.Bytes(), right.Bytes())
}
