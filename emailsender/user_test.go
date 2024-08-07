package emailsender

import (
	"testing"
)

func TestUser_AssignTasks(t *testing.T) {
	testCases := []struct {
		name     string
		user     User
		weeks    int
		expected string
	}{
		{"Week 0", User{Bufferindex: 0}, 0, Task1},
		{"Week 1", User{Bufferindex: 0}, 1, Task2},
		{"Week 2", User{Bufferindex: 0}, 2, "Limpieza profunda (area común)"},
		{"Week 3", User{Bufferindex: 0}, 3, Task4},
		{"Week 4", User{Bufferindex: 0}, 4, Task1},
		{"Week 5", User{Bufferindex: 0}, 5, Task2},
		{"Week 6", User{Bufferindex: 0}, 6, "Limpieza profunda (cocina)"},
		{"Week 7", User{Bufferindex: 0}, 7, Task4},
		{"Week 8", User{Bufferindex: 0}, 8, Task1},
		{"Different Buffer Index", User{Bufferindex: 1}, 0, Task2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.user.AssignTasks(tc.weeks)
			if tc.user.Task != tc.expected {
				t.Errorf("AssignTasks(%d) = %s, expected %s", tc.weeks, tc.user.Task, tc.expected)
			}
		})
	}
}
