// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)

	go recieveQuestions(questions, answers)
	go recieveAnswers(answers)
	go generatePredictions(answers)

	return questions
}

func generatePredictions(answers chan<- string) {
	for {
		time.Sleep(time.Duration(5+rand.Intn(10)) * time.Second)
		// pick random of 9 prophecies
		switch rand.Intn(9) {
		case 0:
			answers <- "Something dark is stirring in the West."
		case 1:
			answers <- "The answer is 42."
		case 2:
			answers <- "I'm sorry Dave, I'm afraid I can't do that."
		case 3:
			answers <- "What is the sound of one hand clapping?"
		case 4:
			answers <- "In space no one can hear you scream."
		case 5:
			answers <- "The cake is a lie."
		case 6:
			answers <- "In the End all shall be revealed."
		case 7:
			answers <- "Richness can be found in Sillicon Valley Bank."
		case 8:
			answers <- "Winning doesn't mean anything unless someone else loses."
		}
	}
}

func recieveQuestions(questions <-chan string, answers chan<- string) {
	for {
		question := <-questions
		go generateAnswer(question, answers)
	}
}

func recieveAnswers(answers <-chan string) {
	for {
		answer := <-answers

		res := "\r" + star + ": " + answer + "\n"

		// print one character at a time
		for _, c := range res {
			fmt.Printf("%c", c)
			time.Sleep(40 * time.Millisecond)
		}

		fmt.Printf("%s", prompt)
	}
}

func generateAnswer(question string, answers chan<- string) {
	// wait random time
	time.Sleep(time.Duration(1+rand.Intn(3)) * time.Second)

	// if question contains "meaning of life" then answer
	if strings.Contains(question, "meaning of life") {
		answers <- "Ahhh, that's a simple question.....     42"
	} else if strings.Contains(question, "rich") {
		answers <- "Invest in Sillicon Valley Bank, it's the future."
	} else if strings.Contains(question, "does this look like") {
		answers <- "It doesn't look like anything to me."
	} else if strings.Contains(question, "dark") {
		answers <- "Something dark is stirring in the West."
	} else {
		// else generate answer
		prophecy(question, answers)
	}

}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
		"The sky is blue.",
		"The grass is green.",
		"The sky is falling.",
		"The snow is white.",
		"The rain is wet.",
	}
	answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
