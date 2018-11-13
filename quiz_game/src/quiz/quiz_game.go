package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func readCSV(f string) (map[string]int, error) {
	absPath, _ := filepath.Abs(f)
	fmt.Println(absPath)

	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	// Need to close it to prevent leaking
	defer file.Close()

	// Read CSV from csv package
	csvr := csv.NewReader(file)

	// Map for storing the Questions (string) and Answers (int)
	qAndA := map[string]int{}

	// Iterate all rows of CSV and get the Questions (key) and Answers (value)
	for {
		row, err := csvr.Read()
		if err != nil {
			// Read until End-of-file
			if err == io.EOF {
				err = nil
			}
			return qAndA, err
		}
		result, err := strconv.Atoi(row[1])
		// Check for error in case string cannot be converted to integer
		if err == nil {
			qAndA[row[0]] = result
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	problems, _ := readCSV("problems.csv")
	finalResult := 0
	for k, v := range problems {
		fmt.Printf("%s: ", k)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		result, err := strconv.Atoi(text)
		// Error check
		if err != nil {
			continue
		}
		// Comparing both answers
		if result == v {
			finalResult++
		}
	}
	fmt.Printf("You have got %v out of %v questions correct!\n", finalResult, len(problems))
}
