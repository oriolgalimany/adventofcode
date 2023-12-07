package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"runtime"
	"strconv"
)

var almanac [][][]int64
var seeds []int64

func load() {
	file, _ := os.Open("../input")
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

func getSeeds() []int64 {
	var seeds []int64
	for i := 0; i < len(almanac[0][0]); i = i + 2 {
		for j := int64(0); j < almanac[0][0][i+1]; j++ {
			seed := almanac[0][0][i] + j
			seeds = append(seeds, seed)
		}
		runtime.GC() // brute force :3
	}
	return seeds
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

	var location = int64(math.MaxInt64)
	seeds := almanac[0][0]
	for i := 0; i < len(seeds); i = i + 2 {
		for j := int64(0); j < seeds[i+1]; j++ {
			soil := convert(1, seeds[i]+j)
			fert := convert(2, soil)
			water := convert(3, fert)
			light := convert(4, water)
			temp := convert(5, light)
			hum := convert(6, temp)
			loc := convert(7, hum)

			if location > loc {
				location = loc
			}

		}
	}

	fmt.Println(location)
}
