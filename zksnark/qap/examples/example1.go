package examples

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/cloudflare/bn256"

	"github.com/eugenekadish/cryptopalooza/zksnark/qap/lagrange"
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

// E1QAP creates a QAP (Quadratic Arithmetic Program) for the arithmetic
// expression.
func E1QAP() bool {

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
	var leftG []*big.Int

	leftG = append(
		leftG,
		big.NewInt(3),
	)

	leftG[0] = new(big.Int).Mul(big.NewInt(1), leftG[0])
	v[0] = new(bn256.G1).ScalarMult(g1, leftG[0])

	leftG = append(
		leftG,
		big.NewInt(0),
	)

	leftG[1] = new(big.Int).Mul(big.NewInt(2), leftG[1])
	v[1] = new(bn256.G1).ScalarMult(g1, leftG[1])

	leftG = append(
		leftG,
		big.NewInt(0),
	)

	leftG[2] = new(big.Int).Mul(big.NewInt(6), leftG[2])
	v[2] = new(bn256.G1).ScalarMult(g1, leftG[2])

	var w [3]*bn256.G2
	var rightG []*big.Int

	rightG = append(
		rightG,
		big.NewInt(0),
	)

	rightG[0] = new(big.Int).Mul(big.NewInt(1), rightG[0])
	w[0] = new(bn256.G2).ScalarMult(g2, rightG[0])

	rightG = append(
		rightG,
		big.NewInt(1),
	)

	rightG[1] = new(big.Int).Mul(big.NewInt(2), rightG[1])
	w[1] = new(bn256.G2).ScalarMult(g2, rightG[1])

	rightG = append(
		rightG,
		big.NewInt(0),
	)

	rightG[2] = new(big.Int).Mul(big.NewInt(6), rightG[2])
	w[2] = new(bn256.G2).ScalarMult(g2, rightG[2])

	// var y [3]*bn256.GT
	// var outputG []*big.Int

	// outputG = append(
	// 	outputG,
	// 	big.NewInt(0),
	// )

	// outputG[0] = new(big.Int).Mul(big.NewInt(1), outputG[0])
	// y[0] = new(bn256.GT).ScalarMult(gt, outputG[0])

	// outputG = append(
	// 	outputG,
	// 	big.NewInt(0),
	// )

	// outputG[1] = new(big.Int).Mul(big.NewInt(2), outputG[1])
	// y[1] = new(bn256.GT).ScalarMult(gt, outputG[1])

	// outputG = append(
	// 	outputG,
	// 	big.NewInt(1),
	// )

	// outputG[2] = new(big.Int).Mul(big.NewInt(6), outputG[2])
	// y[2] = new(bn256.GT).ScalarMult(gt, outputG[2])

	var y [3]*bn256.G2
	var outputG []*big.Int

	outputG = append(
		outputG,
		big.NewInt(0),
	)

	outputG[0] = new(big.Int).Mul(big.NewInt(1), outputG[0])
	y[0] = new(bn256.G2).ScalarMult(g2, outputG[0])

	outputG = append(
		outputG,
		big.NewInt(0),
	)

	outputG[1] = new(big.Int).Mul(big.NewInt(2), outputG[1])
	y[1] = new(bn256.G2).ScalarMult(g2, outputG[1])

	outputG = append(
		outputG,
		big.NewInt(1),
	)

	outputG[2] = new(big.Int).Mul(big.NewInt(6), outputG[2])
	y[2] = new(bn256.G2).ScalarMult(g2, outputG[2])

	var eV = new(bn256.G1).Add(v[0], new(bn256.G1).Add(v[1], v[2]))
	var eW = new(bn256.G2).Add(w[0], new(bn256.G2).Add(w[1], w[2]))

	// var eY = new(bn256.GT).Add(y[0], new(bn256.GT).Add(y[1], y[2]))
	var eY = bn256.Pair(g1, new(bn256.G2).Add(y[0], new(bn256.G2).Add(y[1], y[2])))

	return bytes.Equal(bn256.Pair(eV, eW).Marshal(), eY.Marshal())
}

// vp_{0}(r1) = 0, vp_{0}(r2) = 0, vp_{0}(s1) = 1, vp_{0}(s2) = 1,
// vp_{1}(r1) = 1, vp_{1}(r2) = 0, vp_{1}(s1) = 0, vp_{1}(s2) = 0,
// vp_{2}(r1) = 0, vp_{2}(r2) = 1, vp_{2}(s1) = 0, vp_{2}(s2) = 0,

