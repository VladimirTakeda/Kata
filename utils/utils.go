package utils

import (
	"bufio"
	"strconv"
	"strings"
)

var romanInput = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

var arabicInput = map[string]int{
	"1":  1,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
}

type SystemType int

const (
	Roman   SystemType = 0
	Arabic  SystemType = 1
	Unknown SystemType = 2
)

func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func checkSystemType(firstParam, secondParam string) (SystemType, int, int) {
	var firstValue int
	var secondValue int
	var firstFound bool
	var secondFound bool

	systemType := Unknown

	firstValue, firstFound = romanInput[firstParam]
	secondValue, secondFound = romanInput[secondParam]

	if firstFound && secondFound {
		systemType = Roman
	}

	if systemType == Unknown {
		firstValue, firstFound = arabicInput[firstParam]
		secondValue, secondFound = arabicInput[secondParam]

		if firstFound && secondFound {
			systemType = Arabic
		}
	}
	return systemType, firstValue, secondValue
}

func prepareParams(expression string) (string, string, string) {
	reader := bufio.NewReader(strings.NewReader(expression))
	first, err := reader.ReadString(' ')
	if err != nil {
		panic("Can't read first number in string")
	}
	operation, err := reader.ReadString(' ')
	if err != nil {
		panic("Can't read operation in string")
	}
	second, err := reader.ReadString('\n')
	if err != nil {
		panic("Can't read second number in string")
	}

	first = first[:len(first)-1]
	operation = operation[:len(operation)-1]
	second = second[:len(second)-1]

	return first, operation, second
}

func Calculate(expression string) string {
	firstParam, operation, secondParam := prepareParams(expression)
	systemType, firstValue, secondValue := checkSystemType(firstParam, secondParam)

	if systemType == Unknown {
		panic("params from different systems or out of limits")
	}

	var result int
	switch operation {
	case "+":
		result = firstValue + secondValue
	case "-":
		result = firstValue - secondValue
	case "*":
		result = firstValue * secondValue
	case "/":
		result = firstValue / secondValue
	default:
		panic("Wrong operation")
	}

	if result < 1 && systemType == Roman {
		panic("Roman result should be positive")
	}

	if systemType == Roman {
		return integerToRoman(result)
	}
	if systemType == Arabic {
		return strconv.Itoa(result)
	}
	return ""
}
