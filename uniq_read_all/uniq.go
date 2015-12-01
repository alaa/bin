package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	dic := make(map[string]int)

	for _, f := range validFiles() {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}

		for _, line := range strings.Split(string(data), "\n") {
			dic[line]++
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
