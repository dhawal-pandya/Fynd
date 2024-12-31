package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	words := []string{
		"lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit",
		"sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore",
		"magna", "aliqua", "ut", "enim", "ad", "minim", "veniam", "quis", "nostrud",
		"exercitation", "ullamco", "laboris", "nisi", "aliquip", "ex", "ea", "commodo",
		"consequat", "duis", "aute", "irure", "in", "reprehenderit", "voluptate", "velit",
		"esse", "cillum", "eu", "fugiat", "nulla", "pariatur", "excepteur", "sint",
		"occaecat", "cupidatat", "non", "proident", "sunt", "culpa", "qui", "officia",
		"deserunt", "mollit", "anim", "id", "est", "laborum",
	}

	generateSentence := func() string {
		var sentence []string
		wordCount := rand.Intn(8) + 5
		for i := 0; i < wordCount; i++ {
			word := words[rand.Intn(len(words))]
			sentence = append(sentence, word)
		}
		text := strings.Join(sentence, " ")
		text = strings.ToUpper(string(text[0])) + text[1:]
		return text
	}

	for {
		sentence := generateSentence()

		switch rand.Intn(3) {
		case 0:
			sentence += "."
		case 1:
			sentence += "?"
		case 2:
			sentence += "!"
		}

		fmt.Print(sentence)

		if rand.Intn(4) == 0 {
			fmt.Println() // makes it look like a paragraph
		} else {
			fmt.Print(" ")
		}

		time.Sleep(100 * time.Millisecond) // makes it look real
	}
}
