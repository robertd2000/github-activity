package model

type FilterType string

const (
	PushEvent                     FilterType = "PushEvent"
	CreateEvent                   FilterType = "CreateEvent"
	ForkEvent                     FilterType = "ForkEvent"
	IssueCommentEvent             FilterType = "IssueCommentEvent"
	IssuesEvent                   FilterType = "IssuesEvent"
	WatchEvent                    FilterType = "WatchEvent"
	PullRequestEvent              FilterType = "PullRequestEvent"
	PullRequestReviewEvent        FilterType = "PullRequestReviewEvent"
	PullRequestReviewCommentEvent FilterType = "PullRequestReviewCommentEvent"
	DeleteEvent                   FilterType = "DeleteEvent"
)