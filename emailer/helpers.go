package emailer

import "time"

const (
	StartYear  = 2024
	StartMonth = 7
	StartDay   = 22
)

func GetWeeksUntil(currentDate time.Time) int {
	startDate := time.Date(StartYear, StartMonth, StartDay, 0, 0, 0, 0, time.UTC)
	weeks := currentDate.Sub(startDate).Hours() / 24 / 7
	return int(weeks)
}
