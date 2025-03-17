package utils

import "time"

func FormatDate(date time.Time) string {
	return date.Format("02.01.2006 15:04:05")
}