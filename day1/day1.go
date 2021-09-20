package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"strconv"
)

func parseData(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	buf := make([]int, 0)

    for scanner.Scan() {
		text := scanner.Text()
		value, err := strconv.Atoi(text)
		if err == nil {
			buf = append(buf, value)
		} else {
			fmt.Println(err)
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return buf
}

func getResult(buf []int) int {
	result := 0

	// hey at least it's not bubble sort
	for i := 0; i < len(buf); i++ {
		for j := i; j < len(buf); j++ {
			for k := i; k < len(buf); k++ {
				if buf[i] + buf[j] + buf[k] == 2020 {
					result = buf[i] * buf[j] * buf[k]
				}
			}
		}
	}

	return result
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    buf := parseData(file)
	result := getResult(buf)
	
	fmt.Println(result)
}