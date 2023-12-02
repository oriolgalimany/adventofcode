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
	file, err := os.Open("input")
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
		i, _ := strconv.Atoi(id)

		sets := strings.Split(gameResult, ";")
		valid := true

		for _, set := range sets {
			if getColorCubes("red", &set) > redCubes {
				valid = false
				break
			}

			if getColorCubes("green", &set) > greenCubes {
				valid = false
				break
			}

			if getColorCubes("blue", &set) > blueCubes {
				valid = false
				break
			}
		}
		if valid {
			total += i
		}
	}
	fmt.Println(total)
}
