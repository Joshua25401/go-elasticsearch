package model

import (
	"strings"
	"time"
)

// Search Category
type UserSearchType string

const (
	BY_NAME    UserSearchType = "name"
	BY_ADDRESS UserSearchType = "address"
)

type User struct {
	FirstName  string    `json:"first_name"`
	MiddleName string    `json:"middle_name"`
	LastName   string    `json:"last_name"`
	Age        int       `json:"age"`
	DayOfBirth time.Time `json:"dob"`
	Address    string    `json:"address"`
}

func (u *User) FullName() string {
	fullName := strings.Join([]string{u.FirstName, u.MiddleName, u.LastName}, " ")
	return strings.Join(strings.Fields(fullName), " ")
}

func (u *User) Index() string {
	return "user_index"
}
