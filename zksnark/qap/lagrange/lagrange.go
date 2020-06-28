package lagrange

import (
	"fmt"
	"math/big"

	"github.com/cloudflare/bn256"
)

// BasisPolynomial generates a Lagrange basis polynomial modulo the order of the field.
func BasisPolynomial(j int, xCoords ...*big.Int) func(*big.Int) *big.Int {

	// TODO: Return an error if the index is out of range

	var selected = xCoords[j]
	var denominator = big.NewInt(1)

	xCoords = append(xCoords[:j], xCoords[j+1:]...)

	var xCoord *big.Int
	for _, xCoord = range xCoords {
		denominator = new(big.Int).Mul(denominator, new(big.Int).Sub(selected, xCoord))
	}

	return func(x *big.Int) *big.Int {

		var numerator = big.NewInt(1)

		var xCoord *big.Int
		for _, xCoord = range xCoords {
			numerator = new(big.Int).Mul(numerator, new(big.Int).Sub(x, xCoord))
		}

		// return new(big.Int).Div(numerator, denominator)
		return new(big.Int).Mul(numerator, new(big.Int).ModInverse(denominator, bn256.Order))
	}
}

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
) (accumulator *big.Int) {

	fmt.Printf("basis %d", len(basis))

	var index int
	var base func(*big.Int) *big.Int

	for index, base = range basis {
		accumulator = new(big.Int).Add(accumulator, new(big.Int).Mul(big.NewInt(yCoords[index]), base(x)))
	}

	return
}
