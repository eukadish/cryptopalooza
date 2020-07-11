package qap

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/cloudflare/bn256"

	"github.com/eugenekadish/cryptopalooza/zksnark/qap/lagrange"
)

// f(x1, x2, x3, x4) = 4 * x1 * x2 - 7 * x2 + 3 * x4
//                   =     x3      - 7 * x2 + 3 * x4
//                   = f2(x2, x3, x4)
//                   = x5

// 4 * a1 * a2 - 7 * a2 + 3 * a4 = a5

// This relation is satisfied for a1 = 3, a2 = 2, a4 = 1, and a5 = 13. The index
// 3 is skipped for convenience, because it will be used for an intermediate
// result for the QAP composition.

// f1(x1, x2) = 4 * x1 * x2
//            = (4 * x1) * x2
//            = p1(x1, x2) * p2(x1, x2)
//            = x3
//            = 24

// p1(x1, x2) = v1_{0} + v1_{1} * x_{1} + v1_{2} * x_{2} + v1_{3} * x_{3} + v1_{4} * x_{4} + v1_{5} * x_{5}
//            =   0    +     4  *  a1   +     0  *  a2   +     0  *  a3   +     0  *  a4   +     0  *  a5
//            = 4 * 3

// p2(x1, x2) = w1_{0} + w1_{1} * x_{1} + w1_{2} * x_{2} + w1_{3} * x_{3} + w1_{4} * x_{4} + w1_{5} * x_{5}
//            =   0    +     0  *  a1   +     1  *  a2   +     0  *  a3   +     0  *  a4   +     0  *  a5
//            = 1 * 2

// I_{1} = {1, 2, 3}

/* # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#                                                                                                                                                            #
#    x1       x2    a1       a2    3       2                                                           p1(a1, a2) * p2(a1, 2) - a3 = (4 * a1) * (a2) - a3    #
#      \     /        \     /       \     /                                                                                        = 4 * 3 * 2 - 24          #
#     4 *   /        4 *   /       4 *   /                                                                                         = 0                       #
#        \ /            \ /           \ /                                                                                                                    #
#         *              *             *                                                                                                                     #
#         |              |             |                                                                                                                     #
#         |              |             |                                                                                                                     #
#         |              |             |                                                                                                                     #
#     f(x1, x2)          a3            24                                                                                                                    #
#                                                                                                                                                            #
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # */

// f2(x3, x4, x5) = x3 - 7 * x2 + 3 * x4
//                = 1 * ((-7) * x2 + 1 * x3 + 3 * x4)
//                = p1(x2, x3, x4) * p2(x2, x3, x4)
//                = x5
//                = 13

// p1(x2, x3, x4) = v2_{0} + v2_{1} * x_{1} + v2_{2} * x_{2} + v2_{3} * x_{3} + v2_{4} * x_{4} + v2_{5} * x_{5}
//                =   1    +     0  *  a1   +     0  *  a2   +     0  *  a3   +     0  *  a4   +     0  *  a5
//                = 1

// p2(x3, x4, x5) = w2_{0} + w2_{1} * x_{1} + w2_{2} * x_{2} + w2_{3} * x_{3} + w2_{4} * x_{4} + w2_{5} * x_{5}
//                =   0    +     0  *  a1   +  (-7)  *  a2   +     1  *  a3   +     3  *  a4   +     0  *  a5
//                = (-7) * 2 + 1 * 24 + 3 * 1

// I_{2} = {2, 3, 4, 5}

/* # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#                                                                                                                                                            #
#          x3       x5      a3       a5     24       2                             p(a3, a4, a5) * p(a3, a4, a5) - a6 = (1) * (a3 + 3 * a4 - 7 * a5) - a6    #
#            \     /          \     /         \     /                                                                 = (1) * (24 + 3 * 1 - 7 * 2) - 13      #
#             \   * 7          \   * 7         \   * 7                                                                = 0                                    #
#              \ /              \ /             \ /                                                                                                          #
#      x4       -       a4       -       1       -                                                                                                           #
#        \     /          \     /         \     /                                                                                                            #
#       3 *   /          3 *   /         3 *   /                                                                                                             #
#          \ /              \ /             \ /                                                                                                              #
#           +                +               +                                                                                                               #
#           |                |               |                                                                                                               #
#           * 1              * 1             * 1                                                                                                             #
#           |                |               |                                                                                                               #
#    f2(x3, x4, x5)          a6              13                                                                                                              #
#                                                                                                                                                            #
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # */

