package main

import (
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("Please provide a username")
	}

	username := args[0]

	

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)
}