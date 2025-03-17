package model

type EventType string

const (
	PushEvent                     EventType = "PushEvent"
	CreateEvent                   EventType = "CreateEvent"
	ForkEvent                     EventType = "ForkEvent"
	IssueCommentEvent             EventType = "IssueCommentEvent"
	IssuesEvent                   EventType = "IssuesEvent"
	WatchEvent                    EventType = "WatchEvent"
	PullRequestEvent              EventType = "PullRequestEvent"
	PullRequestReviewEvent        EventType = "PullRequestReviewEvent"
	PullRequestReviewCommentEvent EventType = "PullRequestReviewCommentEvent"
	DeleteEvent                   EventType = "DeleteEvent"
)