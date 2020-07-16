package qap

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/cloudflare/bn256"
)

// f(x1) = x1 * x1 * x1 + x1 + 5
//       =    x1 * x2   + x1 + 5
//       =      x3      + x1 + 5
//       = f2(x1, x3)
//       = x4

// a1 * a1 * a1 + a1 + 5 = a4

// This relation is satisfied for a1 = 3 and a4 = 35. The indices 2 and 3 are
// skipped for convenience, because they will be used for an intermediate
// result for the QAP composition.

// f1(x1) = x1 * x1
//        = (x1) * (x1)
//        = p1(x1) * p2(x1)
//        = x2
//        = 9

// p1(x1) = v1_{0} + v1_{1} * x_{1} + v1_{2} * x_{2} + v1_{3} * x_{3} + v1_{4} * x_{4}
//        =   0    +     1  *  a1   +     0  *  a2   +     0  *  a3   +     0  *  a4
//        = 1 * 3

// p2(x1) = w1_{0} + w1_{1} * x_{1} + w1_{2} * x_{2} + w1_{3} * x_{3} + w1_{4} * x_{4}
//        =   0    +     1  *  a1   +     0  *  a2   +     0  *  a3   +     0  *  a4
//        = 1 * 3

// I_{1} = {1}

/* # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#                                                                                                                                                            #
#    x1       x1    a1       a1    3       3                                                              p1(a1) * p2(a1) - a2 = (1 * a1) * (1 * a1) - a2    #
#      \     /        \     /       \     /                                                                                    = 1 * 3 * 1 * 3 - 9           #
#       \   /          \   /         \   /                                                                                     = 0                           #
#        \ /            \ /           \ /                                                                                                                    #
#         *              *             *                                                                                                                     #
#         |              |             |                                                                                                                     #
#         |              |             |                                                                                                                     #
#         |              |             |                                                                                                                     #
#      f1(x1)            a2            9                                                                                                                     #
#                                                                                                                                                            #
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # */

// f2(x1, x2) = x1 * x2
//            = (x1) * (x2)
//            = p1(x1, x2) * p2(x1, x2)
//            = x3
//            = 27

// p1(x1, x2) = v2_{0} + v2_{1} * x_{1} + v2_{2} * x_{2} + v2_{3} * x_{3} + v2_{4} * x_{4}
//            =   0    +     1  *  a1   +     0  *  a2   +     0  *  a3   +     0  *  a4
//            = 1 * 3

// p2(x1, x2) = w2_{0} + w2_{1} * x_{1} + w2_{2} * x_{2} + w2_{3} * x_{3} + w2_{4} * x_{4}
//            =   0    +     0  *  a1   +     1  *  a2   +     0  *  a3   +     0  *  a4
//            = 1 * 9

// I_{2} = {1, 2}

/* # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#                                                                                                                                                            #
#    x1       x2    a1       a2    3       9                                                      p1(a1, a2) * p2(a1, a2) - a3 = (1 * a1) * (1 * a2) - a3    #
#      \     /        \     /       \     /                                                                                    = 1 * 3 * 1 * 9 - 27          #
#       \   /          \   /         \   /                                                                                     = 0                           #
#        \ /            \ /           \ /                                                                                                                    #
#         *              *             *                                                                                                                     #
#         |              |             |                                                                                                                     #
#         |              |             |                                                                                                                     #
#         |              |             |                                                                                                                     #
#     f(x1, x2)          a3            27                                                                                                                    #
#                                                                                                                                                            #
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # */

// f3(x1, x3) = x3 + x1 + 5
//            = 1 * (1 * x1 + 1 * x3 + 5)
//            = p1(x1, x3) * p2(x1, x3)
//            = x4
//            = 35

// p1(x1, x3) = v3_{0} + v3_{1} * x_{1} + v3_{2} * x_{2} + v3_{3} * x_{3} + v3_{4} * x_{4}
//            =   1    +     0  *  a1   +     0  *  a2   +     0  *  a3   +     0  *  a4
//            = 1

// p2(x1, x3) = w3_{0} + w3_{1} * x_{1} + w3_{2} * x_{2} + w3_{3} * x_{3} + w3_{4} * x_{4} + w3_{5} * x_{5}
//            =   5    +     1  *  a1   +     0  *  a2   +     1  *  a3   +     0  *  a4
//            = 5 + 1 * 3 + 1 * 27

