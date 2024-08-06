package emailsender

import (
	"encoding/json"
	"fmt"
	"os"
)

type UserList struct {
	Users []User `json:"users"`
}

type User struct {
	Username    string `json:"name"`
	Email       string `json:"email"`
	Bufferindex int    `json:"index"`
	Task        string
}

const (
	Task1 = "Sacar Basura / Correspondencia"
	Task2 = "Limpieza profunda (%s)"
	Task3 = "Barrer área común / Limpiar Baño"
	Task4 = "Barrer Cocina"
)

func NewUserList() *UserList {
	return &UserList{Users: []User{}}
}

func (ul *UserList) ReadUsers(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(ul)
	if err != nil {
		return fmt.Errorf("error decoding JSON: %w", err)
	}

	return nil
}

func (u *User) AssignTasks(weeks int) {
	var specialTask string

	cycleOffset := (weeks % 8) / 4 //This returns an integer 0 or 1 depending on current chore cycle

	if weeks%2 == cycleOffset {
		specialTask = fmt.Sprintf(Task2, "area común")
	} else {
		specialTask = fmt.Sprintf(Task2, "cocina")
	}

	tasks := []string{Task1, specialTask, Task3, Task4}
	taskIndex := u.Bufferindex + weeks
	u.Task = tasks[taskIndex%4]
}
