package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func readCSV(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var records []string
	reader := csv.NewReader(file)
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, rec...)
	}
	return records, nil
}

func main() {
	records, err := readCSV("AOC_Day1_Safe.csv")
	if err != nil {
		log.Fatal(err)
	}
	// Loop through the records and "spin the dial" on the safe. Count the number of times the dial lands on 0.
	var currentIndex int = 50
	var password int = 0

	for _, turn := range records {
		firstChar := string(turn[0])

		turnNum, err := strconv.Atoi(turn[1:])
		if err != nil {
			log.Fatal(err)
		}

		if firstChar == "L" {
			for i := 1; i <= turnNum; i++ {
				if currentIndex == 99 {
					currentIndex = 0
				} else {
					currentIndex++
				}
				if currentIndex == 0 {
					password++
				}
			}
		} else {
			for i := 1; i <= turnNum; i++ {
				if currentIndex == 0 {
					currentIndex = 99
				} else {
					currentIndex--
				}
				if currentIndex == 0 {
					password++
				}
			}
		}
	}

	fmt.Println(password)
}
