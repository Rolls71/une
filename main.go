package main

import (
	"fmt"

	"github.com/Rolls71/une/cosc"
)

func crcTest() {
	generatorPolynomial := cosc.BinaryString("1101")
	data := cosc.BinaryString("1001001110")
	message := data + generatorPolynomial.TrailingZeroes()
	fmt.Printf("message:              %s\n", message)
	fmt.Printf("generator polynomial: %s\n", generatorPolynomial)
	fmt.Printf("xor polynomial:       %s\n", message.XOR(generatorPolynomial))
	fmt.Printf("left shifted:         %s\n", message.ShiftLeft())
	fmt.Printf("left shifted twice:   %s\n", message.ShiftLeft().ShiftLeft())
	fmt.Printf("left shifted thrice:  %s\n", message.ShiftLeft().ShiftLeft().ShiftLeft())
	remainder := cosc.CyclicRedundancyCheck(generatorPolynomial, message)
	fmt.Printf("calculate remainder:  %s\n", remainder)
	fmt.Printf("calculate message:    %s\n", cosc.CyclicRedundancyCheck(generatorPolynomial, data+remainder))
}

func main() {
	crcTest()
}
