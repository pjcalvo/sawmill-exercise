package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	maxCut = 3
)

func getLines() ([][]int, error) {
	// Open the file for reading
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, nil
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	dataLines := make([][]int, 0)
	// collect all the data lines
	for scanner.Scan() {
		lineStr := scanner.Text()
		valuesLine := strings.Split(lineStr, " ")
		singles := make([]int, 0)
		for _, v := range valuesLine {
			value, _ := strconv.Atoi(v)
			singles = append(singles, value)
		}
		dataLines = append(dataLines, singles)
	}

	// Check if there were any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
	return dataLines, nil
}

func processLines(lines [][]int) error {
	caseNumber := 1
	maxProfit := 0
	for _, v := range lines {
		stucks := make([]int, 0)
		stuckProfit := 0
		// if it is a new case
		if len(v) == 1 {
			// if last case
			if v[0] == 0 {
				return nil
			}
			fmt.Printf("Case %v\n", caseNumber)
			caseNumber++
			stuckProfit = 0
			maxProfit = 0
			stucks = make([]int, 0)
			// the next amount of lines belong to the same sawnline
			continue
		}

		// process line
		left := 0
		for _, v := range v[1:] {
			stucks = append(stucks, left)
			v = v - left
			remaining := v
			for remaining >= maxCut {
				stucks = append(stucks, maxCut)
				remaining = remaining - maxCut
			}
			stucks = append(stucks, remaining)
			left = maxCut - remaining

			// calculate line profit
			for _, s := range stucks {
				switch s {
				case 0:
					continue
				case 1:
					stuckProfit -= 1
				case 2:
					stuckProfit += 3
				case 3:
					stuckProfit += 1
				default:
					return fmt.Errorf("screw up something")
				}
			}
			maxProfit = maxProfit + stuckProfit
			stucks = make([]int, 0)
		}
		fmt.Printf("Max Profit: %v\n", stuckProfit)
	}
	return nil
}

func main() {
	dataLines, err := getLines()
	if err != nil {
		panic(err)
	}

	err = processLines(dataLines)
	if err != nil {
		panic(err)
	}
}
