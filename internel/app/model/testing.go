package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "User@example.org",
		Password: "password",
	}
}
