package main

import (
	"log"
	"os"

	"github.com/robertd2000/github-activity/internal/model"
	"github.com/robertd2000/github-activity/internal/service"
	"github.com/robertd2000/github-activity/internal/utils"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("Please provide a username")
	}

	var filterType model.EventType
	var err error

	if len(args) == 2 {
		filterType, err = utils.ParseEventType(args[1])
		if err != nil {
			log.Fatalln("Please provide a valid event type")
		}
	}

	username := args[0]

	activities := service.FetchActivity(username)
	service.DisplayActivity(activities, filterType)
}