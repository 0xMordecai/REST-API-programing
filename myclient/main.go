package main

import "net/http"

func main() {
	result, err := http.Get("http://localhost:8888")
	if err != nil {
		panic(err)
	}
}
