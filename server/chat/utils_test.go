package chat

import (
	"context"
	"testing"

	"github.com/car12o/grpc-chat/server/user"
	"google.golang.org/grpc/metadata"
)

func TestGetCtxToken(t *testing.T) {
	token := getCtxToken(context.Background())
	if token != "" {
		t.Errorf("token should be empty but got %s", token)
	}

	ctx := metadata.NewIncomingContext(
		context.Background(),
		metadata.Pairs(),
	)
	token = getCtxToken(ctx)
	if token != "" {
		t.Errorf("token should be empty but got %s", token)
	}

	tk := user.Token("test-token")
	ctx = metadata.NewIncomingContext(
		context.Background(),
		metadata.Pairs("token", string(tk)),
	)
	token = getCtxToken(ctx)
	if token != tk {
		t.Errorf("token should be %s but got %s", tk, token)
	}
}
