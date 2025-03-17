package model

type FilterType string

// Допустимые значения для EventType
const (
	PushEvent   FilterType = "PushEvent"
	CreateEvent FilterType = "CreateEvent"
	ForkEvent   FilterType = "ForkEvent"
)