package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var mapping = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func isUnit(num int) bool {
	units := []int{1, 10, 100, 1000}
	for _, unit := range units {
		if num == unit {
			return true
		}
	}
	return false
}

func isFive(num int) bool {
	fives := []int{5, 50, 500}
	for _, five := range fives {
		if num == five {
			return true
		}
	}
	return false
}

func isValidSubtraction(number, subtract int) bool {
	if number <= (subtract * 10) {
		return true
	}
	return false
}

func matchingNumerals(numerals []int) bool {
	unique := make(map[int]bool)
	for _, val := range numerals {
		unique[val] = true
	}
	if len(unique) == 1 {
		return true
	}
	return false
}

type fragment struct {
	value       int
	subtraction int
}

func romanToArabic(s string) (int, error) {
	var numValues []int
	for _, val := range s {
		_, ok := mapping[string(val)]
		if !ok {
			return 0, errors.New("invalid character in provided numerals")
		}
		numValues = append(numValues, mapping[string(val)])
	}

	// check for illegaly repeating characters
	for i, val := range numValues {
		if isFive(val) && i < len(numValues)-1 {
			if matchingNumerals(numValues[i : i+2]) {
				return 0, errors.New("more than one sequential fives used")
			}
		}
		if isUnit(val) && i < len(numValues)-3 {
			if matchingNumerals(numValues[i : i+4]) {
				return 0, errors.New("more than three sequential units used")
			}
		}
	}

	var fragments []fragment

	for i := 0; i < len(numValues); i++ {
		var fragment fragment

		if len(numValues) == 1 {
			fragment.value = numValues[i]
		}
		if i < len(numValues)-1 {
			if numValues[i] < numValues[i+1] {
				fragment.value = numValues[i+1]
				fragment.subtraction = numValues[i]
				if !isValidSubtraction(fragment.value, fragment.subtraction) {
					return 0, errors.New("invalid subtraction in sequence")
				}
				i++
			} else {
				fragment.value = numValues[i]
			}
		}
		if i == len(numValues)-1 {
			fragment.value = numValues[i]
		}

		fragments = append(fragments, fragment)
		if len(fragments) > 1 {
			if fragments[len(fragments)-1].value > fragments[len(fragments)-2].value {
				return 0, errors.New("sequence is non-decreasing")
			}
		}
	}

	var arabic int
	for _, val := range fragments {
		arabic += (val.value - val.subtraction)
	}

	return arabic, nil
}

func main() {
	if len(os.Args) != 2 {
		err := errors.New("exactly one argument is required")
		log.Fatal(err)
	}
	numeralSequence := os.Args[1]
	arabic, err := romanToArabic(numeralSequence)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(arabic)
}
