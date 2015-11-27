package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	dic := make(map[string]int)

	for _, f := range validFiles() {
		file, err := os.Open(f)
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			dic[scanner.Text()]++
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}

	for k, v := range dic {
		if dic[k] > 1 {
			fmt.Printf("%d %s\n", v, k)
		}
	}
}

func validFiles() (valid []string) {
	args := os.Args[1:]
	for _, arg := range args {
		if validFile(arg) {
			valid = append(valid, arg)
		}
	}
	return
}

func validFile(filepath string) bool {
	if _, err := os.Stat(filepath); err != nil {
		return false
	}
	return true
}
