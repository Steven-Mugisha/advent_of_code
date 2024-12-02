package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const QUESTION = `https://adventofcode.com/2024/day/1`


func main() {

	left, right, err := handleLocationData("./day1.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	// part1: total distance between the two lists
	totalDistance := calcTotalDistance(left, right)
	fmt.Println("Total distance: ", totalDistance)

	// part2: similarity score
	similarityScore := calcSimilarityScore(left, right)
	fmt.Println("Similarity score: ", similarityScore)


}

func readLines (filename string) ([]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func handleLocationData(filename string) ([]int , []int, error) {

	lines, err := readLines(filename)
	if err != nil {
		return nil, nil, err
	}

	left := []int{}
	right := []int{}

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line: %s", line)
		}

		leftVal, err := strconv.Atoi(parts[0])

		if err != nil {
            return nil, nil, err
        }

		rightVal, err := strconv.Atoi(parts[1])

		if err != nil {
			return nil, nil, err
		}

		left = append(left, leftVal)
		right = append(right, rightVal)
	
	}

	return left, right, nil
}


func calcTotalDistance(left []int, right []int) int {

	sort.Ints(left)
	sort.Ints(right)

	res := []int{}

	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		absDiff := int(math.Abs(float64(diff)))
		res = append(res, absDiff)
	}

	total := 0
	for _, v := range res {
		total += v
	}

	return total
}

func calcSimilarityScore(left []int, right []int) int {
	// initialize the map:
	var similarityMap = make(map[int]int)

	for i := 0; i < len(right); i++ {
		if _,ok := similarityMap[right[i]]; ok {
			similarityMap[right[i]]++
		} else {
			similarityMap[right[i]] = 1
		}
	}

	simScore := 0

	for _, v := range left {
		simScore += v * similarityMap[v]
	}
	return simScore
}
