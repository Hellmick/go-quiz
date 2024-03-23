package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readCsv(filename string) map[string]string {

	problems := make(map[string]string)
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Error reading records", err)
	}

	for _, eachrecord := range records {
		problems[eachrecord[0]] = eachrecord[1]
	}

	return problems

}

func main() {
	problems := readCsv("problems.csv")

	for key, value := range problems {
		fmt.Println("The problem is", key, " and the answer is ", value)
	}
}
