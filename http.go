package main

import (
	"fmt"
	"net/http"
	"os"
)

func get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(resp.Body)
}
