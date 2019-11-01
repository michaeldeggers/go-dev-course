package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// read function will read info into the byte slice, but cannot resize it
	_, _ = io.Copy(os.Stdout, resp.Body)
}
