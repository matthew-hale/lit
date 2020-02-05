package main

import (
	"fmt"
	"bufio"
	"os"
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
	for _, line := range input {
		fmt.Println(line)
	}
}
