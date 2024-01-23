package models

type Channel struct {
	ChannelID      int32   `json:"channelid"`
	ServerID       int32   `json:"serverid"`
	ParticipantsID []int32 `json:"participantsid"`
	ChannelName    string  `json:"channelname"`
}
