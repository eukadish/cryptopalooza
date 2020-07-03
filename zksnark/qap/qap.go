package qap

import (
	"math/big"
)

// BasisPolynomial generates a Lagrange basis polynomial modulo the order of the field.
func BasisPolynomial(order *big.Int, j int, xCoords ...*big.Int) func(*big.Int) *big.Int {

	// TODO: Better error handling for index out of range, dividing by zero, etc.

	var selected = xCoords[j]
	var denominator = big.NewInt(1)

	xCoords = append(xCoords[:j], xCoords[j+1:]...)

	var xCoord *big.Int
	for _, xCoord = range xCoords {
		denominator = new(big.Int).Mul(denominator, new(big.Int).Sub(selected, xCoord))
	}

	// TODO: Include the order of the field as a function parameter.
	return func(x *big.Int) *big.Int {

		var numerator = big.NewInt(1)

		var xCoord *big.Int
		for _, xCoord = range xCoords {
			numerator = new(big.Int).Mul(numerator, new(big.Int).Sub(x, xCoord))
		}

		// fmt.Printf(" = num %v \n", numerator)
		// fmt.Printf(" = denom %v \n", denominator)

		// return new(big.Int).Div(numerator, denominator)
		// return new(big.Int).Mul(numerator, new(big.Int).ModInverse(denominator, big.NewInt(7)))

		return new(big.Int).Mul(numerator, new(big.Int).ModInverse(denominator, order))
	}
}

// Interpolate loops through the basis polynomials and y-coordinates for evaluating at a point.
func Interpolate(
	x *big.Int, yCoords []int64, basis ...func(*big.Int) *big.Int,
) *big.Int {

	var accumulator = big.NewInt(0)

	var index int
	var base func(*big.Int) *big.Int

	for index, base = range basis {
		accumulator = new(big.Int).Add(accumulator, new(big.Int).Mul(big.NewInt(yCoords[index]), base(x)))
	}

	return accumulator
}