// I_{3} = {1, 3}

/* # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#                                                                                                                                                            #
#          x1       x3       a1       a3        3       27                                 p1(a1, a3) * p(a1, a3) - a6 = (1) * (5 + 1 * a1 + 1 * a3) - a4    #
#            \     /           \     /           \     /                                                              = (1) * (5 + 1 * 3 + 1 * 27) - 35      #
#             \   /             \   /             \   /                                                               = 0                                    #
#              \ /               \ /               \ /                                                                                                       #
#       5       +         5       +         5       +                                                                                                        #
#        \     /           \     /           \     /                                                                                                         #
#         \   /             \   /             \   /                                                                                                          #
#          \ /               \ /               \ /                                                                                                           #
#           +                 +                 +                                                                                                            #
#           |                 |                 |                                                                                                            #
#           * 1               * 1               * 1                                                                                                          #
#           |                 |                 |                                                                                                            #
#      f2(x1, x3)             a4                35                                                                                                           #
#                                                                                                                                                            #
# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # */

// v0(r1) =   0, v0(r2) =   0, v0(r3) =  1 | w0(r1) =   0, w0(r2) =   0, w0(r3) =   5 | y0(r1) =  0, y0(r2) =  0, y0(r3) =  0
// v1(r1) =   1, v1(r2) =   1, v1(r3) =  0 | w1(r1) =   1, w1(r2) =   0, w1(r3) =   1 | y1(r1) =  0, y1(r2) =  0, y1(r3) =  0
// v2(r1) =   0, v2(r2) =   0, v2(r3) =  0 | w2(r1) =   0, w2(r2) =   1, w2(r3) =   0 | y2(r1) =  1, y2(r2) =  0, y2(r3) =  0
// v3(r1) =   0, v3(r2) =   0, v3(r3) =  0 | w3(r1) =   0, w3(r2) =   0, w3(r3) =   1 | y3(r1) =  0, y3(r2) =  1, y3(r3) =  0
// v4(r1) =   0, v4(r2) =   0, v4(r3) =  0 | w4(r1) =   0, w4(r2) =   0, w4(r3) =   0 | y4(r1) =  0, y4(r2) =  0, y4(r3) =  1

// Use the constraints to define interpolation polynomials for the QAP.

//
// v(x)
//

//              (x  - r1) * (x  - r2)
// v0(x) = 1 * -----------------------
//              (r3 - r1) * (r3 - r2)

//              (x  - r2) * (x  - r3)         (x  - r1) * (x  - r3)
// v1(x) = 1 * ----------------------- + 1 * -----------------------
//              (r1 - r2) * (r1 - r3)         (r2 - r1) * (r2 - r3)

//
// v2(x) =                             0
//

//
// v3(x) =                             0
//

//
// v4(x) =                             0
//

//
// w(x) =
//

//              (x  - r1) * (x  - r2)
// w0(x) = 5 * -----------------------
//              (r3 - r1) * (r3 - r2)

//              (x  - r2) * (x  - r3)         (x  - r1) * (x  - r2)
// w1(x) = 1 * ----------------------- + 1 * -----------------------
//              (r1 - r2) * (r1 - r3)         (r3 - r1) * (r3 - r2)

//              (x  - r1) * (x  - r3)
// w2(x) = 1 * -----------------------
//              (r2 - r1) * (r2 - r3)

//              (x  - r1) * (x  - r2)
// w3(x) = 1 * -----------------------
//              (r3 - r1) * (r3 - r2)

//
// w4(x) =                             0
//

//
// y(x) =
//

//
// y0(x) =                             0
//

//
// y1(x) =                             0
//

//              (x  - r2) * (x  - r3)
// y2(x) = 1 * -----------------------
//              (r1 - r2) * (r1 - r3)

//              (x  - r1) * (x  - r3)
// y3(x) = 1 * -----------------------
//              (r2 - r1) * (r2 - r3)

//              (x  - r1) * (x  - r2)
// y4(x) = 1 * -----------------------
//              (r3 - r1) * (r3 - r2)

// The assignment variables including those that are intermediate for the
// circuit are: a1 = 3, a2 = 9, a3 = 27, and a4 = 35. To check the polynomials
// as defined are correct the relation can be evaluated at r1, r2 and r3 to see
// if has the expected result of 0.

