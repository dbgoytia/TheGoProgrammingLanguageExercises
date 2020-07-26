package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Store the names of the files, and the words and times the owrd has repeated
	// { word: {filename : times}}
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines2(os.Stdin, counts, "os.Stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines2(f, counts, arg)
			f.Close()
		}
	}

	for line, filenames := range counts {

		// Count the number of files in which a line appears in:
		fileCount := len(filenames)

		// Count total number of appearances
		var total = 0

		fmt.Printf("Found [%s] in [%d] files:\n", line, fileCount)
		for name, count := range filenames {
			fmt.Printf("\t%d appearances in %s\n", count, name)
			total += count
		}
		fmt.Printf("\t->Total appearances: %d\n", total)

	}

}

func countLines2(f *os.File, counts map[string]map[string]int, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][filename]++
	}
}
