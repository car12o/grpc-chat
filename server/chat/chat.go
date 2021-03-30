package chat

import (
	"context"

	"github.com/car12o/grpc-chat/proto"
	"github.com/car12o/grpc-chat/server/room"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChatServer struct {
	proto.UnimplementedChatServer
	room *room.Room
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		room: room.NewRoom(),
	}
}

func (cs *ChatServer) RegisterUser(ctx context.Context, u *proto.User) (*proto.Token, error) {
	user, err := cs.room.AddUser(u.GetName())
	if err != nil {
		return nil, err
	}

	return &proto.Token{
		User: &proto.User{
			Name: string(user.Name),
		},
		Hash: string(user.Token),
	}, nil
}

func (cs *ChatServer) ListenBroadcast(_ *proto.Empty, stream proto.Chat_ListenBroadcastServer) error {
	token := getCtxToken(stream.Context())
	user, err := cs.room.GetUser(token)
	if err != nil {
		return err
	}

	for message := range user.Channel {
		if err := stream.Send(message); err != nil {
			status, ok := status.FromError(err)
			if !ok {
				return err
			}
			if status.Code() == codes.Unavailable {
				cs.room.RemoveUser(token)
			}
			return err
		}
	}

	return nil
}

func (cs *ChatServer) PostMessage(ctx context.Context, message *proto.Message) (*proto.Empty, error) {
	token := getCtxToken(ctx)
	if err := cs.room.SendMessage(token, message); err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}
