package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isBinary(number string) bool {
	for _, ch := range number {
		if !unicode.IsDigit(ch) || (ch != '0' && ch != '1') {
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)

	// Trim any leading or trailing whitespace
	input = strings.TrimSpace(input)

	if isBinary(input) {
		fmt.Println(input, "is a binary number.")
	} else {
		fmt.Println(input, "is not a binary number.")
	}
}
