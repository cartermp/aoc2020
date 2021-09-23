package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "regexp"
	// "strconv"
	// "strings"
)

func parseData(){
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	oneTime := false
	idx := 0
	treeCount := 0
	lineCount := 0
	for scanner.Scan() {
		if oneTime {
			oneTime = false
			idx += 3
			lineCount++
			continue
		}

		line := scanner.Text()
		
		if line[idx % len(line)] == '#' {
			treeCount++
		}

		lineCount++
		idx += 3
	}

	fmt.Printf("Tree count: %d\n", treeCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	parseData()
}
