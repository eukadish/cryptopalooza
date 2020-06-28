package examples

import (
	"crypto/rand"
	"log"
	"math/big"

	"github.com/cloudflare/bn256"
)

// type Poly func(x *big.Int) *big.Int

// func lagrangeInterp(r, c1, c2, c3, c4 *big.Int) *big.Int {

// 	var term1 = new(big.Int).Sub(r, c1)
// 	var term2 = new(big.Int).Sub(r, c2)
// 	var term3 = new(big.Int).Sub(r, c3)
// 	var term4 = new(big.Int).Sub(r, c4)

// 	var prod1 = new(big.Int).Mul(term1, term2)
// 	var prod2 = new(big.Int).Mul(term3, term4)

// 	var prod = new(big.Int).Mul(prod1, prod2)

// 	return new(big.Int).Mod(prod, bn256.Order)
// }

func partialTerm(x, x0, x1, x2, x3 *big.Int) *big.Int {

	// log.Printf("x: %d", x)
	// log.Printf("x0: %d", x0)
	// log.Printf("x1: %d", x1)
	// log.Printf("x2: %d", x2)
	// log.Printf("x3: %d", x3)

	var term1 = new(big.Int).Sub(x, x0)
	var term2 = new(big.Int).Sub(x, x1)
	var term3 = new(big.Int).Sub(x, x2)
	var term4 = new(big.Int).Sub(x, x3)

	return new(big.Int).Mul(
		new(big.Int).Mul(term1, term2),
		new(big.Int).Mul(term3, term4),
	)
}

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

// generate strong QAP for the SNARK

