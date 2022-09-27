package discordapi

import (
	"io"
	"time"
)

type SnowFlake string

// asdadas
// asdsaa
// sadsasa
type Me struct {
	ID                string
	Username          string
	Avatar            string
	AvatarDecoration  string
	Discriminator     string
	PublicFlags       int
	Flags             int
	PurchasedFlag     int
	PremiumUsageFlags int
	Banner            string
	BannerColor       string
	AccentColor       string
	Bio               string
	Locale            string
	NFSWAllowed       string
	MFAEnabled        string
	PremiumType       int
	Email             string
	Verified          bool
	Phone             string
}

type APIDiscordUserGuildMessages struct {
	TotalResults int                                    `json:"total_results"`
	Messages     [][]APIDiscordUserGuildMessagesContent `json:"messages"`
}

type APIDiscordUserGuildMessagesContent struct {
	ID        string `json:"id"`
	ChannelID string `json:"channel_id"`
}

type MessageID struct {
	MessageID string
	ChannelID string
}

type GuildQuery struct {
	Offset    int
	ChannelID string
	GuildID   string
}

type ChannelQuery struct {
	Offset    int
	ChannelID string
	GuildID   string
}

type Request struct {
	Method  string
	Path    string
	Body    io.Reader
	Query   []KeyValue
	Headers []KeyValue
}

type KeyValue struct {
	Key   string
	Value string
}

// Asdddd
type RequestTimeout struct {
	Message    string  `json:"message"`
	RetryAfter float64 `json:"retry_after"`
	Global     bool    `json:"global"`
}

type ErrorTimeout struct {
	retryAfter time.Duration
}

func (e *ErrorTimeout) Error() string {
	return ""
}

func (e *ErrorTimeout) RetryAfter() time.Duration {
	return e.retryAfter
}

// asdsa
// asdsadsa
// asdsa

type RequestError struct {
	Code int `json:"code"`
}

type ErrorRequest struct {
	code int
}

func (e *ErrorRequest) Error() string {
	return ""
}

func (e *ErrorRequest) Code() int {
	return e.code
}

type DetailGuild struct {
	ID   string
	Name string
}

// asdsa
type DetailChannel struct {
	ID      string
	GuildID string `json:"guild_id"`
}
