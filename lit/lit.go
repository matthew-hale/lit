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
	for _, line := range input {
		if fileStartMatch.MatchString(line) {
			fmt.Println(strings.Trim(line, "```"))
		}
	}
}
