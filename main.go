package main

import (
	"fmt"
	"os"
	"strconv"
)

// CheckDigitAlgorithms holds the functions for different check digit algorithms
var CheckDigitAlgorithms = map[string]func(string) int{
	"UPC":      calculateUPCCheckDigit,
	"EAN-13":   calculateEAN13CheckDigit,
	"EAN-8":    calculateEAN8CheckDigit,
	"ISBN-10":  calculateISBN10CheckDigit,
	"ISBN-13":  calculateISBN13CheckDigit,
	"ISSN":     calculateISSNCheckDigit,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: barcode-checkdigit <barcode>")
		return
	}

	barcode := os.Args[1]
	found := false

	for name, algorithm := range CheckDigitAlgorithms {
		if isValidCheckDigit(barcode, algorithm) {
			fmt.Printf("The barcode %s uses the %s check digit algorithm.\n", barcode, name)
			found = true
		}
	}

	if !found {
		fmt.Println("No known check digit algorithm found for the barcode.")
	}
}

func isValidCheckDigit(barcode string, algorithm func(string) int) bool {
	if len(barcode) < 1 {
		return false
	}

	checkDigit, err := strconv.Atoi(string(barcode[len(barcode)-1]))
	if err != nil {
		return false
	}

	calculatedCheckDigit := algorithm(barcode[:len(barcode)-1])
	return checkDigit == calculatedCheckDigit
}

func calculateUPCCheckDigit(barcode string) int {
	sum := 0
	for i, r := range barcode {
		digit, _ := strconv.Atoi(string(r))
		if i%2 == 0 {
			sum += digit
		} else {
			sum += digit * 3
		}
	}
	return (10 - (sum % 10)) % 10
}

func calculateEAN13CheckDigit(barcode string) int {
	sum := 0
	for i, r := range barcode {
		digit, _ := strconv.Atoi(string(r))
		if i%2 == 0 {
			sum += digit
		} else {
			sum += digit * 3
		}
	}
	return (10 - (sum % 10)) % 10
}

func calculateEAN8CheckDigit(barcode string) int {
	sum := 0
	for i, r := range barcode {
		digit, _ := strconv.Atoi(string(r))
		if i%2 == 0 {
			sum += digit
		} else {
			sum += digit * 3
		}
	}
	return (10 - (sum % 10)) % 10
}

func calculateISBN10CheckDigit(barcode string) int {
	sum := 0
	for i, r := range barcode {
		digit, _ := strconv.Atoi(string(r))
		sum += digit * (i + 1)
	}
	checkDigit := sum % 11
	if checkDigit == 10 {
		return 'X' // ISBN-10 uses 'X' for check digit 10
	}
	return checkDigit
}

func calculateISBN13CheckDigit(barcode string) int {
	sum := 0
	for i, r := range barcode {
		digit, _ := strconv.Atoi(string(r))
		if i%2 == 0 {
			sum += digit
		} else {
			sum += digit * 3
		}
	}
	return (10 - (sum % 10)) % 10
}

func calculateISSNCheckDigit(barcode string) int {
	sum := 0
	for i, r := range barcode {
		digit, _ := strconv.Atoi(string(r))
		sum += digit * (8 - i)
	}
	checkDigit := sum % 11
	if checkDigit == 10 {
		return 'X' // ISSN uses 'X' for check digit 10
	}
	return checkDigit
}
