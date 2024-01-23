package models

type User struct {
	Username   string                `json:"username"`
	Email      string                `json:"email"`
	UserID     string                `json:"userid"`
	Mentions   []MentionNotification `json:"mentions"`
	DMChannels []Channel             `json:"dmchannels"`
	Servers    []Server              `json:"servers"`
}
