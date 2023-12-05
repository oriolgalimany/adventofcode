package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	total := 0.0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ":")
		card := strings.Split(s[1], "|")

		winning := make(map[string]struct{}) //only key map , struct{} is 0 bytes

		for _, w := range strings.Split(card[0], " ") {
			if w != "" {
				winning[w] = struct{}{}
			}
		}
		win := 0.0
		for _, n := range strings.Split(card[1], " ") {
			_, exist := winning[n]
			if exist {
				win++
			}
		}
		if win > 0 {
			total += math.Pow(2.0, win-1)
		}
	}
	fmt.Println(total)
}

