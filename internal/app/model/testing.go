package model

import "testing"

func TestUser(t *testing.T) *User {

	t.Helper()

	return &User{
		Email:    "email@examole.org",
		Password: "examplePassword",
	}

}
