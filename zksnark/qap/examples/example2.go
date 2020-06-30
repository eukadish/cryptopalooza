package examples

import (
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

// t(x) = x - r

// v_{0}(x) = c_{0} = 3
// v_{1}(x) = c_{1} = 0
// v_{2}(x)         = 0

// w_{0}(x) = d_{0} = 0
// w_{1}(x) = d_{1} = 1
// w_{2}(x)         = 0

// y_{0}(x)         = 0
// y_{1}(x)         = 0
// y_{2}(x)         = 1

// m = 2
// d = 1

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

// LagrangeBase represents the data that will be needed interpolate a
// polynomial using the interpolation polynomial in the Lagrange form:
// https://en.wikipedia.org/wiki/Lagrange_polynomial

type LagrangeBase struct {
	xCoords []*big.Int
}

func (lag LagrangeBase) generatePolynomial(j int) func(*big.Int) *big.Int {

	// TODO: Return an error if the index is out of range

	var xCoord *big.Int
	var denominator *big.Int

	for _, xCoord = range append(lag.xCoords[:j], lag.xCoords[j+1:]...) {
		denominator = new(big.Int).Mul(denominator, new(big.Int).Sub(lag.xCoords[j], xCoord))
	}

	return func(x *big.Int) *big.Int {

		var xCoord *big.Int
		var numerator *big.Int

		for _, xCoord = range append(lag.xCoords[:j], lag.xCoords[j+1:]...) {
			numerator = new(big.Int).Mul(denominator, new(big.Int).Sub(x, xCoord))
		}

		return new(big.Int).Mul(numerator, new(big.Int).ModInverse(denominator, bn256.Order))
	}
}

// New creates a new Lagrange base representation for generating basis
// polynomials
func New(xCoords ...*big.Int) (lag LagrangeBase) {

	var xCoord *big.Int
	for _, xCoord = range xCoords {
		lag.xCoords = append(lag.xCoords, xCoord)
	}

	// TODO: Generate the functions to evaluate the basis here

	return
}

// LagrangeInterpolation loops through the basis polnomials to evaluate the function at a point
func LagrangeInterpolation(
	x *big.Int,
	yCoords []int64,
	basis ...func(*big.Int) *big.Int,
) *big.Int {

	var accumulator = big.NewInt(1)

	var index int
	var base func(*big.Int) *big.Int

	for index, base = range basis {
		accumulator = new(big.Int).Add(accumulator, new(big.Int).Mul(big.NewInt(yCoords[index]), base(x)))
	}

	return accumulator
}

// Pi data used by the prover to share with the verifier
type Pi struct {
	vmid, vpmid *bn256.G1

	w, wp *bn256.G2

	y, h, yp, hp, z *bn256.GT
}

// Example2 generates the Quadratic Arithmetic Program to validate arithmetic circuits in Zero Knowledge
func Example2() bool {

	var err error

	var sk1 *big.Int
	if sk1, _, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var sk2 *big.Int
	if sk2, _, err = bn256.RandomG2(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var skt *big.Int
	if skt, _, err = bn256.RandomGT(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var pk1 *bn256.G1
	if pk1 = new(bn256.G1).ScalarBaseMult(sk1); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var pk2 *bn256.G2
	if pk2 = new(bn256.G2).ScalarBaseMult(sk2); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var pkt *bn256.GT
	if pkt = new(bn256.GT).ScalarBaseMult(skt); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var r1 *big.Int
	if r1, _, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var r2 *big.Int
	if r2, _, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var s1 *big.Int
	if s1, _, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var s2 *big.Int
	if s2, _, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var r *big.Int
	if r, _, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var s *big.Int
	if s, _, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var alpha *big.Int
	if alpha, _, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var betaV *big.Int
	if betaV, _, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var betaW *big.Int
	if betaW, _, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var betaY *big.Int
	if betaY, _, err = bn256.RandomG1(rand.Reader); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var prover Pi

	// var prover []*bigInt

	var leftG []*big.Int
	// var y *big.Int

	// var v0 = LagrangeInterpolation(
	// 	s, []int64{1, 1, 3},
	// 	lagrange.BasisPolynomial(0, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(0)
	// 	lagrange.BasisPolynomial(3, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(3)
	// 	lagrange.BasisPolynomial(4, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(4)
	// )

	leftG = append(
		leftG,
		LagrangeInterpolation(
			s, []int64{1, 1, 3},
			lagrange.BasisPolynomial(0, []*big.Int{r, r1, r2, s1, s2}...),
			lagrange.BasisPolynomial(3, []*big.Int{r, r1, r2, s1, s2}...),
			lagrange.BasisPolynomial(4, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	// var v1 = LagrangeInterpolation(
	// 	s, []int64{1},
	// 	lagrange.BasisPolynomial(1, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(0)
	// )

	leftG = append(
		leftG,
		LagrangeInterpolation(
			s, []int64{1},
			lagrange.BasisPolynomial(1, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	// var v2 = LagrangeInterpolation(
	// 	s, []int64{1},
	// 	lagrange.BasisPolynomial(2, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(0)
	// )

	leftG = append(
		leftG,
		LagrangeInterpolation(
			s, []int64{1},
			lagrange.BasisPolynomial(2, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	// var pi0 *bn256.G1
	// if pi0 = new(bn256.G1).ScalarMult(pk, v[0]); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	// var pi1 *bn256.G1
	// if pi1 = new(bn256.G1).ScalarMult(pk, v[1]); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	// var pi2 *bn256.G1
	// if pi2 = new(bn256.G1).ScalarMult(pk, v[2]); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	// fmt.Printf(" E(v_{0}(s))      = %s \n", pi0.String()[0:18])
	// fmt.Printf(" E(v_{1}(s))      = %s \n", pi1.String()[0:18])
	// fmt.Printf(" E(v_{2}(s))      = %s \n", pi2.String()[0:18])

	var rightG []*big.Int

	// var w0 = LagrangeInterpolation(
	// 	s, []int64{1, 1},
	// 	lagrange.BasisPolynomial(1, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(0)
	// 	lagrange.BasisPolynomial(2, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(3)
	// )

	rightG = append(
		rightG,
		LagrangeInterpolation(
			s, []int64{1, 1},
			lagrange.BasisPolynomial(1, []*big.Int{r, r1, r2, s1, s2}...),
			lagrange.BasisPolynomial(2, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	// var w1 = LagrangeInterpolation(
	// 	s, []int64{1, 1},
	// 	lagrange.BasisPolynomial(0, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(0)
	// 	lagrange.BasisPolynomial(3, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(0)
	// )

	rightG = append(
		rightG,
		LagrangeInterpolation(
			s, []int64{1, 1},
			lagrange.BasisPolynomial(0, []*big.Int{r, r1, r2, s1, s2}...),
			lagrange.BasisPolynomial(3, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	// var w2 = LagrangeInterpolation(
	// 	s, []int64{1},
	// 	lagrange.BasisPolynomial(4, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(0)
	// )

	rightG = append(
		rightG,
		LagrangeInterpolation(
			s, []int64{1},
			lagrange.BasisPolynomial(4, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	var outputG []*big.Int

	// var y0 = LagrangeInterpolation(
	// 	s, []int64{1, 1},
	// 	lagrange.BasisPolynomial(1, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(0)
	// 	lagrange.BasisPolynomial(2, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(3)
	// )

	// var y0 = big.NewInt(0)

	outputG = append(
		outputG,
		big.NewInt(0),
	)

	// var y1 = LagrangeInterpolation(
	// 	s, []int64{1, 1},
	// 	lagrange.BasisPolynomial(1, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(0)
	// 	lagrange.BasisPolynomial(3, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(0)
	// )

	outputG = append(
		outputG,
		LagrangeInterpolation(
			s, []int64{1, 1},
			lagrange.BasisPolynomial(1, []*big.Int{r, r1, r2, s1, s2}...),
			lagrange.BasisPolynomial(3, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	// var y2 = LagrangeInterpolation(
	// 	s, []int64{1, 1, 1},
	// 	lagrange.BasisPolynomial(0, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(0)
	// 	lagrange.BasisPolynomial(2, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(0)
	// 	lagrange.BasisPolynomial(4, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(0)
	// )

	outputG = append(
		outputG,
		LagrangeInterpolation(
			s, []int64{1, 1, 1},
			lagrange.BasisPolynomial(0, []*big.Int{r, r1, r2, s1, s2}...),
			lagrange.BasisPolynomial(2, []*big.Int{r, r1, r2, s1, s2}...),
			lagrange.BasisPolynomial(4, []*big.Int{r, r1, r2, s1, s2}...),
		),
	)

	// Prover = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

	var vmid = new(big.Int).Mul(big.NewInt(6), leftG[2])

	if prover.vmid = new(bn256.G1).ScalarMult(pk1, vmid); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var w = new(big.Int).Add(
		new(big.Int).Mul(big.NewInt(2), rightG[1]),
		new(big.Int).Mul(big.NewInt(6), rightG[2]),
	)

	if prover.w = new(bn256.G2).ScalarMult(pk2, w); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var y = new(big.Int).Add(
		new(big.Int).Mul(big.NewInt(2), outputG[1]),
		new(big.Int).Mul(big.NewInt(6), outputG[2]),
	)

	if prover.y = new(bn256.GT).ScalarMult(pkt, y); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	// TODO: h(s)

	//
	//
	//

	if prover.vpmid = new(bn256.G1).ScalarMult(pk1, new(big.Int).Mul(alpha, vmid)); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	if prover.wp = new(bn256.G2).ScalarMult(pk2, new(big.Int).Mul(alpha, w)); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	if prover.yp = new(bn256.GT).ScalarMult(pkt, new(big.Int).Mul(alpha, y)); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	// TODO: alpha * h(s)

	//
	//
	//

	var z = new(big.Int).Add(
		new(big.Int).Mul(betaV, vmid),
		new(big.Int).Add(
			new(big.Int).Mul(betaW, w),
			new(big.Int).Mul(betaY, y),
		),
	)

	if prover.z = new(bn256.GT).ScalarMult(pkt, z); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var v0 *bn256.G1
	if v0 = new(bn256.G1).ScalarMult(pk1, leftG[0]); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var vin *bn256.G1
	if vin = new(bn256.G1).ScalarMult(pk1, new(big.Int).Mul(big.NewInt(2), leftG[1])); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var w0 *bn256.G2
	if w0 = new(bn256.G2).ScalarMult(pk2, rightG[0]); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var y0 *bn256.GT
	if y0 = new(bn256.GT).ScalarMult(pkt, outputG[0]); err != nil {
		fmt.Printf("parameter generation %v", err)
	}

	var left = new(bn256.G1).Add(new(bn256.G1).Add(v0, vin), prover.vmid)
	var right = new(bn256.G2).Add(w0, prover.w)

	var result = new(bn256.GT).Add(
		bn256.Pair(left, right),
		new(bn256.GT).Neg(new(bn256.GT).Add(y0, prover.y)),
	)

	// get h

	// new(bn256.GT).Add(y0, prover.y)

	// bn256.Pair(left, right)

	// fmt.Println(bytes.Equal(bn256.Pair(left, right).Marshal(), )

	// fmt.Println(bn256.Pair(left, right).Marshal())
	fmt.Printf(" ( . . .) * ( . . . ) - ( . . . ) - ( . . . ) = %s \n", result.String()[0:18])

	// var w *big.Int

	// var w0 = LagrangeInterpolation(
	// 	s, []int64{1, 1},
	// 	lagrange.BasisPolynomial(1, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(0)
	// 	lagrange.BasisPolynomial(2, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(3)
	// )

	// w = append(w, w0)

	// var w1 = LagrangeInterpolation(
	// 	s, []int64{1, 1},
	// 	lagrange.BasisPolynomial(0, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(0)
	// 	lagrange.BasisPolynomial(3, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(0)
	// )

	// w = append(w, w1)

	// var w2 = LagrangeInterpolation(
	// 	s, []int64{1},
	// 	lagrange.BasisPolynomial(4, []*big.Int{r, r1, r2, s1, s2}...), // lag.generatePolynomial(0)
	// )

	// w = append(w, w2)

	// var pi0 *bn256.G1
	// if pi0 = new(bn256.G1).ScalarMult(pk, w[0]); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	// var pi1 *bn256.G1
	// if pi1 = new(bn256.G1).ScalarMult(pk, w[1]); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	// var pi2 *bn256.G1
	// if pi2 = new(bn256.G1).ScalarMult(pk, w[2]); err != nil {
	// 	fmt.Printf("parameter generation %v", err)
	// }

	// fmt.Printf(" E(v_{0}(s))      = %s \n", pi0.String()[0:18])
	// fmt.Printf(" E(v_{1}(s))      = %s \n", pi1.String()[0:18])
	// fmt.Printf(" E(v_{2}(s))      = %s \n", pi2.String()[0:18])

	// fmt.Printf(" E(v_{0}(s)) * E(v_{1}(s))      = %s \n", new(bn256.G1).Add(pi0, pi1).String()[0:18])

	// CRS generation

	// 	var alpha *big.Int
	// 	if alpha, _, err = bn256.RandomG1(rand.Reader); err != nil {
	// 		log.Printf("parameter generation %v", err)
	// 	}

	// 	var crs1 *bn256.G1
	// 	if crs1 = new(bn256.G1).ScalarMult(pk, big.NewInt(1)); err != nil { // Encrypt s^{0}
	// 		log.Printf("parameter generation %v", err)
	// 	}

	// 	var crs2 *bn256.G1
	// 	if crs2 = new(bn256.G1).ScalarMult(pk, s); err != nil { // Encrypt s^{1}
	// 		log.Printf("parameter generation %v", err)
	// 	}

	// 	var crs3 *bn256.G1
	// 	if crs3 = new(bn256.G1).ScalarMult(pk, alpha); err != nil { // Encrypt alpha
	// 		log.Printf("parameter generation %v", err)
	// 	}

	// 	var crs4 *bn256.G1
	// 	if crs4 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(alpha, s)); err != nil { // Encrypt alpha * s
	// 		log.Printf("parameter generation %v", err)
	// 	}

	// 	log.Printf("crs1: %s \n", crs1.String()[0:18])
	// 	log.Printf("crs2: %s \n", crs2.String()[0:18])
	// 	log.Printf("crs3: %s \n", crs3.String()[0:18])
	// 	log.Printf("crs4: %s \n", crs4.String()[0:18])

	// 	// Function specific CRS generation

	// 	var betaV *big.Int
	// 	if betaV, _, err = bn256.RandomG1(rand.Reader); err != nil {
	// 		log.Printf("parameter generation %v", err)
	// 	}

	// 	var betaW *big.Int
	// 	if betaW, _, err = bn256.RandomG1(rand.Reader); err != nil {
	// 		log.Printf("parameter generation %v", err)
	// 	}

	// 	var betaY *big.Int
	// 	if betaY, _, err = bn256.RandomG1(rand.Reader); err != nil {
	// 		log.Printf("parameter generation %v", err)
	// 	}

	// 	var gamma *big.Int
	// 	if gamma, _, err = bn256.RandomG1(rand.Reader); err != nil {
	// 		log.Printf("parameter generation %v", err)
	// 	}

	// 	fmt.Printf("betaV: %s \n", betaV.String()[0:12])
	// 	fmt.Printf("betaW: %s \n", betaW.String()[0:12])
	// 	fmt.Printf("betaY: %s \n", betaY.String()[0:12])

	// 	fmt.Printf("gamma: %s \n", gamma.String()[0:12])

	// 	// crs-f - = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

	// 	// crs

	// 	// I_{free} = {}

	// 	// I_{labeled} = I_{10} U I_{11} = {1, 2}

	// 	// I_{in} = {1}

	// 	// I_{mid} = {2}

	// 	// E(v_{2}(s)) =

	// 	var term1, term2, term3, term4 *big.Int
	// 	var numerator, inverse, denominator *big.Int

	// 	numerator = partialTerm(s, r, r1, s1, s2)
	// 	denominator = partialTerm(r2, r, r1, s1, s2)

	// 	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	// 	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	// 	var crsf1 *bn256.G1
	// 	if crsf1 = new(bn256.G1).ScalarMult(pk, term1); err != nil {
	// 		log.Printf("parameter generation %v", err)
	// 	}

	// 	// use multiplicative inverse check if equal :(
	// 	log.Printf("E(v_{2}(s) = %s \n", crsf1.String()[0:18])

	// 	// E(w_{1}(s)) =

	// 	numerator = partialTerm(s, r, r1, r2, s2)
	// 	denominator = partialTerm(s1, r, r1, r2, s2)

	// 	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	// 	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	// 	numerator = partialTerm(s, r1, r2, s1, s2)
	// 	denominator = partialTerm(r, r1, r2, s1, s2)

	// 	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	// 	term2 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	// 	var crsf2 *bn256.G1
	// 	if crsf2 = new(bn256.G1).ScalarMult(pk, new(big.Int).Add(term1, term2)); err != nil {
	// 		log.Printf("parameter generation %v", err)
	// 	}

	// 	// use multiplicative inverse check if equal :(
	// 	log.Printf("E(w_{1}(s) = %s \n", crsf2.String()[0:18])

	// 	// E(w_{2}(s)) =

	// 	numerator = partialTerm(s, r, r1, r2, s1)
	// 	denominator = partialTerm(s2, r, r1, r2, s1)

	// 	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	// 	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	// 	var crsf3 *bn256.G1
	// 	if crsf3 = new(bn256.G1).ScalarMult(pk, term1); err != nil {
	// 		log.Printf("parameter generation %v", err)
	// 	}

	// 	// use multiplicative inverse check if equal :(
	// 	log.Printf("E(w_{2}(s) = %s \n", crsf3.String()[0:18])

	// 	// E(y_{1}(s)) =

	// 	numerator = partialTerm(s, r, r2, s1, s2)
	// 	denominator = partialTerm(r1, r, r2, s1, s2)

	// 	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	// 	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	// 	numerator = partialTerm(s, r, r1, r2, s2)
	// 	denominator = partialTerm(s1, r, r1, r2, s2)

	// 	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	// 	term2 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	// 	var crsf4 *bn256.G1
	// 	if crsf4 = new(bn256.G1).ScalarMult(pk, new(big.Int).Add(term1, term2)); err != nil {
	// 		log.Printf("parameter generation %v", err)
	// 	}

	// 	// use multiplicative inverse check if equal :(
	// 	log.Printf("E(y_{1}(s) = %s \n", crsf4.String()[0:18])

	// 	// E(y_{2}(s)) =

	// 	numerator = partialTerm(s, r, r1, s1, s2)
	// 	denominator = partialTerm(r2, r, r1, s1, s2)

	// 	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	// 	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	// 	numerator = partialTerm(s, r, r1, r2, s1)
	// 	denominator = partialTerm(s2, r, r1, r2, s1)

	// 	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	// 	term2 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	// 	numerator = partialTerm(s, r1, r2, s1, s2)
	// 	denominator = partialTerm(r, r1, r2, s1, s2)

	// 	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	// 	term3 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	// 	term4 = new(big.Int).Add(term1, term2)

	// 	var crsf5 *bn256.G1
	// 	if crsf5 = new(bn256.G1).ScalarMult(pk, new(big.Int).Add(term3, term4)); err != nil {
	// 		log.Printf("parameter generation %v", err)
	// 	}

	// 	// use multiplicative inverse check if equal :(
	// 	log.Printf("E(y_{2}(s) = %s \n", crsf5.String()[0:18])

	return true
}

// construct the snark`

// use linear interpolation

// v_{0}(r) = c_{0} = 3
// v_{1}(r) = c_{1} = 0
// v_{2}(r)         = 0

// w_{0}(r) = d_{0} = 0
// w_{1}(r) = d_{1} = 1
// w_{2}(r)         = 0

// y_{0}(r)         = 0
// y_{1}(r)         = 0
// y_{2}(r)         = 1

// v_{0}(x) = 3 * x / r
// v_{1}(x) = 0
// v_{2}(x) = 0

// w_{0}(x) = 0
// w_{1}(x) = x / r
// w_{2}(x) = 0

// y_{0}(x) = 0
// y_{1}(x) = 0
// y_{2}(x) = x / r

// construct the snark

// P(x) = (3 * (x / r)) * (x / r) - (x / r)
//      = (3 * (x / r)) * (2 * (x / r)) - (6 * (x / r))
//      = 6 * (x / r) * (x / r - 1)
//      = (6 * (x / r^{2})) * (x - r)

//

func R1CSExample2() bool {

	return true
}
