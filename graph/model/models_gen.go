// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Message struct {
	ID   string  `json:"id"`
	Text *string `json:"text"`
}

type MessagesConnection struct {
	Edges    []*MessagesEdge `json:"edges"`
	PageInfo *PageInfo       `json:"pageInfo"`
}

type MessagesEdge struct {
	Cursor string   `json:"cursor"`
	Node   *Message `json:"node"`
}

type PageInfo struct {
	StartCursor string `json:"startCursor"`
	EndCursor   string `json:"endCursor"`
	HasNextPage *bool  `json:"hasNextPage"`
}
