package examples

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/cloudflare/bn256"
)

// f(x1) = 3 * x1
//       = (3) * (x1)
//       = p1(x1) * p2(x1)
//       = 6

// p1(x1) = c_{0} + Sigma_{i = 1}^{1} c_{i} * x_{i}
//        = c_{0} + c_{1} * x_{1}
//        = 3 + 0 * x1

// p2(x1) = d_{0} + Sigma_{i = 1}^{m - 1} d_{i} * x_{i}
//        = d_{0} + d_{1} * x_{1}
//        = 0 + 1 * 2

/* # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#                                                                                                                                                           #
#    3       x1    3       a1     3       2                                                                         p(a1) * p(a1) - a2 = (3) * (a1) - a2    #
#     \     /       \     /        \     /                                                                                             = 3 * 2 - 6          #
#      \   /         \   /          \   /                                                                                              = 0                  #
#       \ /           \ /            \ /                                                                                                                    #
#        *             *              *                                                                                                                     #
#        |             |              |                                                                                                                     #
#        |             |              |                                                                                                                     #
#        |             |              |                                                                                                                     #
#      f(x1)           a2             6                                                                                                                     #
#                                                                                                                                                           #
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # */

/* # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#                                                                                                                                                            #
#                                                                                                                    p(a1) * p(a1) - a2 = (3) * (a1) - a2    #
#                                                                                                                                        = 3 * 2 - 6         #
#        x1            a1             2                                                                                                  = 0                 #
#        |             |              |                                                                                                                      #
#      3 *           3 *            3 *                                                                                                                      #
#        |             |              |                                                                                                                      #
#      f(x1)           a2             6                                                                                                                      #
#                                                                                                                                                            #
#                                                                                                                                                            #
#                                                                                                                                                            #
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # */

// v_{0}(x) = c_{0} = 3
// v_{1}(x) = c_{1} = 0
// v_{2}(x)         = 0

// w_{0}(x) = d_{0} = 0
// w_{1}(x) = d_{1} = 1
// w_{2}(x)         = 0

// y_{0}(x)         = 0
// y_{1}(x)         = 0
// y_{2}(x)         = 1

// Example1 this function creates a QAP (Quadratic Arithmetic Program) for the arithmetic expression: 3 * x = 6
func Example1() bool {

	var err error

	var g1 *bn256.G1
	if _, g1, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var g2 *bn256.G2
	if _, g2, err = bn256.RandomG2(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	// var gt *bn256.GT
	// if _, gt, err = bn256.RandomGT(rand.Reader); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	var v [3]*bn256.G1

	if v[0] = new(bn256.G1).ScalarMult(g1, big.NewInt(3)); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	if v[1] = new(bn256.G1).ScalarMult(g1, new(big.Int).Mul(big.NewInt(2), big.NewInt(0))); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	if v[2] = new(bn256.G1).ScalarMult(g1, new(big.Int).Mul(big.NewInt(6), big.NewInt(0))); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var w [3]*bn256.G2

	if w[0] = new(bn256.G2).ScalarMult(g2, big.NewInt(0)); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	if w[1] = new(bn256.G2).ScalarMult(g2, new(big.Int).Mul(big.NewInt(2), big.NewInt(1))); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	if w[2] = new(bn256.G2).ScalarMult(g2, new(big.Int).Mul(big.NewInt(6), big.NewInt(0))); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	// var y [3]*bn256.GT

	// if y[0] = new(bn256.GT).ScalarMult(gt, big.NewInt(0)); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	// if y[1] = new(bn256.GT).ScalarMult(gt, new(big.Int).Mul(big.NewInt(2), big.NewInt(0))); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	// if y[2] = new(bn256.GT).ScalarMult(gt, new(big.Int).Mul(big.NewInt(6), big.NewInt(1))); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	var y [3]*bn256.G2

	if y[0] = new(bn256.G2).ScalarMult(g2, big.NewInt(0)); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	if y[1] = new(bn256.G2).ScalarMult(g2, new(big.Int).Mul(big.NewInt(2), big.NewInt(0))); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	if y[2] = new(bn256.G2).ScalarMult(g2, new(big.Int).Mul(big.NewInt(6), big.NewInt(1))); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var left = new(bn256.G1).Add(v[0], new(bn256.G1).Add(v[1], v[2]))
	var right = new(bn256.G2).Add(w[0], new(bn256.G2).Add(w[1], w[2]))

	var product = bn256.Pair(left, right)
	var result = bn256.Pair(g1, new(bn256.G2).Add(y[0], new(bn256.G2).Add(y[1], y[2])))

	fmt.Printf(
		" (v0 + a1 * v1 + a2 * v2) * (w0 + a1 * w1 + a2 * w2) - (y0 + a1 * y1 + a2 * y2) = (%s) * (%s) - %s \n",
		left.String()[0:18], right.String()[0:18], result.String()[0:18],
	)

	return bytes.Equal(product.Marshal(), result.Marshal())
}

// LinearR1CS generates the Quadratic Arithmetic Program to validate arithmetic circuits in Zero Knowledge
func LinearR1CS() bool {

	return true
}
