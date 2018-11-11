package main

import (
	"fmt"
	"net/http"
)

func runAndTest(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("[FAIL] GET %s:\t%s\n", url, err.Error())
	} else {
		fmt.Printf("[ OK ] GET %s:\t%s\n", url, res.Status)
	}
}

func main() {
	runAndTest("http://example.com")
	runAndTest("https://example.com")
}
