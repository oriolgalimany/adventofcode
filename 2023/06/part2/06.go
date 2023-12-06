package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var time int
var distance int

func toInt(str []string) int {
	s := ""
	for _, i := range str {
		s += i
	}
	n, _ := strconv.Atoi(s)
	return n
}

func load() {
	file, _ := os.Open("../input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("\\d+")
	scanner.Scan()
	time = toInt(re.FindAllString(scanner.Text(), -1))
	scanner.Scan()
	distance = toInt(re.FindAllString(scanner.Text(), -1))
}

func main() {
	load()

	total := 1
	records := 0

	for t := 1; t < time; t++ {
		speed := 1 * t
		time_left := time - t
		d := speed * time_left
		if d > distance {
			records++
		}
	}

	total *= records

	fmt.Println(total)
}
