package chat

import (
	"context"

	"github.com/car12o/grpc-chat/server/user"
	"google.golang.org/grpc/metadata"
)

func getCtxToken(ctx context.Context) user.Token {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	tokens := md["token"]
	if tokens == nil {
		return ""
	}
	token := tokens[0]
	if token == "" {
		return ""
	}
	return user.Token(token)
}
