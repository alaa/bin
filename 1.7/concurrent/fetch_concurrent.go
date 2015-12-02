package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("-- \nTotal time: %.2fs\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	url = fixPrefix(url)

	resp, err := http.Get(url)
	defer resp.Body.Close()

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error: %s", err)
	}

	ch <- fmt.Sprintf("%.2fs sec \t=>\t %s \t %d KB", time.Since(start).Seconds(), url, nbytes/1024)
}

func fixPrefix(url string) string {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	return url
}