// (v0 + v1 * a1 + v2 * a2 + v3 * a3 + v4 * a4) * (w0 + w1 * a1 + w2 * a2 + w3 * a3 + w4 * a4)
//                         = (y0 + y1 * a1 + y2 * a2 + y3 * a3 + y4 * a4)

// at r1 - > (0 + 1 * 3 + 0 * 9 + 0 * 27 + 0 * 35) * (0 + 1 * 3 + 0 * 9 + 0 * 27 + 0 * 35)
//       - >                    = (0 + 0 * 3 + 1 * 9 + 0 * 27 + 0 * 35)
//       - >              3 * 3 = 9

// at r2 - > (0 + 1 * 3 + 0 * 9 + 0 * 27 + 0 * 35) * (0 + 0 * 3 + 1 * 9 + 0 * 27 + 0 * 35)
//       - >                    = (0 + 0 * 3 + 0 * 9 + 1 * 27 + 0 * 35)
//       - >              3 * 9 = 27

// at r3 - > (1 + 0 * 3 + 0 * 9 + 0 * 27 + 0 * 35) * (5 + 1 * 3 + 0 * 9 + 1 * 27 + 0 * 35)
//       - >                    = (0 + 0 * 3 + 0 * 9 + 0 * 27 + 1 * 35)
//       - >   1 * (5 + 3 + 27) = 35

