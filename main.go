package main

import (
	"bufio"
	"encoding/csv"
	"flag"
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

func askQuestion(problem, answer string) bool {

	fmt.Println(problem)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	return input == answer

}

func askQuestions(problems map[string]string) int {
	score := 0
	for question, answer := range problems {
		if askQuestion(question, answer) {
			score += 1
		}
	}
	return score
}

func main() {

	fileLocation := flag.String("f", "problems.csv", "provides problem file location")
	flag.Parse()
	problems := readCsv(*fileLocation)

	fmt.Println("The final score is", askQuestions(problems), "out of", len(problems))

}
