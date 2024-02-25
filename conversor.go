package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	binaryInput := os.Args[1]

	binaryParts := strings.Split(binaryInput, ",")
	afterComma := binaryParts[0]

	var decimalAfterCommaList []int
	var decimalBeforeCommaList []int
	var decimalIntegerPart int
	var decimalFloatPart int

	for index := 0; index < len(afterComma); index++ {
		integer := int(afterComma[index] - '0')
		potency := float64(len(afterComma) - 1 - index)
		decimal_based_integer := integer * int(math.Pow(2, potency))
		decimalAfterCommaList = append(decimalAfterCommaList, decimal_based_integer)
	}

	for index := 0; index < len(decimalAfterCommaList); index++ {
		decimalIntegerPart += decimalAfterCommaList[index]
	}

	if len(binaryParts) > 1 {
		beforeComma := binaryParts[1]

		for index := 0; index < len(beforeComma); index++ {
			integer := int(beforeComma[index] - '0')
			potency := float64(len(beforeComma) - 1 - index)
			decimalBeforeComma := integer * int(math.Pow(2, potency))
			decimalBeforeCommaList = append(decimalBeforeCommaList, decimalBeforeComma)
		}

		for index := 0; index < len(decimalBeforeCommaList); index++ {
			decimalFloatPart += decimalBeforeCommaList[index]
		}
	}

	fmt.Println(float64(decimalIntegerPart) + (float64(decimalFloatPart) / 10))
}
