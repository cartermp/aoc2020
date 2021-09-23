package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	// "strings"
)

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func parseData() []passport {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	passports := make([]passport, 0)

	currPassport := passport{}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			if (currPassport != passport{}) {
				passports = append(passports, currPassport)
				currPassport = passport{}
			}
			continue
		}

		if currPassport.byr == "" {
			currPassport.byr = regexp.MustCompile("byr:[0-9]+").FindString(line)
		}

		if currPassport.iyr == "" {
			currPassport.iyr = regexp.MustCompile("iyr:[0-9]+").FindString(line)
		}

		if currPassport.eyr == "" {
			currPassport.eyr = regexp.MustCompile("eyr:[0-9]+").FindString(line)
		}

		if currPassport.hgt == "" {
			currPassport.hgt = regexp.MustCompile("hgt:[a-zA-Z0-9]+").FindString(line)
		}

		if currPassport.hcl == "" {
			currPassport.hcl = regexp.MustCompile("hcl:#?[a-zA-Z]+").FindString(line)
		}

		if currPassport.ecl == "" {
			currPassport.ecl = regexp.MustCompile("ecl:[a-zA-Z]+").FindString(line)
		}

		if currPassport.pid == "" {
			currPassport.pid = regexp.MustCompile("pid:[0-9]+").FindString(line)
		}

		if currPassport.cid == "" {
			currPassport.cid = regexp.MustCompile("cid:[0-9]+").FindString(line)
		}
	}

	if (currPassport != passport{}) {
		passports = append(passports, currPassport)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return passports
}

func getValidPassports(passwords []passport) []passport {
	validPassports := make([]passport, 0)

	for _, pp := range passwords {
		if pp.byr == "" || pp.iyr == "" || pp.eyr == "" || pp.hgt == "" || pp.hcl == "" || pp.ecl == "" || pp.pid == "" {
			continue
		}

		validPassports = append(validPassports, pp)
	}

	return validPassports
}

func main() {
	parsedPassports := parseData()
	validPassports := getValidPassports(parsedPassports)
	fmt.Println(len(validPassports))
}
