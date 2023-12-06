package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var times []int
var distances []int

func toInt(str []string) []int {
	var array []int
	for _, s := range str {
		i, _ := strconv.Atoi(s)
		array = append(array, i)
	}
	return array
}

func load() {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("\\d+")
	scanner.Scan()
	times = toInt(re.FindAllString(scanner.Text(), -1))
	scanner.Scan()
	distances = toInt(re.FindAllString(scanner.Text(), -1))
}

func main() {
	load()

	total := 1
	for i, time := range times {
		records := 0
		for t := 1; t < time; t++ {
			speed := 1 * t
			time_left := time - t
			distance := speed * time_left
			if distance > distances[i] {
				records++
			}
		}
		total *= records
	}
	fmt.Println(total)
}