// v0(r1) =   0, v0(r2) =   1 | w0(r1) =   0, w0(r2) =   0 | y0(r1) =  0, y0(r2) =  0
// v1(r1) =   4, v1(r2) =   0 | w1(r1) =   0, w1(r2) =   0 | y1(r1) =  0, y1(r2) =  0
// v2(r1) =   0, v2(r2) =   0 | w2(r1) =   1, w2(r2) =  -7 | y2(r1) =  0, y2(r2) =  0
// v3(r1) =   0, v3(r2) =   0 | w3(r1) =   0, w3(r2) =   1 | y3(r1) =  1, y3(r2) =  0
// v4(r1) =   0, v4(r2) =   0 | w4(r1) =   0, w4(r2) =   3 | y4(r1) =  0, y4(r2) =  0
// v5(r1) =   0, v5(r2) =   0 | w5(r1) =   0, w5(r2) =   0 | y5(r1) =  0, y5(r2) =  1

// Use the constraints to define interpolation polynomials for the QAP.

//                x - r1   |                                              |
// v0(x) = 1 * ----------- | w0(x) =                 0                    | y0(x) =          0
//               r2 - r1   |                                              |

//                x - r2   |                                              |
// v1(x) = 4 * ----------- | w1(x) =                 0                    | y4(x) =          0
//               r1 - r2   |                                              |

//                         |                x - r2               x - r1   |
// v2(x) =          0      | w2(x) = 1 * ----------- + (-7) * ----------- | y2(x) =          0
//                         |               r1 - r2              r2 - r1   |

//                         |                       x - r1                 |                x - r2
// v3(x) =          0      | w3(x) =        1 * -----------               | y3(x) = 1 * -----------
//                         |                      r2 - r1                 |               r1 - r2

//                         |                       x - r1                 |
// v4(x) =          0      | w4(x) =        3 * -----------               | y4(x) =          0
//                         |                      r2 - r1                 |

//                         |                                              |                x - r1
// v5(x) =          0      | w5(x) =                 0                    | y5(x) = 1 * -----------
//                         |                                              |               r2 - r1

// The assignment variables including those that are intermediate for the
// circuite are: a1 = 3, a2 = 2, a3 = 24, a4 = 1, and a5 = 13. To check the
// polynomials as defined are correct the relation can be evaluated at r1 and
// r2 to see if has the expected result of 0.

// (v_{0} + \Sigma_{k = 1}^{5} a_{k} * v_{k}(x)) (w_{0} + \Sigma_{k = 1}^{5} a_{k} * w_{k}(x)) - (y_{0} + \Sigma_{k = 1}^{5} a_{k} * y_{k}(x))
//       = (v0(x) + 3 * v1(x)) (2 * w2(x) + 24 * w3(x) + 1 * w4(x)) - (24 * y3(x) + 13 * y5(x))

// at r1 = (0 + 3 * 4) (2 * 1 + 24 * 0 + 1 * 0) - (24 * 1 + 13 * 0) = (12)(2) - 24 = 0
// at r2 = (1 + 3 * 0) (2 * (-7) + 24 * 1 + 1 * 3) - (24 * 0 + 13 * 1) = (1)(13) - 13 = 0

