package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func fynd(pattern string, filePath string) {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid regex pattern: %v\n", err)
		return
	}

	var reader *bufio.Reader

	if filePath == "-" {
		reader = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			return
		}
		defer file.Close()
		reader = bufio.NewReader(file)
	}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			}
			break
		}
		line = strings.TrimSpace(line)
		if regex.MatchString(line) {
			fmt.Println(line)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: fynd <pattern> [file]")
		return
	}

	pattern := os.Args[1]
	filePath := "-"
	if len(os.Args) > 2 {
		filePath = os.Args[2]
	}
	fynd(pattern, filePath)
}
