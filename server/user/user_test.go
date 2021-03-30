package user

import "testing"

func TestNewUser(t *testing.T) {
	var username Name = "John"
	user := NewUser(username)
	if user.Name != username {
		t.Errorf("user.Name should be %s but got %s", username, user.Name)
	}
	if user.Token == "" {
		t.Errorf("user.Token should not be empty")
	}
	if user.Channel == nil {
		t.Errorf("user.Channel should be initialized")
	}
}

func TestRenewToken(t *testing.T) {
	user := NewUser("John")
	token := user.Token
	user.RenewToken()
	if token == user.Token {
		t.Errorf("user.Token %s and Token %s should be diferrent", user.Token, token)
	}
}
