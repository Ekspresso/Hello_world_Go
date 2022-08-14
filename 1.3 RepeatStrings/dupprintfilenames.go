// Выводит имена файлов с повторяющимеся значениями

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	filenames := ""
	if len(files) == 0 {
		if countLines(os.Stdin, counts) > 0 {
			filenames += "Stdin"
		}
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			if countLines(f, counts) > 0 {
				filenames += arg + " "
			}
			f.Close()
		}
	}
	fmt.Printf("Files with repeat strings: %s\n", filenames)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) int {
	input := bufio.NewScanner(f)
	c := 0
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 0 {
			c++
		}
	}
	return (c)
}
