package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var universe [][]string
var new_universe [][]string
var galaxies [][2]int
var new_galaxies [][2]int
var width, height int
var free_rows []int
var free_cols []int

const K = 1000000

func load() {
	file, _ := os.Open("../input")
	defer file.Close()

	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line []string
		var free_row = true
		for j, char := range scanner.Text() {
			line = append(line, string(char))
			if string(char) == "#" {
				galaxies = append(galaxies, [...]int{i, j})
				new_galaxies = append(new_galaxies, [...]int{i, j})
				free_row = false
			}
			width = j + 1

		}
		if free_row {
			free_rows = append(free_rows, i)
		}
		universe = append(universe, line)
		i++
	}
	height = i

	for i := 0; i < width; i++ {
		var free_col = true
		for j := 0; j < height; j++ {
			if universe[j][i] == "#" {
				free_col = false
			}
		}
		if free_col {
			free_cols = append(free_cols, i)
		}
	}
}

func expand() {
	for _, col := range free_cols {
		for i := 0; i < len(galaxies); i++ {
			if galaxies[i][1] > col {
				new_galaxies[i][1] = new_galaxies[i][1] + K - 1
			}
		}
	}
	for _, row := range free_rows {
		for i := 0; i < len(galaxies); i++ {
			if galaxies[i][0] > row {
				new_galaxies[i][0] = new_galaxies[i][0] + K - 1
			}
		}
	}
}

func pairs() {
	total := 0
	for i := 0; i < len(new_galaxies); i++ {
		for j := i; j < len(new_galaxies); j++ {
			if i != j {
				distance := math.Abs(path(new_galaxies[i], new_galaxies[j]))
				total += int(distance)
			}
		}
	}
	fmt.Println(total)
}

func path(init [2]int, end [2]int) float64 {
	x := end[0] - init[0]
	y := end[1] - init[1]
	result := math.Abs(float64(x)) + math.Abs(float64(y))
	return result
}

func main() {
	load()
	expand()
	pairs()
}
