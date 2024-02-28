package main

import (
	"fmt"
	"math"
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

	beforeComma := binaryParts[0]

	var decimalBeforeCommaList []float64
	var decimalAfterCommaList []float64
	var decimalIntegerPart float64
	var decimalFloatPart float64

	decimalBeforeCommaList = binaryToDecimalList(beforeComma)

	decimalIntegerPart = sum_list_elements(decimalBeforeCommaList)

	if len(binaryParts) > 1 {
		afterComma := binaryParts[1]

		decimalAfterCommaList = afterCommaBinaryToDecimalList(afterComma)

		decimalFloatPart = sum_list_elements(decimalAfterCommaList)
	}

	fmt.Println(float64(decimalIntegerPart) + (float64(decimalFloatPart)))
}

func afterCommaBinaryToDecimalList(binary string) []float64 {
	var decimalValues []float64
	var decimalValue float64

	for index := 0; index < len(binary); index++ {
		digit := float64(binary[index] - '0')
		exponentiation := float64(index+1) * -1

		decimalValue = digit * math.Pow(2, exponentiation)


		decimalValues = append(decimalValues, decimalValue)
	}

	return decimalValues
}

func binaryToDecimalList(binary string) []float64 {
	var decimalValues []float64

	for index := 0; index < len(binary); index++ {
		digit := float64(binary[index] - '0')
		decimalValue := digit * math.Pow(2, float64(len(binary) - 1 - index))
		decimalValues = append(decimalValues, decimalValue)
	}

	return decimalValues
}

func sum_list_elements(array []float64) float64 {
	var decimalFloatPart float64

	for index := 0; index < len(array); index++ {
		decimalFloatPart += array[index]
	}

	return decimalFloatPart
}
