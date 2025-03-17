package main

import (
	"log"
	"os"

	"github.com/robertd2000/github-activity/internal/service"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("Please provide a username")
	}

	filterType := ""

	if len(args) == 2 {
		filterType = args[1]
	}

	username := args[0]

	activities := service.FetchActivity(username)
	service.DisplayActivity(activities, filterType)
}