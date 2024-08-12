package types

// User represents a user in the system, with a name, email, buffer index, and task.
type User struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	BufferIndex int    `json:"index"`
	Task        string
}

// UserList holds a collection of users.
type UserList struct {
	Users []User `json:"users"`
}

// UserProcessor is a function type that processes a User and returns an error if the processing fails.
type UserProcessor func(*User) error

// ForEach applies the UserProcessor to each User in the UserList.
// It returns a slice of errors if any processing steps fail, or nil if all succeed.
func (ul *UserList) ForEach(processor UserProcessor) []error {
	var errors []error

	for i := range ul.Users {
		if err := processor(&ul.Users[i]); err != nil {
			errors = append(errors, err)
		}
	}

	// Return nil if there are no errors.
	if len(errors) == 0 {
		return nil
	}

	return errors
}
