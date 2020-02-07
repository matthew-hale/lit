package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

func stdinput() []string {
	in := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in = append(in, scanner.Text())
	}
	return in
}

func litOpen(name string) *os.File {
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	return f
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	input := stdinput()
	fileStartMatch := regexp.MustCompilePOSIX("^```[a-zA-Z0-9 _\\-]+\\.?[a-zA-Z]*$")
	fileEndMatch := regexp.MustCompilePOSIX("^```$")
	fileIndexes := make([]int, 0)
	for i, line := range input {
		if fileStartMatch.MatchString(line) {
			fileIndexes = append(fileIndexes, i)
		}
	}
	for _, index := range fileIndexes {
		filename := strings.Trim(input[index], "```")

		f := litOpen(filename)

		// For each index, we're going to want to write all lines to a file 
		// until we hit the end of the code block.
		for _, line := range input[index+1:] {
			if fileEndMatch.MatchString(line) {
				break
			}
			// Write line to file
			_, err := f.WriteString(line + "\n")
			check(err)
		}

		f.Sync()
		f.Close()
	}
}
