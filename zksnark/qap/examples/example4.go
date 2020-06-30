package examples

import (
	"crypto/rand"
	"fmt"
	"log"
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

// Example4 generates the Quadratic Arithmetic Program to validate arithmetic circuits in Zero Knowledge
func Example4() bool {

	var err error

	// var r1 *big.Int
	// if r1, _, err = bn256.RandomG1(rand.Reader); err != nil {
	// 	log.Printf("parameter generation %v", err)
	// }

	// var r2 *big.Int
	// if r2, _, err = bn256.RandomG1(rand.Reader); err != nil {
	// 	log.Printf("parameter generation %v", err)
	// }

	// var s1 *big.Int
	// if s1, _, err = bn256.RandomG1(rand.Reader); err != nil {
	// 	log.Printf("parameter generation %v", err)
	// }

	// var s2 *big.Int
	// if s2, _, err = bn256.RandomG1(rand.Reader); err != nil {
	// 	log.Printf("parameter generation %v", err)
	// }

	// var r *big.Int
	// if r, _, err = bn256.RandomG1(rand.Reader); err != nil {
	// 	log.Printf("parameter generation %v", err)
	// }

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

	// crs - = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

	var pk *bn256.G1
	if pk = new(bn256.G1).ScalarBaseMult(sk); err != nil {
		log.Printf("parameter generation %v", err)
	}

	var crs1 *bn256.G1
	if crs1 = new(bn256.G1).ScalarMult(pk, big.NewInt(1)); err != nil { // Encrypt s^{0}
		log.Printf("parameter generation %v", err)
	}

	var crs2 *bn256.G1
	if crs2 = new(bn256.G1).ScalarMult(pk, s); err != nil { // Encrypt s^{1}
		log.Printf("parameter generation %v", err)
	}

	var crs3 *bn256.G1
	if crs3 = new(bn256.G1).ScalarMult(pk, alpha); err != nil { // Encrypt alpha
		log.Printf("parameter generation %v", err)
	}

	var crs4 *bn256.G1
	if crs4 = new(bn256.G1).ScalarMult(pk, new(big.Int).Mul(alpha, s)); err != nil { // Encrypt alpha * s
		log.Printf("parameter generation %v", err)
	}

	log.Printf("crs1: %s \n", crs1.String()[0:18])
	log.Printf("crs2: %s \n", crs2.String()[0:18])
	log.Printf("crs3: %s \n", crs3.String()[0:18])
	log.Printf("crs4: %s \n", crs4.String()[0:18])

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

	fmt.Printf("betaV: %s \n", betaV.String()[0:12])
	fmt.Printf("betaW: %s \n", betaW.String()[0:12])
	fmt.Printf("betaY: %s \n", betaY.String()[0:12])

	fmt.Printf("gamma: %s \n", gamma.String()[0:12])

	// crs-f - = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

	// crs

	// I_{free} = {}

	// I_{labeled} = I_{10} U I_{11} = {1, 2}

	// I_{in} = {1}

	// I_{mid} = {2}

	// E(v_{2}(s)) =

	// var term1, term2, term3, term4 *big.Int
	// var numerator, inverse, denominator *big.Int

	// numerator = partialTerm(s, r, r1, s1, s2)
	// denominator = partialTerm(r2, r, r1, s1, s2)

	// inverse = new(big.Int).ModInverse(denominator, bn256.Order)
	// term1 = new(big.Int).Mul(big.NewInt(1), new(big.Int).Mul(numerator, inverse))

	// var crsf1 *bn256.G1
	// if crsf1 = new(bn256.G1).ScalarMult(pk, term1); err != nil {
	// 	log.Printf("parameter generation %v", err)
	// }

	// // use multiplicative inverse check if equal :(
	// log.Printf("E(v_{2}(s) = %s \n", crsf1.String()[0:18])

	return true
}
