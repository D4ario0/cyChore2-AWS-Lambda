package utils

import (
	"testing"
	"time"
)

func TestGetWeeksUntil(t *testing.T) {
	testCases := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{"6 weeks after start", time.Date(2024, 9, 2, 0, 0, 0, 0, time.UTC), 6},
		{"1 day after start", time.Date(2024, 7, 23, 0, 0, 0, 0, time.UTC), 0},
		{"1 week before start", time.Date(2024, 7, 15, 0, 0, 0, 0, time.UTC), -1},
		{"exact start date", time.Date(2024, 7, 22, 0, 0, 0, 0, time.UTC), 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetWeeksUntil(&tc.date)
			if got != tc.expected {
				t.Errorf("GetWeeksUntil() = %d; want %d", got, tc.expected)
			}
		})
	}
}
