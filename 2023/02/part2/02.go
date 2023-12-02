package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func getColorCubes(color string, set *string) int {
	regex := "\\d+ " + color
	re := regexp.MustCompile(regex)
	c := re.FindString(*set)
	c = strings.TrimRight(c, " "+color)
	count, _ := strconv.Atoi(c)
	return count
}

func main() {
	file, err := os.Open("../input")
	check(err)
	defer file.Close()

	const redCubes = 12
	const greenCubes = 13
	const blueCubes = 14

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		id, gameResult, _ := strings.Cut(scanner.Text(), ":")
		id = strings.TrimLeft(id, "Game ")
		sets := strings.Split(gameResult, ";")

		maxR, maxG, maxB := 0, 0, 0

		for _, set := range sets {
			r := getColorCubes("red", &set)
			g := getColorCubes("green", &set)
			b := getColorCubes("blue", &set)

			if r > maxR {
				maxR = r
			}
			if g > maxG {
				maxG = g
			}
			if b > maxB {
				maxB = b
			}
		}
		total += maxR * maxG * maxB
	}
	fmt.Println(total)
}
