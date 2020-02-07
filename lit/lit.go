package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// Parses input from stdin
func stdInput() []string {
	in := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in = append(in, scanner.Text())
	}
	return in
}

// Parses input from file
func fileInput(name string) []string {
	in := make([]string, 0)
	f, err := os.Open(name)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
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
	// These are our command line flags
	overwritePtr := flag.Bool("f", false ,"force complete overwrite of named script files (default is to append)")
	//directoryPtr := flag.String("o", "./", "output directory (defaults to current working directory)")
	inputPtr := flag.String("i", "", "input file path (defaults to stdin)")

	flag.Parse()

	// First we determine the output location; if no input flag was 
	// provided, we'll dump everything to the current directory
	input := make([]string, 0)
	if *inputPtr != "" {
		file, err := os.Stat(*inputPtr)
		check(err)
		switch mode := file.Mode(); {
			case mode.IsDir():
				fmt.Fprintf(os.Stderr, "%s is a directory; -i requires a file; exiting\n", *inputPtr)
				os.Exit(1)
			case mode.IsRegular():
				input = fileInput(*inputPtr)
		}
	} else {
		input = stdInput()
	}


	// Regex definitions
	fileStartMatch := regexp.MustCompilePOSIX("^```[a-zA-Z0-9 _\\-]+\\.?[a-zA-Z]*$")
	fileEndMatch := regexp.MustCompilePOSIX("^```$")

	// Here we gather all of the indexes where file names were found in 
	// our input.
	fileIndexes := make([]int, 0)
	for i, line := range input {
		if fileStartMatch.MatchString(line) {
			fileIndexes = append(fileIndexes, i)
		}
	}

	// Here we handle the overwrite flag. If overwrite is set to true, 
	// we'll first want to clear out all of the files found in the input 
	// if they exist. This can be interpreted as "delete," since litOpen() 
	// handles the case of a file not existing.
	if *overwritePtr == true {
		// First, gather all of the unique file names
		filenames := make([]string, 0)
		for _, index := range fileIndexes {
			filename := strings.Trim(input[index], "```")
			// Append only if unique
			skip := false
			for _, ele := range filenames {
				if ele == filename {
					skip = true
					break
				}
			}
			if !skip {
				filenames = append(filenames, filename)
			}
		}

		// Next, remove them if they exist
		for _, filename := range filenames {
			if _, err := os.Stat(filename); err == nil {
				// File exists; if it's a regular file, remove it
				file, _:= os.Stat(filename)
				mode := file.Mode()
				if mode.IsRegular() {
					err := os.Remove(filename)
					check(err)
				}
			} else if os.IsNotExist(err) {
				// file doesn't exist; do nothing
			} else {
				// unknown file error
				check(err)
			}
		}
	}

	// Now we write to the files
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

		fmt.Fprintf(os.Stdout, "%s written\n", filename)
		f.Sync()
		f.Close()
	}
}
