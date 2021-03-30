package room

import (
	"fmt"

	"github.com/car12o/grpc-chat/proto"
	"github.com/car12o/grpc-chat/server/user"
)

type Room struct {
	users  map[user.Name]*user.User
	tokens map[user.Token]*user.User
}

func NewRoom() *Room {
	return &Room{
		users:  make(map[user.Name]*user.User),
		tokens: make(map[user.Token]*user.User),
	}
}

func (r *Room) AddUser(name string) (*user.User, error) {
	user := user.NewUser(user.Name(name))
	if r.users[user.Name] != nil {
		return nil, fmt.Errorf("username %s already exists", user.Name)
	}
	for r.tokens[user.Token] != nil {
		user.RenewToken()
	}

	r.users[user.Name] = user
	r.tokens[user.Token] = user

	return user, nil
}

func (r *Room) GetUser(token user.Token) (*user.User, error) {
	user := r.tokens[token]
	if user == nil {
		return nil, fmt.Errorf("user not found in room")
	}
	return user, nil
}

func (r *Room) RemoveUser(token user.Token) error {
	user, err := r.GetUser(token)
	if err != nil {
		return fmt.Errorf("user not found in room")
	}

	delete(r.users, user.Name)
	delete(r.tokens, user.Token)

	return nil
}

func (r *Room) SendMessage(token user.Token, message *proto.Message) error {
	user, err := r.GetUser(token)
	if err != nil {
		return err
	}
	go r.broadcast(user, message)
	return nil
}

func (r *Room) broadcast(user *user.User, message *proto.Message) {
	for _, u := range r.users {
		message := &proto.Message{
			User: &proto.User{
				Name: string(user.Name),
			},
			Msg: message.Msg,
		}
		u.Channel <- message
	}
}
