package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("ERROR: You should provide an argument")
		return
	}

	binaryInput := os.Args[1]
	var binaryParts []string

	containsDot := strings.Contains(binaryInput, ".")

	if containsDot == true {
		binaryParts = strings.Split(binaryInput, ".")
	} else {
		binaryParts = strings.Split(binaryInput, ",")
	}

	afterComma := binaryParts[0]

	var decimalAfterCommaList []int
	var decimalBeforeCommaList []int
	var decimalIntegerPart int
	var decimalFloatPart int

	decimalAfterCommaList = binaryToDecimalList(afterComma)

	decimalIntegerPart = sum_list_elements(decimalAfterCommaList)

	if len(binaryParts) > 1 {
		beforeComma := binaryParts[1]

		decimalBeforeCommaList = binaryToDecimalList(beforeComma)

		decimalFloatPart = sum_list_elements(decimalBeforeCommaList)
	}

	fmt.Println(float64(decimalIntegerPart) + (float64(decimalFloatPart) / 10))
}

func binaryToDecimalList(binary string) []int {
	var decimalValues []int

	for index := 0; index < len(binary); index++ {
		digit := int(binary[index] - '0')
		decimalValue := digit << (len(binary) - 1 - index)
		decimalValues = append(decimalValues, decimalValue)
	}

	return decimalValues
}

func sum_list_elements(array []int) int {
	var decimalFloatPart int

	for index := 0; index < len(array); index++ {
		decimalFloatPart += array[index]
	}

	return decimalFloatPart
}
