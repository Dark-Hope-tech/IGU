package models

type Server struct {
	ServerID      int32   `json:"serverid"`
	ServerName    string  `json:"servername"`
	CreaterID     int32   `json:"createrid"`
	AdminIDList   []int32 `json:"adminidlist"`
	ChannelIDList []int32 `json:"channelidlist"`
}
