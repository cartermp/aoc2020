package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passwordData struct {
	firstPos       int
	secondPos      int
	requiredLetter string
	password       string
	fullText       string
}

func parseData() []passwordData {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	passwords := make([]passwordData, 0)

	for scanner.Scan() {
		line := scanner.Text()

		firstPosPart := regexp.MustCompile("[0-9]+-").FindString(line)
		secondPosPart := regexp.MustCompile("-[0-9]+").FindString(line)
		letterPart := regexp.MustCompile("[a-z]:").FindString(line)
		passwordPart := regexp.MustCompile(": [a-zA-Z]+").FindString(line)

		firstPos, firstPosErr := strconv.Atoi(strings.Trim(firstPosPart, "-"))
		secondPos, secondPosErr := strconv.Atoi(strings.Trim(secondPosPart, "-"))
		letter := strings.Trim(letterPart, ":")
		password := strings.Trim(passwordPart, ": ")

		if firstPosErr != nil || secondPosErr != nil {
			continue
		}

		data :=
			passwordData{
				firstPos:       firstPos,
				secondPos:      secondPos,
				requiredLetter: letter,
				password:       password,
				fullText:       line,
			}

		passwords = append(passwords, data)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return passwords
}

func getValidPasswords(passwords []passwordData) []passwordData {
	validPasswords := make([]passwordData, 0)

	for _, passwordData := range passwords {
		requiredLetters := 0
		for idx, letter := range passwordData.password {
			if idx == (passwordData.firstPos - 1) && string(letter) == passwordData.requiredLetter {
				requiredLetters++
			}

			if idx == (passwordData.secondPos - 1) && string(letter) == passwordData.requiredLetter {
				requiredLetters++
				break
			}
		}

		if requiredLetters == 1 {
			validPasswords = append(validPasswords, passwordData)
		}
	}

	return validPasswords
}

func main() {
	passwords := parseData()
	result := getValidPasswords(passwords)
	fmt.Println(len(result))
}
