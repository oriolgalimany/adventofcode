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

func main() {
	file, err := os.Open("input")
	// file, err := os.Open("test")
	check(err)
	defer file.Close()

	const redCubes = 12
	const greenCubes = 13
	const blueCubes = 14

	reRed := regexp.MustCompile("\\d+ red")
	reGreen := regexp.MustCompile("\\d+ green")
	reBlue := regexp.MustCompile("\\d+ blue")

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		id, gameResult, _ := strings.Cut(scanner.Text(), ":")
		id = strings.TrimLeft(id, "Game ")
		i, _ := strconv.Atoi(id)

		fmt.Println("id: ", i)
		sets := strings.Split(gameResult, ";")

		fmt.Println(sets)
		reds, greens, blues := 0, 0, 0

		valid := true

		for _, set := range sets {
			red := reRed.FindString(set)
			red = strings.TrimRight(red, " red")
			r, _ := strconv.Atoi(red)
			if r > redCubes {
				valid = false
				break
			}

			green := reGreen.FindString(set)
			green = strings.TrimRight(green, " green")
			g, _ := strconv.Atoi(green)
			if g > greenCubes {
				valid = false
				break
			}

			blue := reBlue.FindString(set)
			blue = strings.TrimRight(blue, " blue")
			b, _ := strconv.Atoi(blue)
			if b > blueCubes {
				valid = false
				break
			}
		}

		fmt.Println("reds: ", reds, " greens: ", greens, " blues: ", blues)
		if valid {
			total += i
		}
	}
	fmt.Println(total)
}
