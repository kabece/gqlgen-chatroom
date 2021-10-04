package graph

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/kabece/gqlgen-chatroom/graph/generated"
	"github.com/kabece/gqlgen-chatroom/graph/model"
)

type Resolver struct {
	ChatRooms  map[string]model.ChatRoom
	Messages map[string][]model.Message
}

func NewResolver() generated.Config {
	const nChatRooms = 20
	const nMessagesPerChatRoom = 100
	r := Resolver{}
	r.ChatRooms = make(map[string]model.ChatRoom, nChatRooms)
	r.Messages = make(map[string][]model.Message, nChatRooms)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < nChatRooms; i++ {
		id := strconv.Itoa(i + 1)
		mockChatRoom := model.ChatRoom{
			ID: id,
			Name: fmt.Sprintf("ChatRoom %s", id),
		}
		r.ChatRooms[id] = mockChatRoom
		r.Messages[id] = make([]model.Message, nMessagesPerChatRoom)

		// Generate messages for the ChatRoom
		for k := 0; k < nMessagesPerChatRoom; k++ {
			text := fmt.Sprintf("Message %d", k + 1)

			mockMessage := model.Message{
				ID: strconv.Itoa(k + 1),
				Text: &text,
			}

			r.Messages[id][k] = mockMessage
		}
	}

	return generated.Config{
		Resolvers: &r,
	}
}
