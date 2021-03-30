package user

import (
	"testing"
)

func TestTokenGenerator(t *testing.T) {
	token1 := tokenGenerator()
	if len(token1) != 16 {
		t.Errorf("token(1) lengh should be 16 but got %d", len(token1))
	}

	token2 := tokenGenerator()
	if len(token2) != 16 {
		t.Errorf("token(2) lengh should be 16 but got %d", len(token2))
	}

	if token1 == token2 {
		t.Errorf("token(1) [%s] and Token(2) [%s] should be different", token1, token2)
	}
}
