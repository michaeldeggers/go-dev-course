package main

import (
	"github.com/michaeldeggers/go-dev-course/utils"
	"io"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	lw := utils.LogWriter{}
	_, _ = io.Copy(lw, file)
}
