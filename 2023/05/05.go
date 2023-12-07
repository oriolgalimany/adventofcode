package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

var almanac [][][]int64
var seeds []int64
var conversion []map[int64]int64

func load() {
	file, _ := os.Open("input")
	defer file.Close()

	var mapa [][]int64
	re := regexp.MustCompile("\\d+")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == "" {
			almanac = append(almanac, mapa)
			mapa = nil
			continue
		}
		digits_str := re.FindAllString(scanner.Text(), -1)
		if len(digits_str) == 0 { //text
			continue
		}

		var digits []int64
		for _, d := range digits_str {
			digit, _ := strconv.ParseInt(d, 10, 64)
			digits = append(digits, digit)
		}
		mapa = append(mapa, digits)
	}
	almanac = append(almanac, mapa)
}

func convert(mapId int, value int64) int64 {
	var mapped int64 = value
	for _, a := range almanac[mapId] {
		source := a[1]
		destiny := a[0]
		scope := a[2]

		if value >= source && value <= source+scope {
			delta := destiny - source
			mapped = value + delta
			return mapped
		}
	}
	return mapped
}

func main() {
	load()

	seeds = almanac[0][0]
	var location = int64(math.MaxInt64)
	for _, seed := range seeds {
		soil := convert(1, seed)
		fert := convert(2, soil)
		water := convert(3, fert)
		light := convert(4, water)
		temp := convert(5, light)
		hum := convert(6, temp)
		loc := convert(7, hum)

		/* fmt.Println("seed", seed)
		fmt.Println("soil", soil)
		fmt.Println("fert", fert)
		fmt.Println("water", water)
		fmt.Println("light", light)
		fmt.Println("temperature", temp)
		fmt.Println("humidity", hum)
		fmt.Println("=") */
		// fmt.Println("location", loc)
		// fmt.Println("-----")
		if location > loc {
			location = loc
		}
	}
	fmt.Println(location)
}
