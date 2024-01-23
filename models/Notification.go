package models

type MentionNotification struct {
	ServerID        int32  `json:"serverid"`
	ChannelID       int32  `json:"channelid"`
	MentionedById   int32  `json:"mentionedbyid"`
	MentionedByName string `json:"mentionedbyname"`
	Message         string `json:"message"`
}
