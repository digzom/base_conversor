package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	binary_number := os.Args[1]

	binary_numbers := strings.Split(binary_number, ",")
	after_comma := binary_numbers[0]

	var decimal_based_after_comma []int
	var decimal_based_before_comma []int
	var decimal_integer int
	var decimal_float_part int

	for index := 0; index < len(after_comma); index++ {
		integer := int(after_comma[index] - '0')
		potency := float64(len(after_comma) - 1 - index)
		decimal_based_integer := integer * int(math.Pow(2, potency))
		decimal_based_after_comma = append(decimal_based_after_comma, decimal_based_integer)
	}

	for index := 0; index < len(decimal_based_after_comma); index++ {
		decimal_integer += decimal_based_after_comma[index]
	}

	if len(binary_numbers) > 1 {
		before_comma := binary_numbers[1]

		for index := 0; index < len(before_comma); index++ {
			integer := int(before_comma[index] - '0')
			potency := float64(len(before_comma) - 1 - index)
			decimal_before_comma := integer * int(math.Pow(2, potency))
			decimal_based_before_comma = append(decimal_based_before_comma, decimal_before_comma)
		}

		for index := 0; index < len(decimal_based_before_comma); index++ {
			decimal_float_part += decimal_based_before_comma[index]
		}
	}

	fmt.Println(float64(decimal_integer) + (float64(decimal_float_part) / 10))
}
