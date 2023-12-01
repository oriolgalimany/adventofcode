package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("01input")
	check(err)

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("[0-9]")
	total := 0

	for scanner.Scan() {
		digits := re.FindAllString(scanner.Text(), -1)
		value, err := strconv.Atoi((digits[0] + digits[len(digits)-1]))
		check(err)
		total += value
	}

	fmt.Println(total)
	file.Close()
}
