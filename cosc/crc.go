// Cyclic Redundancy Check (CRC)

package cosc

// BinaryString is a string consisting of exclusively 1s and 0s
type BinaryString string

// XOR performs XOR between the two given string and returns the result.
func (b BinaryString) XOR(divisor BinaryString) BinaryString {
	result := ""
	runedivisor := []rune(divisor)
	for index, bit := range b {
		if len(runedivisor) <= index {
			result += string(bit)
			continue
		}

		if bit == runedivisor[index] {
			result += "0"
		} else {
			result += "1"
		}
	}

	return BinaryString(result)
}

// ShiftLeft returns the starting string with the first character at the end.
func (b BinaryString) ShiftLeft() BinaryString {
	start := b[:1]
	end := b[1:]
	return end + start
}

// TrailingZeroes returns a string of repeating 0s, with length equal to
// one less than that of the given string's.
// This is useful for appending zeroes to data before passing through
// CyclicRedundancyCheck.
func (b BinaryString) TrailingZeroes() BinaryString {
	zeroes := ""
	for i := len(b) - 1; i > 0; i-- {
		zeroes += "0"
	}
	return BinaryString(zeroes)
}

// CyclicRedundancyCheck performs a Cyclic Redundancy Check on the given message
// using the given generator polynomial. The given message must have enough
// n-1 traling zeroes, where n i equal to the length of the generator polynomial
// string.
func CyclicRedundancyCheck(
	generatorPolynomial BinaryString,
	message BinaryString,
) BinaryString {
	remainder := message
	for i := len(message) - len(generatorPolynomial) + 1; i > 0; i-- {
		if remainder[:1] == "1" {
			remainder = remainder.XOR(generatorPolynomial)
		}
		remainder = remainder.ShiftLeft()
	}
	return remainder[:len(generatorPolynomial)-1]
}

// CalculatorCRC allows the user to input a generator polynomial and a binary
// message, and displays the remainder.
// Given values must consist entirely of '1's and '0's
func CalculatorCRC() bool
