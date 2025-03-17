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
		log.Fatalln("Failed to fetch activity for", username, "-", resp.StatusCode)
	}

	var activities []model.Activity

	if err := json.NewDecoder(resp.Body).Decode(&activities); err != nil {
		log.Fatalln(err)
	}

	return activities
}

func DisplayActivity(activities []model.Activity, filterType model.EventType) {
	if filterType != "" {
		activities = utils.Filter(activities, func(activity model.Activity) bool {
			return activity.Type == filterType
		})
	}
	for _, activity := range activities {
		switch activity.Type {
		case model.PushEvent:
			fmt.Println("- Pushed to", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt), "with", len(activity.Payload.Commits), "commits")
		case model.CreateEvent:
			fmt.Println("- Created", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt))
		case model.ForkEvent:
			fmt.Println("- Forked", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt))
		case model.IssueCommentEvent:
			fmt.Println("- Commented on", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt))
		case model.IssuesEvent:
			fmt.Println("- Created an issue on", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt))
		case model.WatchEvent:
			fmt.Println("- Starred ", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt))
		case model.PullRequestEvent:
			fmt.Println("- Created pull request", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt))
		case model.PullRequestReviewEvent:
			fmt.Println("- Reviewed pull request", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt))
		case model.PullRequestReviewCommentEvent:
			fmt.Println("- Commented on pull request", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt))
		case model.DeleteEvent:
			fmt.Println("- Delete from", activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt))
		default:
			fmt.Println("-", activity.Type, activity.Repo.Name, "at", utils.FormatDate(activity.CreatedAt))
		}
	}
}