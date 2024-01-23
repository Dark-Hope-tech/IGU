package models

import "time"

type Message struct {
	MessageID   int32     `json:"messageid"`
	ChannelID   int32     `json:"channelid"`
	ServerID    int32     `json:"serverid"`
	AuthorID    int32     `json:"authorid"`
	Content     string    `json:"content"`
	DateCreated time.Time `json:"datecreated"`
}
