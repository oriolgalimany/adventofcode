package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numbersMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("../01input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {

		var digits []string

		var digitString = ""
		for _, ch := range scanner.Text() {

			if unicode.IsNumber((ch)) {
				digitString = ""
				digits = append(digits, string(ch))
				break
			}

			digitString += string(ch)
			for k, v := range numbersMap {
				if strings.Contains(digitString, k) {
					digits = append(digits, v)
					digitString = digitString[len(digitString)-1:]
				}
			}

		}

		first, err := strconv.Atoi(digits[0])
		check(err)
		last, err := strconv.Atoi(digits[len(digits)-1])
		check(err)

		value := first*10 + last
		total += value
	}

	fmt.Println(total)
}
