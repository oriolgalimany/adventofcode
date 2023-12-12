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
var width, height int
var free_rows map[int]struct{}
var free_cols map[int]struct{}

func load() {
	file, _ := os.Open("input")
	defer file.Close()

	free_rows = make(map[int]struct{})
	free_cols = make(map[int]struct{})
	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line []string
		var free_row = true
		for j, char := range scanner.Text() {
			line = append(line, string(char))
			if string(char) == "#" {
				free_row = false
			}
			width = j + 1

		}
		if free_row {
			free_rows[i] = struct{}{}
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
			free_cols[i] = struct{}{}
		}
	}
}

func expand() {
	for i := 0; i < height; i++ {
		var line []string
		if _, exist := free_rows[i]; exist {
			for n := 0; n < (width + len(free_cols)); n++ {
				line = append(line, string("."))
			}
			new_universe = append(new_universe, line)
			new_universe = append(new_universe, line)
			continue
		}

		for j := 0; j < width; j++ {
			if _, exist := free_cols[j]; exist {
				line = append(line, string("."))
				line = append(line, string("."))
			} else {
				line = append(line, universe[i][j])
			}
		}
		new_universe = append(new_universe, line)
	}
}

func setGalaxies() {
	for i := 0; i < (height + len(free_rows)); i++ {
		for j := 0; j < (width + len(free_cols)); j++ {
			if new_universe[i][j] == "#" {
				galaxies = append(galaxies, [...]int{i, j})
			}
		}
	}
}

func pairs() {
	total := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i; j < len(galaxies); j++ {
			if i != j {
				distance := math.Abs(path(galaxies[i], galaxies[j]))
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
	setGalaxies()
	// for _, u := range new_universe {
	// 	fmt.Println(u)
	// }
	pairs()
}
