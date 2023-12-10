package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Node struct {
	left  string
	right string
	id    string
}

var network map[string]Node
var instructions string

var steps int64
var n_steps []int64

var current_nodes []string
var goals map[string]struct{}

func load() {
	file, _ := os.Open("../input")
	defer file.Close()

	goals = make(map[string]struct{})

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions = scanner.Text()
	re := regexp.MustCompile("\\w+")
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		values := re.FindAllString(scanner.Text(), -1)
		var node Node
		node.id = values[0]
		node.left = values[1]
		node.right = values[2]
		network[node.id] = node

		if string(values[0][2]) == "A" {
			current_nodes = append(current_nodes, values[0])
		}
		if string(values[0][2]) == "Z" {
			goals[values[0]] = struct{}{}
		}
	}
}

func LCM(array []int64) int64 {
	n := int64(1)
	for i := 0; i < len(array); i++ {
		n = LCM2(array[i], n)
	}
	return n
}

func LCM2(a int64, b int64) int64 {
	max := int64(0)
	min := int64(0)
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	for i := max; ; i += max {
		if i%min == 0 {
			return i
		}
	}
}

func run() {
	for i := 0; i < len(current_nodes); i++ {
		for j := 0; j < len(instructions); j++ {
			step := string(instructions[j])
			steps++
			if step == "L" {
				current_nodes[i] = network[current_nodes[i]].left
			} else if step == "R" {
				current_nodes[i] = network[current_nodes[i]].right
			}
			if _, exist := goals[current_nodes[i]]; exist {
				n_steps = append(n_steps, steps)
				steps = 0
				break
			}
			if j == len(instructions)-1 {
				j = -1
			}
		}
	}
}

func main() {
	network = make(map[string]Node)
	load()
	run()
	fmt.Println(n_steps)
	fmt.Println(LCM(n_steps))
}
