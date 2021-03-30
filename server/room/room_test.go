package room

import (
	"testing"

	"github.com/car12o/grpc-chat/proto"
)

func TestNewUser(t *testing.T) {
	room := NewRoom()
	if len(room.users) != 0 {
		t.Errorf("room.users length should be 0 but got %d", len(room.users))
	}
	if len(room.tokens) != 0 {
		t.Errorf("room.tokens length should be 0 but got %d", len(room.tokens))
	}
}

func TestAddUser(t *testing.T) {
	room := NewRoom()
	username := "John"
	user, err := room.AddUser(username)
	if err != nil {
		t.Errorf("error should be nil but got %s", err)
	}
	if room.users[user.Name] != user {
		t.Errorf("user should had been added to room users")
	}
	if room.tokens[user.Token] != user {
		t.Errorf("user should had been added to room tokens")
	}

	if _, err := room.AddUser(username); err == nil {
		t.Errorf("room should not allow duplicated usernames")
	}
}

func TestBroadcast(t *testing.T) {
	room := NewRoom()
	user, _ := room.AddUser("John")
	msg := "new message"
	room.SendMessage(user.Token, &proto.Message{Msg: "new message"})
	message := <-user.Channel
	if message.User.Name != string(user.Name) {
		t.Errorf("channel message user name `%s` should be equal to user name `%s`", message.User, user.Name)
	}
	if message.Msg != msg {
		t.Errorf("channel message `%s` should be equal to message %s", message.Msg, msg)
	}
}
