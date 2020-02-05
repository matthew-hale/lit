package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"regexp"
	"strings"
)

func input() []string {
	in := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in = append(in, scanner.Text())
	}
	return in
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input := input()
	fileStartMatch := regexp.MustCompilePOSIX("^```[a-zA-Z0-9 _\\-]+\\.?[a-zA-Z]*$")
	fileEndMatch := regexp.MustCompilePOSIX("^```$")
	fileIndexes:= make([]int, 0)
	for i, line := range input {
		if fileStartMatch.MatchString(line) {
			fileIndexes = append(fileIndexes, i)
		}
	}
	for _, index := range fileIndexes {
		filename := strings.Trim(input[index], "```")
		fmt.Println(filename)

		if _, err := os.Stat(filename); err == nil {
			// File exists; we open it
			f, err := os.Open(filename)
			check(err)
			defer f.Close()
		} else if os.IsNotExist(err) {
			// File doesn't exist; we create it
			f, err := os.Create(filename)
			check(err)
			defer f.Close()
		} else {
			// Some other error occured; we log.Fatal
			log.Fatal(err)
		}

		for _, line := range input[index+1:] {
			if fileEndMatch.MatchString(line) {
				break
			}
			fmt.Println(line)
		}
	}
}
