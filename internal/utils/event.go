package utils

import (
	"errors"

	"github.com/robertd2000/github-activity/internal/model"
)

func ParseEventType(input string) (model.EventType, error) {
	switch input {
	case string(model.PushEvent),
		string(model.CreateEvent),
		string(model.ForkEvent),
		string(model.IssueCommentEvent),
		string(model.IssuesEvent),
		string(model.WatchEvent),
		string(model.PullRequestEvent),
		string(model.PullRequestReviewEvent),
		string(model.PullRequestReviewCommentEvent),
		string(model.DeleteEvent):
		return model.EventType(input), nil
	default:
		return "", errors.New("недопустимый тип события")
	}
}