// E2QAP defines a QAP for the arithmetic expression, uses it to create a SNARK,
// and evaluates it.
func E2QAP() bool {

	var err error

	var order = bn256.Order

	var g1 *bn256.G1
	if _, g1, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var g2 *bn256.G2
	if _, g2, err = bn256.RandomG2(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var r1 = big.NewInt(3)
	var r2 = big.NewInt(7)

	var s = big.NewInt(27)

	var v [6]*bn256.G1
	var leftG []*big.Int

	leftG = append(
		leftG,
		Interpolate(
			s, []int64{1},
			BasisPolynomial(order, 1, []*big.Int{r1, r2}...),
		), // v0(s)
	)

	leftG[0] = new(big.Int).Mul(big.NewInt(1), leftG[0])
	v[0] = new(bn256.G1).ScalarMult(g1, leftG[0]) // E(v0(s))

	leftG = append(
		leftG,
		lagrange.Interpolate(
			s, []int64{1},
			BasisPolynomial(order, 0, []*big.Int{r1, r2}...),
		), // v1(s)
	)

	leftG[1] = new(big.Int).Mul(big.NewInt(3), leftG[1]) // a1 = 3
	v[1] = new(bn256.G1).ScalarMult(g1, leftG[1])        // E(a1 * v1(s))

	leftG = append(
		leftG,
		big.NewInt(0), // v2(s)
	)

	leftG[2] = new(big.Int).Mul(big.NewInt(2), leftG[2]) // a2 = 2
	v[2] = new(bn256.G1).ScalarMult(g1, leftG[2])        // E(a2 * v2(s))

	leftG = append(
		leftG,
		big.NewInt(0), // v3(s)
	)

	leftG[3] = new(big.Int).Mul(big.NewInt(24), leftG[3]) // a3 = 24
	v[3] = new(bn256.G1).ScalarMult(g1, leftG[3])         // E(a3 * v3(s))

	leftG = append(
		leftG,
		big.NewInt(0), // v4(s)
	)

	leftG[4] = new(big.Int).Mul(big.NewInt(1), leftG[4]) // a4 = 1
	v[4] = new(bn256.G1).ScalarMult(g1, leftG[4])        // E(a4 * v4(s))

	leftG = append(
		leftG,
		big.NewInt(0), // v5(s)
	)

	leftG[5] = new(big.Int).Mul(big.NewInt(13), leftG[5]) // a4 = 13
	v[5] = new(bn256.G1).ScalarMult(g1, leftG[5])         // E(a5 * v5(s))

	var w [6]*bn256.G2
	var rightG []*big.Int

	rightG = append(
		rightG,
		big.NewInt(0), // w0(s)
	)

	rightG[0] = new(big.Int).Mul(big.NewInt(1), rightG[0])
	w[0] = new(bn256.G2).ScalarMult(g2, rightG[0]) // E(w0(s))

	rightG = append(
		rightG,
		big.NewInt(0), // w1(s)
	)

	rightG[1] = new(big.Int).Mul(big.NewInt(3), rightG[1]) // a1 = 3
	w[1] = new(bn256.G2).ScalarMult(g2, rightG[1])         // E(a1 * w1(s))

	rightG = append(
		rightG,
		Interpolate(
			s, []int64{1, -7},
			BasisPolynomial(order, 0, []*big.Int{r1, r2}...),
			BasisPolynomial(order, 1, []*big.Int{r1, r2}...),
		), // w2(s)
	)

	rightG[2] = new(big.Int).Mul(big.NewInt(2), rightG[2]) // a2 = 2
	w[2] = new(bn256.G2).ScalarMult(g2, rightG[2])         // E(a2 * w2(s))

	rightG = append(
		rightG,
		Interpolate(
			s, []int64{1},
			BasisPolynomial(order, 1, []*big.Int{r1, r2}...),
		), // w3(s)
	)

	rightG[3] = new(big.Int).Mul(big.NewInt(24), rightG[3]) // a3 = 24
	w[3] = new(bn256.G2).ScalarMult(g2, rightG[3])          // E(a3 * v3(s))

	rightG = append(
		rightG,
		Interpolate(
			s, []int64{3},
			BasisPolynomial(order, 1, []*big.Int{r1, r2}...),
		), // w4(s)
	)

	rightG[4] = new(big.Int).Mul(big.NewInt(1), rightG[4]) // a4 = 1
	w[4] = new(bn256.G2).ScalarMult(g2, rightG[4])         // E(a4 * w4(s))

	rightG = append(
		rightG,
		big.NewInt(0), // w5(s)
	)

	rightG[5] = new(big.Int).Mul(big.NewInt(13), rightG[5]) // a4 = 13
	w[5] = new(bn256.G2).ScalarMult(g2, rightG[5])          // E(a5 * w5(s))

	var y [6]*bn256.G2
	var outputG []*big.Int

	outputG = append(
		outputG,
		big.NewInt(0), // y0(s)
	)

	outputG[0] = new(big.Int).Mul(big.NewInt(1), outputG[0])
	y[0] = new(bn256.G2).ScalarMult(g2, new(big.Int).Sub(order, outputG[0])) // E(-y0(s))

	outputG = append(
		outputG,
		big.NewInt(0), // y1(s)
	)

	outputG[1] = new(big.Int).Mul(big.NewInt(3), outputG[1])                 // a1 = 3
	y[1] = new(bn256.G2).ScalarMult(g2, new(big.Int).Sub(order, outputG[1])) // E(-a1 * v1(s))

	outputG = append(
		outputG,
		big.NewInt(0), // y2(s)
	)

	outputG[2] = new(big.Int).Mul(big.NewInt(2), outputG[2])                 // a2 = 2
	y[2] = new(bn256.G2).ScalarMult(g2, new(big.Int).Sub(order, outputG[2])) // E(-a2 * y2(s))

	outputG = append(
		outputG,
		Interpolate(
			s, []int64{0},
			BasisPolynomial(order, 0, []*big.Int{r1, r2}...),
		), // y3(s)
	)

	outputG[3] = new(big.Int).Mul(big.NewInt(24), outputG[3])                // a3 = 24
	y[3] = new(bn256.G2).ScalarMult(g2, new(big.Int).Sub(order, outputG[3])) // E(-a3 * y3(s))

	outputG = append(
		outputG,
		big.NewInt(0), // y4(s)
	)

	outputG[4] = new(big.Int).Mul(big.NewInt(3), outputG[4])                 // a4 = 1
	y[4] = new(bn256.G2).ScalarMult(g2, new(big.Int).Sub(order, outputG[4])) // E(-a4 * y4(s))

	outputG = append(
		outputG,
		Interpolate(
			s, []int64{1},
			BasisPolynomial(order, 1, []*big.Int{r1, r2}...),
		), // y5(s)
	)

	outputG[5] = new(big.Int).Mul(big.NewInt(13), outputG[5])                // a5 = 2
	y[5] = new(bn256.G2).ScalarMult(g2, new(big.Int).Sub(order, outputG[5])) // E(-a5 * y5(s))

	// E(v0) + E(a1 * v1) + E(a2 * v2) + E(a3 * v3) + E(a4 * v4) + E(a5 * v5)
	//     = E(v0 + a1 * v1 + a2 * v2 + a3 * v3 + a4 * v4 + a5 * v5)
	var eV = new(bn256.G1).Add(
		new(bn256.G1).Add(v[0], v[1]),
		new(bn256.G1).Add(
			new(bn256.G1).Add(v[2], v[3]),
			new(bn256.G1).Add(v[4], v[5]),
		),
	)

	// E(w0) + E(a1 * w1) + E(a2 * w2) + E(a3 * w3) + E(a4 * w4) + E(a5 * w5)
	//     = E(v0 + a1 * w1 + a2 * w2 + a3 * w3 + a4 * w4 + a5 * w5)
	var eW = new(bn256.G2).Add(
		new(bn256.G2).Add(w[0], w[1]),
		new(bn256.G2).Add(
			new(bn256.G2).Add(w[2], w[3]),
			new(bn256.G2).Add(w[4], w[5]),
		),
	)

	// var eY = new(bn256.GT).Add(y[0], new(bn256.GT).Add(y[1], y[2]))
	var eY = bn256.Pair(
		g1,
		new(bn256.G2).Add(
			new(bn256.G2).Add(y[0], y[1]),
			new(bn256.G2).Add(
				new(bn256.G2).Add(y[2], y[3]),
				new(bn256.G2).Add(y[4], y[5]),
			),
		),
	)

	return bytes.Equal(bn256.Pair(eV, eW).Marshal(), eY.Marshal())
}

// E2SQAP defines a strong QAP for the arithmetic expression, uses it to create
// a SNARK, and evaluates it.
func E2SQAP() bool {

	return true
}

// E2R1CS generates the Quadratic Arithmetic Program to validate arithmetic
//  circuits in Zero Knowledge
func E2R1CS() bool {

	// Using the intermediate results.

	// 4 * x1 * x2 - 7 * x2 + 3 * x4

	// 24 = 4 * 2 * 3
	// 13 = 24 - 7 * 3 + 3 * 1

	// The assignment variables (including all intermediate) solutions are then:
	// a1 = 3, a2 = 2, a3 = 24, a4 = 1, a5 = 13

	// Putting the values in a vector and appending a 1 for the constant terms
	// in the linear equations gives:

	// s =

	// [1, 3, 2, 24, 1, 13]

	// For each of the two intermediate results create vectors whose dot product
	// with the solution vector are consistent with the intermediate results.

	// A =

	// [0,  4,  0,  0,  0,  0]
	// [1,  0,  0,  0,  0,  0]

	// B =

	// [0,  0,  1,  0,  0,  0]
	// [0,  0, -7,  1,  3,  0]

	// C =

	// [0,  0,  0,  1,  0,  0]
	// [0,  0,  0,  0,  0,  1]

	// (A1 . s) * (B1 . s) = (4 * a1) * (1 * a2)
	//                     = (4 * 3) * (1 * 2)
	//                     = 24
	//                     = (C1 . s)

	// The Lagrange interpolation polynomials for the QAP can then be derived
	// with the constraints, which is done in: E2QAP

	return true
}
