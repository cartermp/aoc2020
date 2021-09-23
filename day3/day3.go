package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readData() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func doWork(data []string, right int, down int) int {
	treeCount := 0
	tobogganIdx := right
	first := true

	for i := 0; i < len(data); i += down {
		if first {
			first = false
			continue
		}

		line := data[i]
		toCheck := tobogganIdx % len(line)
		if line[toCheck] == '#' {
			treeCount++
		}

		tobogganIdx += right
	}

	return treeCount
}

func main() {
	data := readData()

	one := doWork(data, 1, 1)
	two := doWork(data, 3, 1)
	three := doWork(data, 5, 1)
	four := doWork(data, 7, 1)
	five := doWork(data, 1, 2)

	res := one * two * three * four * five
	fmt.Println(res)
}