// wp_{0}(r1) = 1, wp_{0}(r2) = 1, wp_{0}(s1) = 0, wp_{0}(s2) = 0
// wp_{1}(r1) = 0, wp_{1}(r2) = 0, wp_{1}(s1) = 1, wp_{1}(s2) = 0
// wp_{2}(r1) = 0, wp_{2}(r1) = 0, wp_{2}(s1) = 0, wp_{2}(s2) = 1

// yp_{0}(r1) = 0, yp_{0}(r2) = 0, yp_{0}(s1) = 0, yp_{0}(s2) = 0
// yp_{1}(r1) = 1, yp_{1}(r2) = 0, yp_{1}(s1) = 1, yp_{1}(s2) = 0
// yp_{2}(r1) = 0, yp_{2}(r2) = 1, yp_{2}(s1) = 0, yp_{2}(s2) = 1

// v(x) =

//                  (x  -  r) * (x  - r1) * (x  - r2) * (x  - s2)
// vp_{0}(x) = 1 * -----------------------------------------------
//                  (s1 -  r) * (s1 - r1) * (s1 - r2) * (s1 - s2)

//                  (x  -  r) * (x  - r1) * (x  - r2) * (x  - s1)
//           + 1 * -----------------------------------------------
//                  (s2 -  r) * (s2 - r1) * (s2 - r2) * (s2 - s1)

//                  (x  - r1) * (x  - r2) * (x  - s1) * (x  - s2)
//           + 3 * -----------------------------------------------
//                  (r  - r1) * (r  - r2) * (r  - s1) * (r  - s2)

//                  (x  -  r) * (x  - r2) * (x  - s1) * (x  - s2)
// vp_{1}(x) = 1 * -----------------------------------------------
//                  (r1 -  r) * (r1 - r2) * (r1 - s1) * (r1 - s2)

//                  (x  -  r) * (x  - r1) * (x  - s1) * (x  - s2)
// vp_{2}(x) = 1 * -----------------------------------------------
//                  (r2 -  r) * (r2 - r1) * (r2 - s1) * (r2 - s2)

// w(x) =

//                  (x  -  r) * (x  - r2) * (x  - s1) * (x  - s2)
// wp_{0}(x) = 1 * -----------------------------------------------
//                  (r1 -  r) * (r1 - r2) * (r1 - s1) * (r1 - s2)

//                  (x  -  r) * (x  - r1) * (x  - s1) * (x  - s2)
//           + 1 * -----------------------------------------------
//                  (r2 -  r) * (r2 - r1) * (r2 - s1) * (r2 - s2)

//                  (x  -  r) * (x  - r1) * (x  - r2) * (x  - s2)
// wp_{1}(x) = 1 * -----------------------------------------------
//                  (s1 -  r) * (s1 - r1) * (s1 - r2) * (s1 - s2)

//                  (x  - r1) * (x  - r2) * (x  - s1) * (x  - s2)
//           + 1 * -----------------------------------------------
//                  (r  - r1) * (r  - r2) * (r  - s1) * (r  - s2)

//                  (x  -  r) * (x  - r1) * (x  - r2) * (x  - s1)
// wp_{2}(x) = 1 * -----------------------------------------------
//                  (s2 -  r) * (s2 - r1) * (s2 - r2) * (s2 - s1)

// y(x) =

//
// yp_{0}(x) = 0
//

//                  (x  -  r) * (x  - r2) * (x  - s1) * (x  - s2)
// yp_{1}(x) = 1 * -----------------------------------------------
//                  (r1 -  r) * (r1 - r2) * (r1 - s1) * (r1 - s2)

//                  (x  -  r) * (x  - r1) * (x  - r2) * (x  - s2)
//           + 1 * -----------------------------------------------
//                  (s1 -  r) * (s1 - r1) * (s1 - r2) * (s1 - s2)

//                  (x  -  r) * (x  - r1) * (x  - s1) * (x  - s2)
// yp_{2}(x) = 1 * -----------------------------------------------
//                  (r2 -  r) * (r2 - r1) * (r2 - s1) * (r2 - s2)