// LinearQAP generates the Quadratic Arithmetic Program to validate arithmetic circuits in Zero Knowledge
func LinearQAP() bool {

	var err error

	// var v, w, y []*big.Int

	var r1 *big.Int
	if r1, _, err = bn256.RandomG1(rand.Reader); err != nil {
		log.Printf("parameter generation %v", err)
	}

	var r2 *big.Int
	if r2, _, err = bn256.RandomG1(rand.Reader); err != nil {
		log.Printf("parameter generation %v", err)
	}

	var s1 *big.Int
	if s1, _, err = bn256.RandomG1(rand.Reader); err != nil {
		log.Printf("parameter generation %v", err)
	}

	var s2 *big.Int
	if s2, _, err = bn256.RandomG1(rand.Reader); err != nil {
		log.Printf("parameter generation %v", err)
	}

	var r *big.Int
	if r, _, err = bn256.RandomG1(rand.Reader); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// log.Printf("r: %d", r)

	// log.Printf("r1: %d", r1)
	// log.Printf("r2: %d", r2)

	// log.Printf("s1: %d", s1)
	// log.Printf("s2: %d", s2)

	// t(x) = x - r

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

	// CRS generation

	var alpha *big.Int
	if alpha, _, err = bn256.RandomG1(rand.Reader); err != nil {
		log.Printf("parameter generation %v", err)
	}

	var s *big.Int
	if s, _, err = bn256.RandomG1(rand.Reader); err != nil {
		log.Printf("parameter generation %v", err)
	}

	var sk *big.Int
	if sk, _, err = bn256.RandomG1(rand.Reader); err != nil {
		log.Printf("parameter generation %v", err)
	}

	log.Printf(" = = crs = = = = = = = = = = = = = = = = = \n \n")

	var pk *bn256.G1
	if pk = new(bn256.G1).ScalarBaseMult(sk); err != nil {
		log.Printf("parameter generation %v", err)
	}

	var crs1 *bn256.G1
	if crs1 = new(bn256.G1).ScalarMult(pk, big.NewInt(1)); err != nil { // Encrypt s^{0}
		log.Printf("parameter generation %v", err)
	}

	log.Printf(" E(1)             = %s \n", crs1.String()[0:18])

	var crs2 *bn256.G1
	if crs2 = new(bn256.G1).ScalarMult(pk, s); err != nil { // Encrypt s^{1}
		log.Printf("parameter generation %v", err)
	}

	log.Printf(" E(s)             = %s \n", crs2.String()[0:18])

	var crs3 *bn256.G1
	if crs3 = new(bn256.G1).ScalarMult(pk, alpha); err != nil { // Encrypt alpha
		log.Printf("parameter generation %v", err)
	}

	log.Printf(" E(\u03B1)             = %s \n", crs3.String()[0:18])

	var crs4 *bn256.G1
	if crs4 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(alpha, s)); err != nil { // Encrypt alpha * s
		log.Printf("parameter generation %v", err)
	}

	log.Printf(" E(\u03B1 * s)         = %s \n \n", crs4.String()[0:18])

	// Function specific CRS generation

	var betaV *big.Int
	if betaV, _, err = bn256.RandomG1(rand.Reader); err != nil {
		log.Printf("parameter generation %v", err)
	}

	var betaW *big.Int
	if betaW, _, err = bn256.RandomG1(rand.Reader); err != nil {
		log.Printf("parameter generation %v", err)
	}

	var betaY *big.Int
	if betaY, _, err = bn256.RandomG1(rand.Reader); err != nil {
		log.Printf("parameter generation %v", err)
	}

	var gamma *big.Int
	if gamma, _, err = bn256.RandomG1(rand.Reader); err != nil {
		log.Printf("parameter generation %v", err)
	}

	log.Printf(" = = crs - f = = = = = = = = = = = = = = = \n \n")

	// I_{free} = {}

	// I_{labeled} = I_{10} U I_{11} = {1, 2}

	// I_{in} = {1}

	// I_{mid} = {2}

	// E(v_{2}(s)) =

	var numerator, inverse, denominator *big.Int
	var term1, term2, term3, term4, term5, term6, term7, term8 *big.Int

	numerator = partialTerm(s, r, r1, s1, s2)
	denominator = partialTerm(r2, r, r1, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	var crsf1 *bn256.G1
	if crsf1 = new(bn256.G1).ScalarMult(pk, term1); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(v_{2}(s))      = %s \n", crsf1.String()[0:18])

	// E(w_{1}(s)) =

	numerator = partialTerm(s, r, r1, r2, s2)
	denominator = partialTerm(s1, r, r1, r2, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	numerator = partialTerm(s, r1, r2, s1, s2)
	denominator = partialTerm(r, r1, r2, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term2 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	var crsf2 *bn256.G1
	if crsf2 = new(bn256.G1).ScalarMult(pk, new(big.Int).Add(term1, term2)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(w_{1}(s))      = %s \n", crsf2.String()[0:18])

	// E(w_{2}(s)) =

	numerator = partialTerm(s, r, r1, r2, s1)
	denominator = partialTerm(s2, r, r1, r2, s1)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	var crsf3 *bn256.G1
	if crsf3 = new(bn256.G1).ScalarMult(pk, term1); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(w_{2}(s))      = %s \n", crsf3.String()[0:18])

	// E(y_{1}(s)) =

	numerator = partialTerm(s, r, r2, s1, s2)
	denominator = partialTerm(r1, r, r2, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	numerator = partialTerm(s, r, r1, r2, s2)
	denominator = partialTerm(s1, r, r1, r2, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term2 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	var crsf4 *bn256.G1
	if crsf4 = new(bn256.G1).ScalarMult(pk, new(big.Int).Add(term1, term2)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(y_{1}(s))      = %s \n", crsf4.String()[0:18])

	// E(y_{2}(s)) =

	numerator = partialTerm(s, r, r1, s1, s2)
	denominator = partialTerm(r2, r, r1, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	numerator = partialTerm(s, r, r1, r2, s1)
	denominator = partialTerm(s2, r, r1, r2, s1)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term2 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	numerator = partialTerm(s, r1, r2, s1, s2)
	denominator = partialTerm(r, r1, r2, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term3 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	term4 = new(big.Int).Add(term1, term2)

	var crsf5 *bn256.G1
	if crsf5 = new(bn256.G1).ScalarMult(pk, new(big.Int).Add(term3, term4)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(y_{2}(s))      = %s \n", crsf5.String()[0:18])

	// E(1) =

	var crsf6 *bn256.G1
	if crsf6 = new(bn256.G1).ScalarMult(pk, big.NewInt(1)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(1)             = %s \n", crsf6.String()[0:18])

	// E(s) =

	var crsf7 *bn256.G1
	if crsf7 = new(bn256.G1).ScalarMult(pk, s); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(s)             = %s \n", crsf7.String()[0:18])

	// E(alpha * v_{2}(s)) =

	numerator = partialTerm(s, r, r1, s1, s2)
	denominator = partialTerm(r2, r, r1, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	var crsf8 *bn256.G1
	if crsf8 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(alpha, term1)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(\u03B1 * v_{2}(s))  = %s \n", crsf8.String()[0:18])

	// E(alpha * w_{1}(s)) =

	numerator = partialTerm(s, r, r1, r2, s2)
	denominator = partialTerm(s1, r, r1, r2, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	numerator = partialTerm(s, r1, r2, s1, s2)
	denominator = partialTerm(r, r1, r2, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term2 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	var crsf9 *bn256.G1
	if crsf9 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(alpha, term2)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(\u03B1 * w_{1}(s))  = %s \n", crsf9.String()[0:18])

	// E(alpha * w_{2}(s)) =

	numerator = partialTerm(s, r, r1, r2, s1)
	denominator = partialTerm(s2, r, r1, r2, s1)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	var crsf10 *bn256.G1
	if crsf10 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(alpha, term1)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(\u03B1 * w_{2}(s))  = %s \n", crsf10.String()[0:18])

	// E(alpha * y_{1}(s)) =

	numerator = partialTerm(s, r, r2, s1, s2)
	denominator = partialTerm(r1, r, r2, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	numerator = partialTerm(s, r, r1, r2, s2)
	denominator = partialTerm(s1, r, r1, r2, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term2 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	term3 = new(big.Int).Add(term1, term2)

	var crsf11 *bn256.G1
	if crsf11 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(alpha, term3)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(\u03B1 * y_{1}(s))  = %s \n", crsf11.String()[0:18])

	// E(alpha * y_{2}(s)) =

	numerator = partialTerm(s, r, r1, s1, s2)
	denominator = partialTerm(r2, r, r1, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	numerator = partialTerm(s, r, r1, r2, s1)
	denominator = partialTerm(s2, r, r1, r2, s1)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term2 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	numerator = partialTerm(s, r1, r2, s1, s2)
	denominator = partialTerm(r, r1, r2, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term3 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	term4 = new(big.Int).Add(term1, term2)
	term5 = new(big.Int).Add(term3, term4)

	var crsf12 *bn256.G1
	if crsf12 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(alpha, term5)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(\u03B1 * y_{2}(s))  = %s \n", crsf12.String()[0:18])

	// E(alpha) =

	var crsf13 *bn256.G1
	if crsf13 = new(bn256.G1).ScalarMult(pk, alpha); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(\u03B1)             = %s \n", crsf13.String()[0:18])

	// E(alpha * s) =

	var crsf14 *bn256.G1
	if crsf14 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(alpha, s)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(\u03B1 * s)         = %s \n", crsf14.String()[0:18])

	// E(betaV * v_{2}(s)) =

	numerator = partialTerm(s, r, r1, s1, s2)
	denominator = partialTerm(r2, r, r1, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	var crsf15 *bn256.G1
	if crsf15 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(betaV, term1)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(\u03B2v * v_{2}(s)) = %s \n", crsf15.String()[0:18])

	// E(betaW * w_{1}(s)) =

	numerator = partialTerm(s, r, r1, r2, s2)
	denominator = partialTerm(s1, r, r1, r2, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	numerator = partialTerm(s, r1, r2, s1, s2)
	denominator = partialTerm(r, r1, r2, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term2 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	var crsf16 *bn256.G1
	if crsf16 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(betaW, term2)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(\u03B2w * w_{1}(s)) = %s \n", crsf16.String()[0:18])

	// E(betaW * w_{2}(s)) =

	numerator = partialTerm(s, r, r1, r2, s1)
	denominator = partialTerm(s2, r, r1, r2, s1)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	var crsf17 *bn256.G1
	if crsf17 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(betaW, term1)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(\u03B2w * w_{2}(s)) = %s \n", crsf17.String()[0:18])

	// E(betaY * y_{1}(s)) =

	numerator = partialTerm(s, r, r2, s1, s2)
	denominator = partialTerm(r1, r, r2, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	numerator = partialTerm(s, r, r1, r2, s2)
	denominator = partialTerm(s1, r, r1, r2, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term2 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	term3 = new(big.Int).Add(term1, term2)

	var crsf18 *bn256.G1
	if crsf18 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(betaY, term3)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(\u03B2y * y_{1}(s)) = %s \n", crsf18.String()[0:18])

	// E(betaY * y_{2}(s)) =

	numerator = partialTerm(s, r, r1, s1, s2)
	denominator = partialTerm(r2, r, r1, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	numerator = partialTerm(s, r, r1, r2, s1)
	denominator = partialTerm(s2, r, r1, r2, s1)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term2 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	numerator = partialTerm(s, r1, r2, s1, s2)
	denominator = partialTerm(r, r1, r2, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term3 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	term4 = new(big.Int).Add(term1, term2)
	term5 = new(big.Int).Add(term3, term4)

	var crsf19 *bn256.G1
	if crsf19 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(betaY, term5)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(\u03B2y * y_{2}(s)) = %s \n \n", crsf19.String()[0:18])

	log.Printf(" = = shortcrs - f = = = = = = = = = = = = \n \n")

	// E(1) =

	var shortcrsf1 *bn256.G1
	if shortcrsf1 = new(bn256.G1).ScalarMult(pk, big.NewInt(1)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(1)             = %s \n", shortcrsf1.String()[0:18])

	// E(alpha) =

	var shortcrsf2 *bn256.G1
	if shortcrsf2 = new(bn256.G1).ScalarMult(pk, alpha); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(\u03B1)             = %s \n", shortcrsf2.String()[0:18])

	// E(gamma) =

	var shortcrsf3 *bn256.G1
	if shortcrsf3 = new(bn256.G1).ScalarMult(pk, gamma); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(\u03B3)             = %s \n", shortcrsf3.String()[0:18])

	// E(betaV * gamma) =

	var shortcrsf4 *bn256.G1
	if shortcrsf4 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(betaV, gamma)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(\u03B2v * \u03B3)        = %s \n", shortcrsf4.String()[0:18])

	// E(betaW * gamma)(s)) =

	var shortcrsf5 *bn256.G1
	if shortcrsf5 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(betaW, gamma)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(\u03B2w * \u03B3)        = %s \n", shortcrsf5.String()[0:18])

	// E(betaY * y_{1}(s)) =

	var shortcrsf6 *bn256.G1
	if shortcrsf6 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(betaY, gamma)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(\u03B2w * \u03B3)        = %s \n", shortcrsf6.String()[0:18])

	// E(v_{1}(s)) =

	numerator = partialTerm(s, r, r2, s1, s2)
	denominator = partialTerm(r1, r, r2, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	var shortcrsf7 *bn256.G1
	if shortcrsf7 = new(bn256.G1).ScalarMult(pk, term1); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(v_{1}(s))      = %s \n", shortcrsf7.String()[0:18])

	// E(w_{0}(s)) =

	numerator = partialTerm(s, r, r2, s1, s2)
	denominator = partialTerm(r1, r, r2, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	numerator = partialTerm(s, r, r1, s1, s2)
	denominator = partialTerm(r2, r, r1, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term2 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	var shortcrsf8 *bn256.G1
	if shortcrsf8 = new(bn256.G1).ScalarMult(pk, new(big.Int).Add(term1, term2)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(w_{0}(s))      = %s \n", shortcrsf8.String()[0:18])

	// E(t(s)) =

	var shortcrsf9 *bn256.G1
	if shortcrsf9 = new(bn256.G1).ScalarMult(pk, new(big.Int).Sub(s, r)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(t(s))          = %s \n \n", shortcrsf9.String()[0:18])

	log.Printf(" = = prover - P = = = = = = = = = = = = = = = \n \n")

	// v_{mid}(s) = Sigma_{k = 2}^{2} a_{k} * v_{k}(s) = a_{2} v_{2}(s)

	numerator = partialTerm(s, r, r1, s1, s2)
	denominator = partialTerm(r2, r, r1, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	var prover1 *bn256.G1
	if prover1 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(big.NewInt(2), term1)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(v_{mid}(s))     = %s \n", prover1.String()[0:18])

	// w(s) = Sigma_{k = 1}^{2} a_{k} * v_{k}(s) = a_{1} w_{1}(s) + a_{2} w_{2}(s)

	numerator = partialTerm(s, r, r1, r2, s2)
	denominator = partialTerm(s1, r, r1, r2, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	numerator = partialTerm(s, r1, r2, s1, s2)
	denominator = partialTerm(r, r1, r2, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term2 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	term3 = new(big.Int).Add(term1, term2)
	term4 = new(big.Int).Mul(big.NewInt(2), term3) // 2 * w_{1}(s)

	numerator = partialTerm(s, r, r1, r2, s1)
	denominator = partialTerm(s2, r, r1, r2, s1)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term5 = new(big.Int).Mul(big.NewInt(2), new(big.Int).Mul(numerator, inverse))

	term6 = new(big.Int).Mul(big.NewInt(6), term5) // 6 * w_{2}(s)

	var crsf30 *bn256.G1
	if crsf30 = new(bn256.G1).ScalarMult(pk, new(big.Int).Add(term4, term6)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(w(s))         = %s \n", crsf30.String()[0:18])

	// y(s) = Sigma_{k = 1}^{2} a_{k} * y_{k}(s) = a_{1} y_{1}(s) + a_{2} y_{2}(s)

	numerator = partialTerm(s, r, r2, s1, s2)
	denominator = partialTerm(r1, r, r2, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	numerator = partialTerm(s, r, r1, r2, s2)
	denominator = partialTerm(s1, r, r1, r2, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term2 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	term3 = new(big.Int).Mul(big.NewInt(2), new(big.Int).Add(term1, term2)) // 2 * y_{1}(s)

	numerator = partialTerm(s, r, r1, s1, s2)
	denominator = partialTerm(r2, r, r1, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term4 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	numerator = partialTerm(s, r, r1, r2, s1)
	denominator = partialTerm(s2, r, r1, r2, s1)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term5 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	numerator = partialTerm(s, r1, r2, s1, s2)
	denominator = partialTerm(r, r1, r2, s1, s2)

	inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	term7 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	term6 = new(big.Int).Add(term4, term5)
	term8 = new(big.Int).Mul(big.NewInt(6), new(big.Int).Add(term6, term7)) // 6 * y_{2}(s)

	var crsf31 *bn256.G1
	if crsf31 = new(bn256.G1).ScalarMult(pk, new(big.Int).Add(term3, term8)); err != nil {
		log.Printf("parameter generation %v", err)
	}

	// use multiplicative inverse check if equal :(
	log.Printf(" E(y_{2}(s))      = %s \n", crsf31.String()[0:18])

	return true
}

// LinearQAP generates the Quadratic Arithmetic Program to validate arithmetic circuits in Zero Knowledge
func LinearR1CS() bool {

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
