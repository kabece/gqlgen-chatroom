package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/kabece/gqlgen-chatroom/graph/generated"
	"github.com/kabece/gqlgen-chatroom/graph/model"
)

func (r *chatRoomResolver) MessagesConnection(ctx context.Context, obj *model.ChatRoom, first *int, after *string) (*model.MessagesConnection, error) {
	// The cursor is base64 encoded by convention, so we need to decode it first
	var decodedCursor string
	if after != nil {
		b, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return nil, err
		}
		decodedCursor = string(b)
	}

	// Here we could query the DB to get data, e.g. SELECT * FROM messages WHERE chat_room_id = obj.ID AND timestamp < decodedCursor
	// Mocking for now
	edges := make([]*model.MessagesEdge, *first)
	count := 0
	currentPage := false
	// If no cursor present start from the top
	if decodedCursor == "" {
		currentPage = true
	}
	hasNextPage := false

	// Iterating over the mocked messages to find the current page
	// In real world use-case you should fetch only the required part of data from the database
	for i, v := range r.Messages[obj.ID] {
		node := v

		if currentPage && count < *first {
			edges[count] = &model.MessagesEdge{
				Cursor: base64.StdEncoding.EncodeToString([]byte(v.ID)),
				Node:   &node,
			}
			count++
		}

		if v.ID == decodedCursor {
			currentPage = true
		}

		// If there are any elements left after the current page we indicate that in the response
		if count == *first && i < len(r.Messages[obj.ID]) {
			hasNextPage = true
		}
	}

	pageInfo := model.PageInfo{
		StartCursor: base64.StdEncoding.EncodeToString([]byte(edges[0].Node.ID)),
		EndCursor:   base64.StdEncoding.EncodeToString([]byte(edges[count-1].Node.ID)),
		HasNextPage: &hasNextPage,
	}

	fmt.Printf("MessagesConnection | first: %v, pageInfo: %+v \n", *first, pageInfo)

	mc := model.MessagesConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return &mc, nil
}

func (r *queryResolver) ChatRoom(ctx context.Context, id string) (*model.ChatRoom, error) {
	if t, ok := r.ChatRooms[id]; ok {
		return &t, nil
	}
	return nil, errors.New("chat room not found")
}

// ChatRoom returns generated.ChatRoomResolver implementation.
func (r *Resolver) ChatRoom() generated.ChatRoomResolver { return &chatRoomResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type chatRoomResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
