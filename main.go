package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("Please provide a username")
	}

	username := args[0]
	
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)
}