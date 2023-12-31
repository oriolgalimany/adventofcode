package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Node struct {
	left  string
	right string
	id    string
}

var historyset [][]int
var total = 0

func load() {
	// file, _ := os.Open("input")
	file, _ := os.Open("test")
	defer file.Close()

	re := regexp.MustCompile("(-)*\\d+")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := re.FindAllString(scanner.Text(), -1)
		var history []int
		for _, v := range values {
			n, _ := strconv.Atoi(v)
			history = append(history, n)
		}
		historyset = append(historyset, history)
	}
}
func allZeroes(array []int) bool {
	for _, a := range array {
		if a != 0 {
			return false
		}
	}
	return true
}

func calc(history []int, process *[][]int) {
	var array []int
	for i := 0; i < len(history); i++ {
		if i < len(history)-1 {
			array = append(array, history[i+1]-history[i])
		}
	}
	if allZeroes(array) {
		*process = append(*process, array)
		return
	}
	*process = append(*process, array)
	calc(array, process)
}

func predict(process_ptr *[][]int) {
	prediction := 0
	process := *process_ptr
	for j := len(process) - 1; j >= 0; j-- {
		if j == len(process)-1 {
			prediction = process[j][len(process[j])-1]
			process[j] = append(process[j], prediction)
		} else {
			prediction = process[j][len(process[j])-1] + process[j+1][len(process[j+1])-1]
			process[j] = append(process[j], prediction)
		}
	}
	total += prediction
}

func main() {
	load()
	for _, history := range historyset {
		var process [][]int
		process = append(process, history)
		calc(history, &process)
		predict(&process)
	}
	fmt.Println(total)
}
