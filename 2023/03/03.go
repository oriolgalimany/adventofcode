package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func number(c string) (int, bool) {
	isNumber := false
	value, err := strconv.Atoi(c)
	if err == nil {
		isNumber = true
	}
	return value, isNumber
}

type Digit struct {
	number    string
	positions [][]int
	valid     bool
}

func specialChar(i, j int, matrix *[][]string) bool {
	if _, err := strconv.Atoi((*matrix)[i][j]); err != nil && (*matrix)[i][j] != "." {
		return true
	}
	return false
}

func valid(i, j, rows, columns int, matrix *[][]string) bool {
	if i-1 >= 0 && specialChar(i-1, j, matrix) || // top i-1, j
		i+1 < rows && specialChar(i+1, j, matrix) || // bottom i+1, j
		j-1 >= 0 && specialChar(i, j-1, matrix) || // left i,j-1
		j+1 < columns && specialChar(i, j+1, matrix) || // right i,j+1
		i-1 >= 0 && j-1 >= 0 && specialChar(i-1, j-1, matrix) || // top left i-1, j-1
		i-1 >= 0 && j+1 < columns && specialChar(i-1, j+1, matrix) || // top right i-1, j+1
		i+1 < rows && j-1 >= 0 && specialChar(i+1, j-1, matrix) || // bottom left i+1, j-1
		i+1 < rows && j+1 < columns && specialChar(i+1, j+1, matrix) { // bottom right i+1, j+1
		return true
	}
	return false
}

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	var matrix [][]string
	rows := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matrix = append(matrix, strings.Split(scanner.Text(), ""))
		rows++
	}

	total := 0

	var digit Digit

	for i := 0; i < rows; i++ {
		row := matrix[i]

		for j := 0; j < len(row); j++ {

			if _, err := strconv.Atoi(row[j]); err == nil {
				digit.number += row[j]
				digit.positions = append(digit.positions, []int{i, j})
				if !digit.valid {
					digit.valid = valid(i, j, rows, len(row), &matrix)
				}
			} else {
				if digit.number != "" {

					if digit.valid {
						n, _ := strconv.Atoi(digit.number)
						total += n
					}
					digit.number = ""
					digit.positions = nil
					digit.valid = false
				}
			}
		}

		if digit.valid {
			n, _ := strconv.Atoi(digit.number)
			total += n
		}
		digit.number = ""
		digit.positions = nil
		digit.valid = false
	}
	fmt.Println(total)
}
