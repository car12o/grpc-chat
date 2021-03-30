package user

import "github.com/car12o/grpc-chat/proto"

type Name string
type Token string

type User struct {
	Name    Name
	Token   Token
	Channel chan *proto.Message
}

func NewUser(name Name) *User {
	return &User{
		Name:    name,
		Token:   tokenGenerator(),
		Channel: make(chan *proto.Message),
	}
}

func (u *User) RenewToken() {
	u.Token = tokenGenerator()
}
