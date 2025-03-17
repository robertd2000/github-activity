package model

import "time"

type Activity struct {
	Type      EventType  `json:"type"`
	Repo      Repo   `json:"repo"`
	CreatedAt time.Time `json:"created_at"`
	Payload   struct {
		Action  string `json:"action"`
		Ref     string `json:"ref"`
		RefType string `json:"ref_type"`
		Commits []struct {
			Message string `json:"message"`
		} `json:"commits"`
	} `json:"payload"`
}

type Repo struct {
	Name string `json:"name"`
}