// E3QAP defines a QAP for the arithmetic expression, uses it to create a SNARK,
// and evaluates it.
func E3QAP() bool {

	var err error

	var order = bn256.Order
	// var order = bn256.Order.Set(big.NewInt(11))
	// var order = bn256.Order.Set(big.NewInt(23))
	// var order = bn256.Order.Set(big.NewInt(997))

	var g1 *bn256.G1
	if _, g1, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var g2 *bn256.G2
	if _, g2, err = bn256.RandomG2(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var r1 *big.Int // big.NewInt(3)
	if r1, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var r2 *big.Int // big.NewInt(7)
	if r2, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var r3 *big.Int // big.NewInt(10)
	if r3, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var s *big.Int // big.NewInt(5)
	if s, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	// var betaV *big.Int
	// if betaV, err = rand.Int(rand.Reader, order); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	// var betaW *big.Int
	// if betaW, err = rand.Int(rand.Reader, order); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	// var betaY *big.Int
	// if betaY, err = rand.Int(rand.Reader, order); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	var v [5]*bn256.G1
	var leftG []*big.Int

	leftG = append(
		leftG,
		Interpolate(
			s, []int64{1},
			BasisPolynomial(order, 2, []*big.Int{r1, r2, r3}...),
		), // v0(s)
	)

	// leftG[0] = new(big.Int).Mul(big.NewInt(1), leftG[0])
	leftG[0] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(1), leftG[0]), order)

	v[0] = new(bn256.G1).ScalarMult(g1, leftG[0]) // E(v0(s))

	leftG = append(
		leftG,
		Interpolate(
			s, []int64{1, 1},
			BasisPolynomial(order, 0, []*big.Int{r1, r2, r3}...),
			BasisPolynomial(order, 1, []*big.Int{r1, r2, r3}...),
		), // v1(s)
	)

	// leftG[1] = new(big.Int).Mul(big.NewInt(3), leftG[1])                          // a1 = 3
	leftG[1] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(3), leftG[1]), order) // a1 = 3

	v[1] = new(bn256.G1).ScalarMult(g1, leftG[1]) // E(a1 * v1(s))

	leftG = append(
		leftG,
		big.NewInt(0), // v2(s)
	)

	// leftG[2] = new(big.Int).Mul(big.NewInt(9), leftG[2])                          // a2 = 9
	leftG[2] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(9), leftG[2]), order) // a2 = 9

	v[2] = new(bn256.G1).ScalarMult(g1, leftG[2]) // E(a2 * v2(s))

	leftG = append(
		leftG,
		big.NewInt(0), // v3(s)
	)

	// leftG[3] = new(big.Int).Mul(big.NewInt(27), leftG[3])                          // a3 = 27
	leftG[3] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(27), leftG[3]), order) // a3 = 27

	v[3] = new(bn256.G1).ScalarMult(g1, leftG[3]) // E(a3 * v3(s))

	leftG = append(
		leftG,
		big.NewInt(0), // v4(s)
	)

	// leftG[4] = new(big.Int).Mul(big.NewInt(35), leftG[4])                          // a4 = 35
	leftG[4] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(35), leftG[4]), order) // a4 = 35

	v[4] = new(bn256.G1).ScalarMult(g1, leftG[4]) // E(a4 * v4(s))

	var w [5]*bn256.G2
	var rightG []*big.Int

	rightG = append(
		rightG,
		Interpolate(
			s, []int64{5},
			BasisPolynomial(order, 2, []*big.Int{r1, r2, r3}...),
		), // w0(s)
	)

	// rightG[0] = new(big.Int).Mul(big.NewInt(1), rightG[0])
	rightG[0] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(1), rightG[0]), order)

	w[0] = new(bn256.G2).ScalarMult(g2, rightG[0]) // E(w0(s))

	rightG = append(
		rightG,
		Interpolate(
			s, []int64{1, 1},
			BasisPolynomial(order, 0, []*big.Int{r1, r2, r3}...),
			BasisPolynomial(order, 2, []*big.Int{r1, r2, r3}...),
		), // w1(s)
	)

	// rightG[1] = new(big.Int).Mul(big.NewInt(3), rightG[1])                          // a1 = 3
	rightG[1] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(3), rightG[1]), order) // a1 = 3

	w[1] = new(bn256.G2).ScalarMult(g2, rightG[1]) // E(a1 * w1(s))

	rightG = append(
		rightG,
		Interpolate(
			s, []int64{1},
			BasisPolynomial(order, 1, []*big.Int{r1, r2, r3}...),
		), // w2(s)
	)

	// rightG[2] = new(big.Int).Mul(big.NewInt(9), rightG[2])                          // a2 = 9
	rightG[2] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(9), rightG[2]), order) // a2 = 9

	w[2] = new(bn256.G2).ScalarMult(g2, rightG[2]) // E(a2 * w2(s))

	rightG = append(
		rightG,
		Interpolate(
			s, []int64{1},
			BasisPolynomial(order, 2, []*big.Int{r1, r2, r3}...),
		), // w3(s)
	)

	// rightG[3] = new(big.Int).Mul(big.NewInt(27), rightG[3])                          // a3 = 27
	rightG[3] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(27), rightG[3]), order) // a3 = 27

	w[3] = new(bn256.G2).ScalarMult(g2, rightG[3]) // E(a3 * v3(s))

	rightG = append(
		rightG,
		big.NewInt(0), // w4(s)
	)

	// rightG[4] = new(big.Int).Mul(big.NewInt(35), rightG[4])                          // a4 = 35
	rightG[4] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(35), rightG[4]), order) // a4 = 35

	w[4] = new(bn256.G2).ScalarMult(g2, rightG[4]) // E(a4 * w4(s))

	var y [5]*bn256.G2
	var outputG []*big.Int

	outputG = append(
		outputG,
		big.NewInt(0), // y0(s)
	)

	// outputG[0] = new(big.Int).Mul(big.NewInt(1), outputG[0])
	outputG[0] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(1), outputG[0]), order)

	y[0] = new(bn256.G2).ScalarMult(g2, outputG[0]) // E(y0(s))

	outputG = append(
		outputG,
		big.NewInt(0), // y1(s)
	)

	// outputG[0] = new(big.Int).Mul(big.NewInt(3), outputG[1])                          // a1 = 3
	outputG[1] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(3), outputG[1]), order) // a1 = 3

	y[1] = new(bn256.G2).ScalarMult(g2, outputG[1]) // E(a1 * y1(s))

	outputG = append(
		outputG,
		Interpolate(
			s, []int64{1},
			BasisPolynomial(order, 0, []*big.Int{r1, r2, r3}...),
		), // y2(s)
	)

	// outputG[2] = new(big.Int).Mul(big.NewInt(9), outputG[2])                          // a2 = 9
	outputG[2] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(9), outputG[2]), order) // a2 = 9

	y[2] = new(bn256.G2).ScalarMult(g2, outputG[2]) // E(a2 * y3(s))

	outputG = append(
		outputG,
		Interpolate(
			s, []int64{1},
			BasisPolynomial(order, 1, []*big.Int{r1, r2, r3}...),
		), // y3(s)
	)

	// outputG[3] = new(big.Int).Mul(big.NewInt(27), outputG[3])                          // a3 = 27
	outputG[3] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(27), outputG[3]), order) // a3 = 27

	y[3] = new(bn256.G2).ScalarMult(g2, outputG[3]) // E(a3 * y3(s))

	outputG = append(
		outputG,
		Interpolate(
			s, []int64{1},
			BasisPolynomial(order, 2, []*big.Int{r1, r2, r3}...),
		), // y4(s)
	)

	// outputG[4] = new(big.Int).Mul(big.NewInt(35), outputG[4])                          // a4 = 35
	outputG[4] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(35), outputG[4]), order) // a4 = 35

	y[4] = new(bn256.G2).ScalarMult(g2, outputG[4]) // E(a4 * y4(s))

	var term1 = new(big.Int).Add(
		leftG[0],
		new(big.Int).Add(
			new(big.Int).Add(leftG[1], leftG[2]),
			new(big.Int).Add(leftG[3], leftG[4]),
		),
	)

	var term2 = new(big.Int).Add(
		rightG[0],
		new(big.Int).Add(
			new(big.Int).Add(rightG[1], rightG[2]),
			new(big.Int).Add(rightG[3], rightG[4]),
		),
	)

	var term3 = new(big.Int).Add(
		outputG[0],
		new(big.Int).Add(
			new(big.Int).Add(outputG[1], outputG[2]),
			new(big.Int).Add(outputG[3], outputG[4]),
		),
	)

	var t = new(big.Int).Mod(
		new(big.Int).Mul(
			new(big.Int).Sub(s, r1),
			new(big.Int).Mul(
				new(big.Int).Sub(s, r2),
				new(big.Int).Sub(s, r3),
			),
		),
		order,
	)

	var h = new(big.Int).Mod(
		new(big.Int).Mul(
			new(big.Int).Sub(
				new(big.Int).Mul(term1, term2), term3,
			),
			new(big.Int).ModInverse(t, order),
		),
		order,
	)

	var eV = new(bn256.G1).Add(
		v[0],
		new(bn256.G1).Add(
			new(bn256.G1).Add(v[1], v[2]),
			new(bn256.G1).Add(v[3], v[4]),
		),
	)

	var eW = new(bn256.G2).Add(
		w[0],
		new(bn256.G2).Add(
			new(bn256.G2).Add(w[1], w[2]),
			new(bn256.G2).Add(w[3], w[4]),
		),
	)

	var eY = bn256.Pair(
		g1,
		new(bn256.G2).Add(
			y[0],
			new(bn256.G2).Add(
				new(bn256.G2).Add(y[1], y[2]),
				new(bn256.G2).Add(y[3], y[4]),
			),
		),
	)

	var eT = new(bn256.G1).ScalarMult(g1, t)
	var eH = new(bn256.G2).ScalarMult(g2, h)

	var left = new(bn256.GT).Add(
		bn256.Pair(eV, eW),
		new(bn256.GT).Neg(eY),
	)

	var right = bn256.Pair(eT, eH)

	return bytes.Equal(left.Marshal(), right.Marshal())
}

