package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	lines := fileToLines("input.txt")
	elves, max := linesToElvesAndMax(lines)
	fmt.Println("Max", max)
	top3Total := getTopThreeSum(elves)
	fmt.Println("Total Top Three", top3Total)
}

func fileToLines(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return
}

func linesToElvesAndMax(lines []string) (elves []int, mostCalories int) {
	elf := 0
	max := 0
	for _, line := range lines {
		num, _ := strconv.ParseInt(line, 0, 64)
		if num > 0 {
			elf += int(num)
		} else {
			// blank line

			// We should total up and collect the maximum
			if elf > max {
				max = elf
			}
			// Insert elf into array
			elves = append(elves, elf)
			// make a new elf
			elf = 0
		}
	}
	// add the last  elf
	// Insert elf into array
	elves = append(elves, elf)

	return elves, max
}

func getTopThreeSum(elves []int) int {
	sort.Ints(elves)
	ReverseSlice(elves)
	return elves[0] + elves[1] + elves[2]
}

func ReverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}
