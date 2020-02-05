package main

import (
	"fmt"
	"bufio"
	"os"
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
		for _, line := range input[index+1:] {
			if fileEndMatch.MatchString(line) {
				break
			}
			fmt.Println(line)
		}
	}
}
