package qap

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"

	// "golang.org/x/crypto/bn256"
	"github.com/cloudflare/bn256"

	// "github.com/ethereum/go-ethereum/crypto/bn256"
	// bn256 "github.com/ethereum/go-ethereum/crypto/bn256/google"

	// "github.com/ing-bank/zkrp/crypto/bn256"

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

	// bn256.Order.Set(big.NewInt(11))
	// bn256.Order.Set(big.NewInt(997))
	// bn256.Order.Set(big.NewInt(999983))

	var order = bn256.Order
	// var order = bn256.Order.Set(big.NewInt(11))
	// var order = bn256.Order.Set(big.NewInt(997))
	fmt.Println(order)

	var g1 *bn256.G1
	if _, g1, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var g2 *bn256.G2
	if _, g2, err = bn256.RandomG2(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var r1 *big.Int
	// if r1, err = rand.Prime(rand.Reader, order.BitLen()); err != nil {
	if r1, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}
	r1 = big.NewInt(1)
	// r1 = big.NewInt(3)
	// fmt.Println(r1)

	var r2 *big.Int
	// if r2, err = rand.Prime(rand.Reader, order.BitLen()); err != nil {
	if r2, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}
	r2 = big.NewInt(2)
	// r2 = big.NewInt(5)
	// fmt.Println(r2)

	var s1 *big.Int
	// if s1, err = rand.Prime(rand.Reader, order.BitLen()); err != nil {
	if s1, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}
	s1 = big.NewInt(3)
	// s1 = big.NewInt(7)
	// fmt.Println(s1)

	var s2 *big.Int
	// if s2, err = rand.Prime(rand.Reader, order.BitLen()); err != nil {
	if s2, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}
	s2 = big.NewInt(4)
	// s2 = big.NewInt(11)
	// fmt.Println(s2)

	var r *big.Int
	// if r, err = rand.Prime(rand.Reader, order.BitLen()); err != nil {
	if r, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}
	// r = big.NewInt(10)
	// r = big.NewInt(13)
	fmt.Println(r)

	var s *big.Int
	// if s, err = rand.Prime(rand.Reader, order.BitLen()); err != nil {
	if s, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}
	// s = big.NewInt(22)
	fmt.Println(s)

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

	// Data points: (r, 3), (s1, 1), (s2, 1)
	leftG = append(
		leftG,
		lagrange.Interpolate(
			s, []int64{3, 1, 1},
			BasisPolynomial(order, 0, []*big.Int{r, r1, r2, s1, s2}...),
			BasisPolynomial(order, 3, []*big.Int{r, r1, r2, s1, s2}...),
			BasisPolynomial(order, 4, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	leftG[0] = new(big.Int).Mul(big.NewInt(1), leftG[0])
	v[0] = new(bn256.G1).ScalarMult(g1, leftG[0])

	leftG = append(
		leftG,
		lagrange.Interpolate(
			s, []int64{1},
			lagrange.BasisPolynomial(1, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	leftG[1] = new(big.Int).Mul(big.NewInt(2), leftG[1])
	v[1] = new(bn256.G1).ScalarMult(g1, leftG[1])

	leftG = append(
		leftG,
		lagrange.Interpolate(
			s, []int64{1},
			lagrange.BasisPolynomial(2, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	leftG[2] = new(big.Int).Mul(big.NewInt(6), leftG[2])
	v[2] = new(bn256.G1).ScalarMult(g1, leftG[2])

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

	rightG = append(
		rightG,
		lagrange.Interpolate(
			s, []int64{1, 1},
			BasisPolynomial(order, 0, []*big.Int{r, r1, r2, s1, s2}...),
			BasisPolynomial(order, 3, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	rightG[1] = new(big.Int).Mul(big.NewInt(2), rightG[1])
	w[1] = new(bn256.G2).ScalarMult(g2, rightG[1])

	rightG = append(
		rightG,
		lagrange.Interpolate(
			s, []int64{1},
			BasisPolynomial(order, 4, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	rightG[2] = new(big.Int).Mul(big.NewInt(6), rightG[2])
	w[2] = new(bn256.G2).ScalarMult(g2, rightG[2])

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
		lagrange.Interpolate(
			s, []int64{1, 1},
			BasisPolynomial(order, 1, []*big.Int{r, r1, r2, s1, s2}...),
			BasisPolynomial(order, 3, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	outputG[1] = new(big.Int).Mul(big.NewInt(2), outputG[1])
	y[1] = new(bn256.G2).ScalarMult(g2, outputG[1])

	outputG = append(
		outputG,
		lagrange.Interpolate(
			s, []int64{1, 1, 1},
			BasisPolynomial(order, 0, []*big.Int{r, r1, r2, s1, s2}...),
			BasisPolynomial(order, 2, []*big.Int{r, r1, r2, s1, s2}...),
			BasisPolynomial(order, 4, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	outputG[2] = new(big.Int).Mul(big.NewInt(6), outputG[2])
	y[2] = new(bn256.G2).ScalarMult(g2, outputG[2])

	var term1 = new(big.Int).Add(
		leftG[0],
		new(big.Int).Add(
			leftG[1],
			leftG[2],
		),
	)

	fmt.Printf(" = = term1: %d \n", new(big.Int).Mod(term1, order))

	var term2 = new(big.Int).Add(
		rightG[0],
		new(big.Int).Add(
			rightG[1],
			rightG[2],
		),
	)

	fmt.Printf(" = = term2: %d \n", new(big.Int).Mod(term2, order))

	var term3 = new(big.Int).Add(
		outputG[0],
		new(big.Int).Add(
			outputG[1],
			outputG[2],
		),
	)

	fmt.Printf(" = = term3: %d \n", new(big.Int).Mod(term3, order))

	var t = new(big.Int).Mul(
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
	)

	fmt.Printf(" = = t: %d \n", new(big.Int).Mod(t, order))

	var h = new(big.Int).Mul(
		new(big.Int).Sub(
			new(big.Int).Mul(term1, term2), term3,
		),
		new(big.Int).ModInverse(t, order),
	)

	fmt.Printf(" = = h: %d \n", new(big.Int).Mod(h, order))

	var eV = new(bn256.G1).Add(v[0], new(bn256.G1).Add(v[1], v[2]))
	var eW = new(bn256.G2).Add(w[0], new(bn256.G2).Add(w[1], w[2]))
	var eY = bn256.Pair(g1, new(bn256.G2).Add(y[0], new(bn256.G2).Add(y[1], y[2])))

	var eT = new(bn256.G1).ScalarMult(g1, t)
	var eH = new(bn256.G2).ScalarMult(g2, h)

	var left = new(bn256.GT).Add(bn256.Pair(eV, eW), new(bn256.GT).Neg(eY))
	// var left = new(bn256.GT).Add(bn256.Pair(eV, eW), new(bn256.GT).ScalarMult(eY, big.NewInt(-1)))

	var right = bn256.Pair(eT, eH)

	// var res = new(bn256.GT).Add(left, new(bn256.GT).Neg(right))
	// fmt.Printf("parameter generation %v", res)

	// fmt.Printf(
	// 	" (v0 + a1 * v1 + a2 * v2) * (w0 + a1 * w1 + a2 * w2) - (y0 + a1 * y1 + a2 * y2) = (%s) * (%s) - %s \n",
	// 	eV.String()[0:18], eW.String()[0:18], eY.String()[0:18],
	// )

	// _, _ = left.Unmarshal(left.Marshal())
	// _, _ = right.Unmarshal(right.Marshal())

	// return bytes.Equal(left.Marshal(), right.Marshal())

	return bytes.Equal(left.Marshal(), right.Marshal())
}

// E1R1CS generates the Quadratic Arithmetic Program to validate arithmetic circuits in Zero Knowledge
func E1R1CS() bool {

	var err error

	var order = bn256.Order
	// var order = bn256.Order.Set(big.NewInt(11))
	// var order = bn256.Order.Set(big.NewInt(23))
	fmt.Println(order)

	// <[3, 0, 0] , [1, 2, 6]> * <[0, 1, 0] , [1, 2, 6]> - <[0, 0, 1] , [1, 2, 6]> = 3 * 2 - 6 = 0

	// A =

	// [3, 0, 0]

	// B =

	// [0, 1, 0]

	// C =

	// [0, 0, 1]

	// Two constraints:

	// <[3, 0, 0, 0, 0] , [1, 2, 6, 6, 1]> * <[0, 1, 0, 0, 0] , [1, 2, 6, 6, 1]> - <[0, 0, 1, 0, 0] , [1, 2, 6, 6, 1]> = 3 * 2 - 6 = 0
	// <[0, 0, 0, 0, 1] , [1, 2, 6, 6, 1]> * <[0, 0, 0, 1, 0] , [1, 2, 6, 6, 1]> - <[0, 0, 0, 1, 0] , [1, 2, 6, 6, 1]> = 1 * 6 - 6 = 0

	// A =

	// [3, 0, 0, 0, 0]
	// [0, 0, 0, 0, 1]

	// B =

	// [0, 1, 0, 0, 0]
	// [0, 0, 0, 1, 0]

	// C =

	// [0, 0, 1, 0, 0]
	// [0, 0, 0, 1, 0]

	var alpha1 *big.Int
	// if r1, err = rand.Prime(rand.Reader, order.BitLen()); err != nil {
	if alpha1, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}
	alpha1 = big.NewInt(1)
	fmt.Println(alpha1)

	var alpha2 *big.Int
	// if r1, err = rand.Prime(rand.Reader, order.BitLen()); err != nil {
	if alpha2, err = rand.Int(rand.Reader, order); err != nil {
		fmt.Printf("parameter generation %v", err)
	}
	alpha2 = big.NewInt(2)
	fmt.Println(alpha2)

	// Data points: (alpha1, 3), (alpha1, 0), (alpha1, 0)

	//
	// A(x) =
	//

	//                   (x - alpha2)
	// A_{0}(x) = 3 * -------------------
	//                 (alpha1 - alpha2)

	//
	// A_{1}(x) = 0
	//

	//
	// A_{2}(x) = 0
	//

	//
	// A_{3}(x) = 0
	//

	//                    (x - alpha1)
	// A_{4}(x) = 1 * -------------------
	//                 (alpha1 - alpha2)

	//
	// B(x) =
	//

	//
	// B_{0}(x) = 0
	//

	//                   x
	// B_{1}(x) = 1 * --------
	//                 alpha1

	//
	// B_{2}(x) = 0
	//

	//
	// B_{3}(x) = 0
	//

	//
	// B_{4}(x) = 0
	//

	//
	// C(x) =
	//

	//
	// C_{0}(x) = 0
	//

	//
	// C_{1}(x) = 0
	//

	//                   x
	// C_{2}(x) = 1 * --------
	//                 alpha1

	//                   x
	// C_{3}(x) = 1 * --------
	//                 alpha1

	//                   x
	// C_{4}(x) = 1 * --------
	//                 alpha1

	// A polynomials

	// [0, 3 * (3/alpha)]
	// [0,         0]
	// [0,         0]
	// [0,          ]
	// [0,          ]

	// B polynomials

	// [0,         0]
	// [0, (1/alpha)]
	// [0,         0]

	// C polynomials

	// [0,         0]
	// [0,         0]
	// [0, (1/alpha)]

	// A results at x = alpha1: [3, 0, 0]
	// B results at x = alpha1: [0, 1, 0]
	// C results at x = alpha1: [0, 0, 1]

	// 3 * 2 - 6 = 0

	// TODO: Check that the Lagrange interpolations evaluate to the correct
	// control points here:

	return true
}
