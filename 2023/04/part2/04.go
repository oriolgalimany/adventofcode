package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cards []*Card

type Card struct {
	winningNumbers map[string]struct{}
	ourNumbers     []string
	copies         []*Card
	wins           int
}

func play() {
	total := 0
	for i, card := range cards {
		total++
		//process copies
		for _, cp := range card.copies {
			for nwins := 1; nwins <= cp.wins; nwins++ {
				cards[i+nwins].copies = append(cards[i+nwins].copies, cards[i+nwins])
				total++
			}
		}
		//process card
		for nwins := 1; nwins <= card.wins; nwins++ {
			cards[i+nwins].copies = append(cards[i+nwins].copies, cards[i+nwins])
			total++
		}
	}

	fmt.Println(total)
}

func load() {
	file, _ := os.Open("../input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ":")
		cardSplit := strings.Split(s[1], "|")

		winning := make(map[string]struct{}) //only key map , struct{} is 0 bytes

		for _, w := range strings.Split(cardSplit[0], " ") {
			if w != "" {
				winning[w] = struct{}{}
			}
		}

		win := 0

		ourNumbers := strings.Split(cardSplit[1], " ")
		for _, n := range ourNumbers {
			_, exist := winning[n]
			if exist {
				win++
			}
		}
		card := new(Card)
		card.winningNumbers = winning
		card.ourNumbers = ourNumbers
		card.wins = win

		cards = append(cards, card)
	}
}

func main() {
	load()
	play()
}
