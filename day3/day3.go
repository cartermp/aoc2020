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
	minOcc         int
	maxOcc         int
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

		minPart := regexp.MustCompile("[0-9]+-").FindString(line)
		maxPart := regexp.MustCompile("-[0-9]+").FindString(line)
		letterPart := regexp.MustCompile("[a-z]:").FindString(line)
		passwordPart := regexp.MustCompile(": [a-zA-Z]+").FindString(line)

		min, minErr := strconv.Atoi(strings.Trim(minPart, "-"))
		max, maxErr := strconv.Atoi(strings.Trim(maxPart, "-"))
		letter := strings.Trim(letterPart, ":")
		password := strings.Trim(passwordPart, ": ")

		if minErr != nil || maxErr != nil {
			continue
		}

		data :=
			passwordData{
				minOcc:         min,
				maxOcc:         max,
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
		letterCount := strings.Count(passwordData.password, passwordData.requiredLetter)
		if letterCount >= passwordData.minOcc && letterCount <= passwordData.maxOcc {
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
