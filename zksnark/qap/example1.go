package qap

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

// p1(x1) = v1_{0} + v1_{1} * x_{1}
//        = 3 + 0 * x1

// p2(x1) = w1_{0} + w1_{1} * x_{1}
//        = 0 + 1 * 2

/* # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # #
#                                                                                                                                                            #
#    3       x1    3       a1    3       2                                                                           p(a1) * p(a1) - a2 = (3) * (a1) - a2    #
#     \     /       \     /       \     /                                                                                               = 3 * 2 - 6          #
#      \   /         \   /         \   /                                                                                                = 0                  #
#       \ /           \ /           \ /                                                                                                                      #
#        *             *             *                                                                                                                       #
#        |             |             |                                                                                                                       #
#        |             |             |                                                                                                                       #
#        |             |             |                                                                                                                       #
#      f(x1)           a2            6                                                                                                                       #
#                                                                                                                                                            #
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

// v_{0}(x) = 3 | w_{0}(x) = 0 | y_{0}(x) = 0
// v_{1}(x) = 0 | w_{1}(x) = 1 | y_{1}(x) = 0
// v_{2}(x) = 0 | w_{2}(x) = 0 | y_{2}(x) = 1

// E1QAP defines a QAP for the arithmetic expression, uses it to create a SNARK,
// and evaluates it.
func E1QAP() bool {

	var err error

	var g1 *bn256.G1
	if _, g1, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("error generating group element: %v \n", err)
	}

	var g2 *bn256.G2
	if _, g2, err = bn256.RandomG2(rand.Reader); err != nil {
		fmt.Printf("error generating group element %v \n", err)
	}

	var v [3]*bn256.G1
	var leftG []*big.Int

	leftG = append(
		leftG,
		big.NewInt(3), // v0(s)
	)

	leftG[0] = new(big.Int).Mul(big.NewInt(1), leftG[0])
	v[0] = new(bn256.G1).ScalarMult(g1, leftG[0]) // E(v0(s))

	leftG = append(
		leftG,
		big.NewInt(0), // v1(s)
	)

	leftG[1] = new(big.Int).Mul(big.NewInt(2), leftG[1]) // a1 = 2
	v[1] = new(bn256.G1).ScalarMult(g1, leftG[1])        // E(a1 * v1(s))

	leftG = append(
		leftG,
		big.NewInt(0),
	)

	leftG[2] = new(big.Int).Mul(big.NewInt(6), leftG[2]) // a2 = 6
	v[2] = new(bn256.G1).ScalarMult(g1, leftG[2])        // E(a2 * v2(s))

	var w [3]*bn256.G2
	var rightG []*big.Int

	rightG = append(
		rightG,
		big.NewInt(0), // w0(s)
	)

	rightG[0] = new(big.Int).Mul(big.NewInt(1), rightG[0])
	w[0] = new(bn256.G2).ScalarMult(g2, rightG[0]) // E(w0(s))

	rightG = append(
		rightG,
		big.NewInt(1), // w1(s)
	)

	rightG[1] = new(big.Int).Mul(big.NewInt(2), rightG[1]) // a1 = 2
	w[1] = new(bn256.G2).ScalarMult(g2, rightG[1])         // E(a1 * w1(s))

	rightG = append(
		rightG,
		big.NewInt(0), // w2(s)
	)

	rightG[2] = new(big.Int).Mul(big.NewInt(6), rightG[2]) // a2 = 6
	w[2] = new(bn256.G2).ScalarMult(g2, rightG[2])         // E(a2 * v2(s))

	var y [3]*bn256.G2
	var outputG []*big.Int

	outputG = append(
		outputG,
		big.NewInt(0), // y0(s)
	)

	outputG[0] = new(big.Int).Mul(big.NewInt(1), outputG[0])
	y[0] = new(bn256.G2).ScalarMult(g2, outputG[0]) // E(y0(s))

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

	// Quadratic root detection to validate the SNARK was constructed with
	// values that satisfy the arithmetic circuit.

	var eV = new(bn256.G1).Add(v[0], new(bn256.G1).Add(v[1], v[2]))
	var eW = new(bn256.G2).Add(w[0], new(bn256.G2).Add(w[1], w[2]))
	var eY = bn256.Pair(g1, new(bn256.G2).Add(y[0], new(bn256.G2).Add(y[1], y[2])))

	// return bytes.Equal(bn256.Pair(eV, eW).Marshal(), eY.Marshal())

	var left = new(bn256.GT).Add(
		bn256.Pair(eV, eW),
		new(bn256.GT).Neg(eY),
	)

	var right = bn256.Pair(
		g1,
		new(bn256.G2).ScalarMult(g2, big.NewInt(0)),
	)

	return bytes.Equal(left.Marshal(), right.Marshal())
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

// E1SQAP defines a string QAP for the arithmetic expression, uses it to create
// a SNARK, and evaluates it.
func E1SQAP() bool {

	var err error

	var order = bn256.Order
	// var order = bn256.Order.Set(big.NewInt(11))
	// var order = bn256.Order.Set(big.NewInt(997))

	var g1 *bn256.G1
	if _, g1, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var g2 *bn256.G2
	if _, g2, err = bn256.RandomG2(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var r *big.Int // big.NewInt(10)
	if r, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var r1 *big.Int // big.NewInt(2)
	if r1, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var r2 *big.Int // big.NewInt(3)
	if r2, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var s *big.Int // big.NewInt(22)
	if s, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var s1 *big.Int // big.NewInt(5)
	if s1, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var s2 *big.Int // big.NewInt(7)
	if s2, err = rand.Int(rand.Reader, order); err != nil {
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

	var v [3]*bn256.G1
	var leftG []*big.Int

	leftG = append(
		leftG,
		Interpolate(
			s, []int64{3, 1, 1},
			BasisPolynomial(order, 0, []*big.Int{r, r1, r2, s1, s2}...),
			BasisPolynomial(order, 3, []*big.Int{r, r1, r2, s1, s2}...),
			BasisPolynomial(order, 4, []*big.Int{r, r1, r2, s1, s2}...),
		), // v0(s)
	)

	// leftG[0] = new(big.Int).Mul(big.NewInt(1), leftG[0])
	leftG[0] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(1), leftG[0]), order)

	v[0] = new(bn256.G1).ScalarMult(g1, leftG[0]) // E(v0(s))

	leftG = append(
		leftG,
		Interpolate(
			s, []int64{1},
			BasisPolynomial(order, 1, []*big.Int{r, r1, r2, s1, s2}...),
		), // v1(s)
	)

	// leftG[1] = new(big.Int).Mul(big.NewInt(2), leftG[1])                          // a1 = 2
	leftG[1] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(2), leftG[1]), order) // a1 = 2

	v[1] = new(bn256.G1).ScalarMult(g1, leftG[1]) // E(a1 * v1(s))

	leftG = append(
		leftG,
		Interpolate(
			s, []int64{1},
			BasisPolynomial(order, 2, []*big.Int{r, r1, r2, s1, s2}...),
		), // v2(s)
	)

	// leftG[2] = new(big.Int).Mul(big.NewInt(6), leftG[2])                          // a2 = 6
	leftG[2] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(6), leftG[2]), order) // a2 = 6

	v[2] = new(bn256.G1).ScalarMult(g1, leftG[2]) // E(a2 * v2(s))

	var w [3]*bn256.G2
	var rightG []*big.Int

	rightG = append(
		rightG,
		Interpolate(
			s, []int64{1, 1},
			BasisPolynomial(order, 1, []*big.Int{r, r1, r2, s1, s2}...),
			BasisPolynomial(order, 2, []*big.Int{r, r1, r2, s1, s2}...),
		), // w0(s)
	)

	// rightG[0] = new(big.Int).Mul(big.NewInt(1), rightG[0])
	rightG[0] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(1), rightG[0]), order)

	w[0] = new(bn256.G2).ScalarMult(g2, rightG[0]) // E(w0(s))

	rightG = append(
		rightG,
		Interpolate(
			s, []int64{1, 1},
			BasisPolynomial(order, 0, []*big.Int{r, r1, r2, s1, s2}...),
			BasisPolynomial(order, 3, []*big.Int{r, r1, r2, s1, s2}...),
		), // w1(s)
	)

	// rightG[1] = new(big.Int).Mul(big.NewInt(2), rightG[1])                          // a1 = 2
	rightG[1] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(2), rightG[1]), order) // a1 = 2

	w[1] = new(bn256.G2).ScalarMult(g2, rightG[1]) // E(a1 * w1(s))

	rightG = append(
		rightG,
		Interpolate(
			s, []int64{1},
			BasisPolynomial(order, 4, []*big.Int{r, r1, r2, s1, s2}...),
		), // w2(s)
	)

	// rightG[2] = new(big.Int).Mul(big.NewInt(6), rightG[2])                          // a2 = 6
	rightG[2] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(6), rightG[2]), order) // a2 = 6

	w[2] = new(bn256.G2).ScalarMult(g2, rightG[2]) // E(a2 * w2(s))

	var y [3]*bn256.G2
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
		Interpolate(
			s, []int64{1, 1},
			BasisPolynomial(order, 1, []*big.Int{r, r1, r2, s1, s2}...),
			BasisPolynomial(order, 3, []*big.Int{r, r1, r2, s1, s2}...),
		), // y1(s)
	)

	// outputG[1] = new(big.Int).Mul(big.NewInt(2), outputG[1])                          // a1 = 2
	outputG[1] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(2), outputG[1]), order) // a1 = 2

	y[1] = new(bn256.G2).ScalarMult(g2, outputG[1]) // E(a1 * y1(s))

	outputG = append(
		outputG,
		Interpolate(
			s, []int64{1, 1, 1},
			BasisPolynomial(order, 0, []*big.Int{r, r1, r2, s1, s2}...),
			BasisPolynomial(order, 2, []*big.Int{r, r1, r2, s1, s2}...),
			BasisPolynomial(order, 4, []*big.Int{r, r1, r2, s1, s2}...),
		), // y2(s)
	)

	// outputG[2] = new(big.Int).Mul(big.NewInt(6), outputG[2])                          // a2 = 6
	outputG[2] = new(big.Int).Mod(new(big.Int).Mul(big.NewInt(6), outputG[2]), order) // a2 = 6

	y[2] = new(bn256.G2).ScalarMult(g2, outputG[2]) // E(a2 * y2(s))

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

	// var t = new(big.Int).Mul(
	// 	new(big.Int).Sub(s, r),

	// 	new(big.Int).Mul(
	// 		new(big.Int).Mul(
	// 			new(big.Int).Sub(s, r1),
	// 			new(big.Int).Sub(s, r2),
	// 		),
	// 		new(big.Int).Mul(
	// 			new(big.Int).Sub(s, s1),
	// 			new(big.Int).Sub(s, s2),
	// 		),
	// 	),
	// )

	var t = new(big.Int).Mod(
		new(big.Int).Mul(
			new(big.Int).Sub(s, r),
			new(big.Int).Mul(
				new(big.Int).Mul(
					new(big.Int).Sub(s, r1),
					new(big.Int).Sub(s, r2),
				),
				new(big.Int).Mul(
					new(big.Int).Sub(s, s1),
					new(big.Int).Sub(s, s2),
				),
			),
		),
		order,
	)

	// var h = new(big.Int).Mul(
	// 	new(big.Int).Sub(
	// 		new(big.Int).Mul(term1, term2), term3,
	// 	),
	// 	new(big.Int).ModInverse(t, order),
	// )

	var h = new(big.Int).Mod(
		new(big.Int).Mul(
			new(big.Int).Sub(
				new(big.Int).Mul(term1, term2), term3,
			),
			new(big.Int).ModInverse(t, order),
		),
		order,
	)

	// Quadratic root detection to validate the SNARK was constructed with
	// values that satisfy the arithmetic circuit.

	var eV = new(bn256.G1).Add(v[0], new(bn256.G1).Add(v[1], v[2]))
	var eW = new(bn256.G2).Add(w[0], new(bn256.G2).Add(w[1], w[2]))
	var eY = bn256.Pair(g1, new(bn256.G2).Add(y[0], new(bn256.G2).Add(y[1], y[2])))

	var eT = new(bn256.G1).ScalarMult(g1, t)
	var eH = new(bn256.G2).ScalarMult(g2, h)

	var left = new(bn256.GT).Add(
		bn256.Pair(eV, eW),
		new(bn256.GT).Neg(eY),
	)

	var right = bn256.Pair(eT, eH)

	return bytes.Equal(left.Marshal(), right.Marshal())
}

// E1R1CS defines a R1CS that simplifies deriving the constraints for creating
// the QAP.
func E1R1CS() bool {

	// 3 * x1 = x2

	// a1 = 2, a2 = 6

	// s =

	//  [1, 2, 6]

	// A =

	// [3, 0, 0]

	// B =

	// [0, 1, 0]

	// C =

	// [0, 0, 1]

	// (A . s) * (B . s) = (3 * 1) * (1 * a1)
	//                   = (3 * 1) * (1 * 2)
	//                   = 6
	//                   = (C . s)

	return true
}
