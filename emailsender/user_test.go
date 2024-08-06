package emailsender

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestNewUserList(t *testing.T) {
	ul := NewUserList()
	if ul == nil {
		t.Fatal("NewUserList() returned nil")
	}
	if ul.Users == nil {
		t.Error("NewUserList() returned UserList with nil Users slice")
	}
	if len(ul.Users) != 0 {
		t.Errorf("NewUserList() returned UserList with %d users, expected 0", len(ul.Users))
	}
}

func TestUserList_ReadUsers(t *testing.T) {
	// Create a temporary JSON file for testing
	tempFile, err := os.CreateTemp("", "test_users_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	testUsers := []User{
		{Username: "User1", Email: "user1@example.com", Bufferindex: 0},
		{Username: "User2", Email: "user2@example.com", Bufferindex: 1},
	}
	json.NewEncoder(tempFile).Encode(testUsers)
	tempFile.Close()

	ul := NewUserList()
	ul.ReadUsers(tempFile.Name())

	if len(ul.Users) != len(testUsers) {
		t.Errorf("ReadUsers() loaded %d users, expected %d", len(ul.Users), len(testUsers))
	}

	if !reflect.DeepEqual(ul.Users, testUsers) {
		t.Errorf("ReadUsers() loaded incorrect user data")
	}
}

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
