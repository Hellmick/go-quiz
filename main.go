package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var timeLimit time.Duration

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

func askQuestion(question, answer string) bool {

	fmt.Println(question)

	timer := time.NewTimer(timeLimit)
	answerCh := make(chan string)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		answerCh <- strings.TrimSpace(input)
	}()

	select {

	case input := <-answerCh:
		return input == answer

	case <-timer.C:
		fmt.Println("Time's up!")
		os.Exit(0)

	}

	return false
}

func runQuiz(problems map[string]string) (score int) {

	for question, answer := range problems {

		if askQuestion(question, answer) {
			score += 1
		}
	}

	fmt.Println("The final score is", score, "out of", len(problems))

	return

}

func main() {

	fileLocation := flag.String("f", "problems.csv", "provides problem file location")
	questionTime := flag.Int("t", 30, "provides time to answer for each question")
	flag.Parse()

	problems := readCsv(*fileLocation)
	timeLimit = time.Duration(*questionTime) * time.Second

	runQuiz(problems)

}