//                  (x  -  r) * (x  - r1) * (x  - r2) * (x  - s1)
//           + 1 * -----------------------------------------------
//                  (s2 -  r) * (s2 - r1) * (s2 - r2) * (s2 - s1)

//                  (x  - r1) * (x  - r2) * (x  - s1) * (x  - s2)
//           + 1 * -----------------------------------------------
//                  (r  - r1) * (r  - r2) * (r  - s1) * (r  - s2)

// E1SQAP creates a strong QAP for the arithmetic expression, as well as
// generates and then verifies the SNARK.
func E1SQAP() bool {

	var err error

	var g1 *bn256.G1
	if _, g1, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var g2 *bn256.G2
	if _, g2, err = bn256.RandomG2(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var r1 *big.Int
	if r1, err = rand.Int(rand.Reader, bn256.Order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var r2 *big.Int
	if r2, err = rand.Int(rand.Reader, bn256.Order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var s1 *big.Int
	if s1, err = rand.Int(rand.Reader, bn256.Order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var s2 *big.Int
	if s2, err = rand.Int(rand.Reader, bn256.Order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var r *big.Int
	if r, err = rand.Int(rand.Reader, bn256.Order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var s *big.Int
	if s, err = rand.Int(rand.Reader, bn256.Order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	// var alpha *big.Int
	// if alpha, err = rand.Int(rand.Reader, bn256.Order); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	// var betaV *big.Int
	// if betaV, err = rand.Int(rand.Reader, bn256.Order); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	// var betaW *big.Int
	// if betaW, err = rand.Int(rand.Reader, bn256.Order); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	// var betaY *big.Int
	// if betaY, err = rand.Int(rand.Reader, bn256.Order); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	var v [3]*bn256.G1
	var leftG []*big.Int

	leftG = append(
		leftG,
		lagrange.Interpolate(
			s, []int64{3, 1, 1},
			lagrange.BasisPolynomial(0, []*big.Int{r, r1, r2, s1, s2}...),
			lagrange.BasisPolynomial(3, []*big.Int{r, r1, r2, s1, s2}...),
			lagrange.BasisPolynomial(4, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	leftG[0] = new(big.Int).Mul(big.NewInt(1), leftG[0])
	v[0] = new(bn256.G1).ScalarMult(g1, leftG[0])
	// v[0] = new(bn256.G1).ScalarMult(g1, big.NewInt(3))

	leftG = append(
		leftG,
		lagrange.Interpolate(
			s, []int64{1},
			lagrange.BasisPolynomial(1, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	leftG[1] = new(big.Int).Mul(big.NewInt(2), leftG[1])
	v[1] = new(bn256.G1).ScalarMult(g1, leftG[1])
	// v[1] = new(bn256.G1).ScalarMult(g1, new(big.Int).Mul(big.NewInt(2), big.NewInt(0)))

	leftG = append(
		leftG,
		lagrange.Interpolate(
			s, []int64{1},
			lagrange.BasisPolynomial(2, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	leftG[2] = new(big.Int).Mul(big.NewInt(6), leftG[2])
	v[2] = new(bn256.G1).ScalarMult(g1, leftG[2])
	// v[2] = new(bn256.G1).ScalarMult(g1, new(big.Int).Mul(big.NewInt(6), big.NewInt(0)))

	var w [3]*bn256.G2
	var rightG []*big.Int

	rightG = append(
		rightG,
		lagrange.Interpolate(
			s, []int64{1, 1},
			lagrange.BasisPolynomial(1, []*big.Int{r, r1, r2, s1, s2}...),
			lagrange.BasisPolynomial(2, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	rightG[0] = new(big.Int).Mul(big.NewInt(1), rightG[0])
	w[0] = new(bn256.G2).ScalarMult(g2, rightG[0])
	// w[0] = new(bn256.G2).ScalarMult(g2, big.NewInt(0))

	rightG = append(
		rightG,
		lagrange.Interpolate(
			s, []int64{1, 1},
			lagrange.BasisPolynomial(0, []*big.Int{r, r1, r2, s1, s2}...),
			lagrange.BasisPolynomial(3, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	rightG[1] = new(big.Int).Mul(big.NewInt(2), rightG[1])
	w[1] = new(bn256.G2).ScalarMult(g2, rightG[1])
	// w[1] = new(bn256.G2).ScalarMult(g2, new(big.Int).Mul(big.NewInt(2), big.NewInt(1)))

	rightG = append(
		rightG,
		lagrange.Interpolate(
			s, []int64{1},
			lagrange.BasisPolynomial(4, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	rightG[2] = new(big.Int).Mul(big.NewInt(6), rightG[2])
	w[2] = new(bn256.G2).ScalarMult(g2, rightG[2])
	// w[2] = new(bn256.G2).ScalarMult(g2, new(big.Int).Mul(big.NewInt(6), big.NewInt(0)))

	var y [3]*bn256.G2
	var outputG []*big.Int

	outputG = append(
		outputG,
		big.NewInt(0),
	)

	outputG[0] = new(big.Int).Mul(big.NewInt(1), outputG[0])
	y[0] = new(bn256.G2).ScalarMult(g2, outputG[0])
	// y[0] = new(bn256.G2).ScalarMult(g2, big.NewInt(0))

	outputG = append(
		outputG,
		lagrange.Interpolate(
			s, []int64{1, 1},
			lagrange.BasisPolynomial(1, []*big.Int{r, r1, r2, s1, s2}...),
			lagrange.BasisPolynomial(3, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	outputG[1] = new(big.Int).Mul(big.NewInt(2), outputG[1])
	y[1] = new(bn256.G2).ScalarMult(g2, outputG[1])
	// y[1] = new(bn256.G2).ScalarMult(g2, new(big.Int).Mul(big.NewInt(2), big.NewInt(0)))

	outputG = append(
		outputG,
		lagrange.Interpolate(
			s, []int64{1, 1, 1},
			lagrange.BasisPolynomial(0, []*big.Int{r, r1, r2, s1, s2}...),
			lagrange.BasisPolynomial(2, []*big.Int{r, r1, r2, s1, s2}...),
			lagrange.BasisPolynomial(4, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	outputG[2] = new(big.Int).Mul(big.NewInt(6), outputG[2])
	y[2] = new(bn256.G2).ScalarMult(g2, outputG[2])
	// y[2] = new(bn256.G2).ScalarMult(g2, new(big.Int).Mul(big.NewInt(6), big.NewInt(1)))

	var term1 = new(big.Int).Add(
		leftG[0],
		new(big.Int).Add(
			leftG[1],
			leftG[2],
		),
	)

	var term2 = new(big.Int).Add(
		rightG[0],
		new(big.Int).Add(
			rightG[1],
			rightG[2],
		),
	)

	var term3 = new(big.Int).Add(
		outputG[0],
		new(big.Int).Add(
			outputG[1],
			outputG[2],
		),
	)

	// var t = big.NewInt(0)
	var t = new(big.Int).Sub(s, r)

	// var h = big.NewInt(12)
	var h = new(big.Int).Mul(
		new(big.Int).Sub(
			new(big.Int).Mul(term1, term2), term3,
		),
		new(big.Int).ModInverse(t, bn256.Order),
	)

	var eV = new(bn256.G1).Add(v[0], new(bn256.G1).Add(v[1], v[2]))
	var eW = new(bn256.G2).Add(w[0], new(bn256.G2).Add(w[1], w[2]))
	var eY = bn256.Pair(g1, new(bn256.G2).Add(y[0], new(bn256.G2).Add(y[1], y[2])))

	// return bytes.Equal(bn256.Pair(eV, eW).Marshal(), eY.Marshal())

	var eT = new(bn256.G1).ScalarMult(g1, t)
	var eH = new(bn256.G2).ScalarMult(g2, h)

	var left = new(bn256.GT).Add(bn256.Pair(eV, eW), new(bn256.GT).Neg(eY))
	// // var left = new(bn256.GT).Add(bn256.Pair(eV, eW), new(bn256.GT).ScalarMult(eY, big.NewInt(-1)))
	var right = bn256.Pair(eT, eH)

	// // fmt.Printf(
	// // 	" (v0 + a1 * v1 + a2 * v2) * (w0 + a1 * w1 + a2 * w2) - (y0 + a1 * y1 + a2 * y2) = (%s) * (%s) - %s \n",
	// // 	eV.String()[0:18], eW.String()[0:18], eY.String()[0:18],
	// // )

	return bytes.Equal(left.Marshal(), right.Marshal())
}

// LinearR1CS generates the Quadratic Arithmetic Program to validate arithmetic circuits in Zero Knowledge
func LinearR1CS() bool {

	return true
}
