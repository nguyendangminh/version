package main

import (
	"log"
	"os"

	"github.com/nguyendangminh/version"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "version" {
		version.Print()
		return
	}
	log.Println("error: invalid command")
	log.Printf("usage: %s version", os.Args[0])
}
