package main

import (
	"bufio"
	"fmt"
	"os"
)

var labyrinth [][]string
var width, height int
var start [2]int
var path [][2]int
var results [][]string
var enclosed int

func load() {
	file, _ := os.Open("../input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		var line []string
		for j, char := range scanner.Text() {
			line = append(line, string(char))
			if string(char) == "S" {
				start = [...]int{i, j}
			}
			width = j + 1
		}
		labyrinth = append(labyrinth, line)
		i++
	}
	height = i
}

func step(i int, j int) {
	last_pos := path[len(path)-1]
	var direction string

	if j > last_pos[1] {
		direction = "E"
	} else if j < last_pos[1] {
		direction = "W"
	} else if i > last_pos[0] {
		direction = "S"
	} else if i < last_pos[0] {
		direction = "N"
	}

	path = append(path, [2]int{i, j})
	pipe := labyrinth[i][j]

	switch pipe {
	case "S":
		return
	case "|":
		if direction == "N" {
			step(i-1, j)
		} else if direction == "S" {
			step(i+1, j)
		}
	case "-":
		if direction == "W" {
			step(i, j-1)
		} else if direction == "E" {
			step(i, j+1)
		}
	case "L":
		if direction == "S" { //turn right
			step(i, j+1)
		} else if direction == "W" { //go up
			step(i-1, j)
		}
	case "J":
		if direction == "S" { //turn left
			step(i, j-1)
		} else if direction == "E" { //go up
			step(i-1, j)
		}
	case "7":
		if direction == "N" { //turn left
			step(i, j-1)
		} else if direction == "E" { //go down
			step(i+1, j)
		}
	case "F":
		if direction == "N" { //turn right
			step(i, j+1)
		} else if direction == "W" { //go down
			step(i+1, j)
		}
	}
}

func run() {
	i := start[0]
	j := start[1]
	path = append(path, start)

	if j < width { //right
		pipe := labyrinth[i][j+1]
		if pipe == "-" || pipe == "7" || pipe == "J" {
			step(i, j+1)
			return
		}
	}
	if i < height { //down
		pipe := labyrinth[i+1][j]
		if pipe == "|" || pipe == "L" || pipe == "J" {
			step(i+1, j)
			return
		}
	}
	if j > 0 { //left
		pipe := labyrinth[i][j-1]
		if pipe == "-" || pipe == "L" || pipe == "F" {
			step(i, j-1)
			return
		}
	}
	if i > 0 { //up
		pipe := labyrinth[i-1][j]
		if pipe == "|" || pipe == "7" || pipe == "F" {
			step(i-1, j)
			return
		}
	}
}

func visualize() {
	results = [][]string{}
	for i := 0; i < height; i++ {
		line := []string{}
		for j := 0; j < width; j++ {
			in_path := false
			for _, v := range path {
				if v[0] == i && v[1] == j {
					in_path = true
				}
			}
			if in_path {
				line = append(line, labyrinth[i][j])
			} else {
				line = append(line, ".")
			}
		}
		results = append(results, line)
	}

	enclosed = 0
	enter := false
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if results[i][j] == "|" || results[i][j] == "L" || results[i][j] == "J" {
				enter = !enter
			}
			if enter && results[i][j] == "." {
				results[i][j] = "@"
				enclosed++
			}
			if i == 0 || i == height-1 || j == 0 || j == width-1 {
				if results[i][j] == "@" {
					results[i][j] = "."
				}
			}
		}
	}
	for i := 0; i < height; i++ {
		fmt.Println(results[i])
	}
}

func main() {
	load()
	run()
	visualize()
	fmt.Println(enclosed)
}
