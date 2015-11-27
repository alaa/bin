package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dic := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		dic[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	for k, v := range dic {
		if dic[k] > 1 {
			fmt.Printf("%d %s\n", v, k)
		}
	}
}
