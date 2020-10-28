package main

import (
	"fmt"

	"github.com/cloudflare/bn256"
	"github.com/eugenekadish/cryptopalooza/sm"
	"github.com/eugenekadish/cryptopalooza/zksm"
	"github.com/eugenekadish/cryptopalooza/zksnark/qap"
)

func main() {

	var order = bn256.Order

	// var order = bn256.Order.Set(big.NewInt(11))
	// var order = bn256.Order.Set(big.NewInt(23))
	// var order = bn256.Order.Set(big.NewInt(997))

	fmt.Println()

	fmt.Printf(" = Arithmetic Circuit = = \n")

	fmt.Println()

	fmt.Printf("  - Example 1 QAP         %t \n", qap.E1QAP(order))
	fmt.Printf("  - Example 1 Strong QAP  %t \n", qap.E1SQAP(order))

	fmt.Printf("  - Example 2 QAP         %t \n", qap.E2QAP(order))
	fmt.Printf("  - Example 2 R1CS        %t \n", qap.E2R1CS(order))

	fmt.Printf("  - Example 3 QAP         %t \n", qap.E3QAP(order))
	fmt.Printf("  - Example 3 R1CS        %t \n", qap.E3R1CS(order))

	fmt.Println()

	fmt.Printf(" = Set Membership = = \n")

	fmt.Println()

	fmt.Printf("  - Zero Knowledge                                  %t \n", zksm.E1SM(order))
	fmt.Printf("  - Bilinear-map Accumulator                        %t \n", sm.E2ACCUM(order))

	fmt.Println()

	fmt.Printf("  - RSA Accumulator                                 %t \n", sm.E1ACCUM(order))
	fmt.Printf("  - RSA Accumulator (hash to prime)                 %t \n", sm.E3ACCUM(order))
	fmt.Printf("  - RSA Accumulator (two-universal hash functions)  %t \n", sm.E4ACCUM(order))

	fmt.Println()
}
