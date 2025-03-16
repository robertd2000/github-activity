package service

import (
	"fmt"
	"log"
	"net/http"
)

func FetchActivity(username string) string {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}