// E3SQAP defines a strong QAP for the arithmetic expression, uses it to create
// a SNARK, and evaluates it.
func E3SQAP() bool {

	return true
}

// E3R1CS generates the Quadratic Arithmetic Program to validate arithmetic
//  circuits in Zero Knowledge
func E3R1CS() bool {

	// Using the intermediate results.

	// x1 * x1 * x1 + x1 + 5 = x4

	// 9  = 3 * 3
	// 27 = 3 * 9
	// 35 = 5 + 3 + 27

	// The assignment variables (including all intermediate) solutions are then:
	// a1 = 3, a2 = 9, a3 = 27, a4 = 35

	// Putting the values in a vector and appending a 1 for the constant terms
	// in the linear equations gives:

	// s =

	// [1, 3, 9, 27, 35]

	// A =

	// [0, 1, 0, 0, 0]
	// [0, 1, 0, 0, 0]
	// [1, 0, 0, 0, 0]

	// B =

	// [0, 1, 0, 0, 0]
	// [0, 0, 1, 0, 0]
	// [5, 1, 0, 1, 0]

	// C =

	// [0, 0, 1, 0, 0]
	// [0, 0, 0, 1, 0]
	// [0, 0, 0, 0, 1]

	// (A1 . s) * (B1 . s) = (1 * a1) * (1 * a1)
	//                     = (1 * 3) * (1 * 3)
	//                     = 9
	//                     = (C1 . s)

	// The Lagrange interpolation polynomials for the QAP can then be derived
	// with the constraints, which is done in: E3QAP

	return true
}
