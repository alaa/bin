package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := fetch(url)
		if err != nil {
			log.Fatalf("Cloud not fetch url %s, Error: %s", url, err)
		}

		if _, err := io.Copy(os.Stdout, resp); err != nil {
			log.Fatalf("Error: ", err)
		}
	}
}

func fetch(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bodyReader := bytes.NewReader(body)

	if err != nil {
		return bodyReader, err
	}
	return bodyReader, nil
}
