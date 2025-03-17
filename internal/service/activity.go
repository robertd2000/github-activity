package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/robertd2000/github-activity/internal/model"
	"github.com/robertd2000/github-activity/internal/utils"
)

func FetchActivity(username string) []model.Activity {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalln("Failed to fetch activity")
	}

	var activities []model.Activity

	if err := json.NewDecoder(resp.Body).Decode(&activities); err != nil {
		log.Fatalln(err)
	}

	return activities
}

func DisplayActivity(activities []model.Activity) {
	for _, activity := range activities {
		switch activity.Type {
		case "PushEvent":
			fmt.Println("Pushed to", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt), "with", len(activity.Payload.Commits), "commits")
		case "CreateEvent":
			fmt.Println("Created", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt))
		case "ForkEvent":
			fmt.Println("Forked", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt))
		case "IssueCommentEvent":
			fmt.Println("Commented on", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt))
		case "IssuesEvent":
			fmt.Println("Created an issue on", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt))
		case "WatchEvent":
			fmt.Println("Starred ", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt))
		case "PullRequestEvent":
			fmt.Println("Created pull request", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt))
		case "PullRequestReviewEvent":
			fmt.Println("Reviewed pull request ", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt))

		}
		// fmt.Println(activity.Type, activity.Repo.Name, activity.CreatedAt, activity.Payload.Action, activity.Payload.Ref, activity.Payload.RefType, activity.Payload.Commits)

	}
}