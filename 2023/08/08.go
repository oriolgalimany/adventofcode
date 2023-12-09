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
var start = "AAA"
var goal = "ZZZ"
var steps []string

func load() {
	file, _ := os.Open("input")
	defer file.Close()

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
	}
}

func run() {
	current_node := start
	for i := 0; i < len(instructions); i++ {
		step := string(instructions[i])
		if step == "L" {
			current_node = network[current_node].left
		} else if step == "R" {
			current_node = network[current_node].right
		}
		steps = append(steps, step)
		if current_node == goal {
			break
		}
		if i == len(instructions)-1 {
			i = -1
		}
	}
}

func main() {
	network = make(map[string]Node)
	load()
	run()
	fmt.Println(len(steps))
}
