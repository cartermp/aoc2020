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

type PasswordData struct {
	minOcc         int
	maxOcc         int
	requiredLetter string
	password       string
	fullText       string
}

func parseData() []PasswordData {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	passwords := make([]PasswordData, 0)

	for scanner.Scan() {
		line := scanner.Text()

		minPart := regexp.MustCompile("[0-9]+-").FindString(line)
		maxPart := regexp.MustCompile("-[0-9]+").FindString(line)
		letterPart := regexp.MustCompile("[a-z]:").FindString(line)
		passwordPart := regexp.MustCompile(": [a-zA-Z]+").FindString(line)

		min, _ := strconv.Atoi(strings.Trim(minPart, "-"))
		max, _ := strconv.Atoi(strings.Trim(maxPart, "-"))
		letter := strings.Trim(letterPart, ":")
		password := strings.Trim(passwordPart, ": ")

		data :=
			PasswordData{
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

func getValidPasswords(passwords []PasswordData) []PasswordData {
	validPasswords := make([]PasswordData, 0)

	for i := 0; i < len(passwords); i++ {
		passwordData := passwords[i]
		letterCount := strings.Count(passwordData.fullText, passwordData.requiredLetter)
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
