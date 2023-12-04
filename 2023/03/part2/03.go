package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var matrix [][]string
var gears map[string][]*Digit //ij as key, #adjacents as value

type Digit struct {
	number string
	valid  bool
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func adjacent(i, j int, digit *Digit) bool {
	if _, err := strconv.Atoi(matrix[i][j]); err != nil && matrix[i][j] == "*" {
		id := strconv.Itoa(i) + strconv.Itoa(j)

		for _, value := range gears[id] {
			if value == digit {
				return true
			}
		}

		gears[id] = append(gears[id], digit)
		return true
	}
	return false
}

func valid(i, j, rows, columns int, digit *Digit) {
	// top i-1, j
	if i-1 >= 0 && adjacent(i-1, j, digit) {
		digit.valid = true
	}
	// bottom i+1, j
	if i+1 < rows && adjacent(i+1, j, digit) {
		digit.valid = true
	}
	// left i,j-1
	if j-1 >= 0 && adjacent(i, j-1, digit) {
		digit.valid = true
	}
	// right i,j+1
	if j+1 < columns && adjacent(i, j+1, digit) {
		digit.valid = true
	}
	// top left i-1, j-1
	if i-1 >= 0 && j-1 >= 0 && adjacent(i-1, j-1, digit) {
		digit.valid = true
	}
	// top right i-1, j+1
	if i-1 >= 0 && j+1 < columns && adjacent(i-1, j+1, digit) {
		digit.valid = true
	}
	// bottom left i+1, j-1
	if i+1 < rows && j-1 >= 0 && adjacent(i+1, j-1, digit) {
		digit.valid = true
	}
	// bottom right i+1, j+1
	if i+1 < rows && j+1 < columns && adjacent(i+1, j+1, digit) {
		digit.valid = true
	}
}

func main() {
	file, err := os.Open("../input")
	check(err)
	defer file.Close()

	rows := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matrix = append(matrix, strings.Split(scanner.Text(), ""))
		rows++
	}

	total := 0

	var digit *Digit
	digit = new(Digit)
	gears = make(map[string][]*Digit)

	for i := 0; i < rows; i++ {
		row := matrix[i]

		for j := 0; j < len(row); j++ {
			if _, err := strconv.Atoi(row[j]); err == nil {
				digit.number += row[j]
				valid(i, j, rows, len(row), digit)
			} else {
				if digit.number != "" {
					digit = new(Digit)
				}
			}
		}
		digit = new(Digit)
	}

	for _, v := range gears {
		if len(v) == 2 {
			v0, _ := strconv.Atoi(v[0].number)
			v1, _ := strconv.Atoi(v[1].number)
			total += v0 * v1
		}
	}
	fmt.Println(total)
}
