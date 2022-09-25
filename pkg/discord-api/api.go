package discordapi

import (
	"discord-msg-cleaner/pkg/args"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const DISCORD_HOST = "https://discord.com"

var DiscordApi IDiscordAPI

type IDiscordAPI interface {
	// Discord Guild API
	GuildIsValid() bool
	GetUserMessagesID() []UserMessagesID
	GetUserGuildMessagesID(*GuildQuery) ([]UserMessagesID, error)
	DeleteMessageById(*UserMessagesID) error
	GetTotalUserGuildMessages() (int, error)
	GetDetailGuild() (*DetailGuildResponse, error)
}

type discordAPI struct {
	*args.Args
	DiscordMe *APIDiscordMe
}

func Init(me *APIDiscordMe, args *args.Args) {
	DiscordApi = &discordAPI{
		Args:      args,
		DiscordMe: me,
	}
}

type APIDiscordMe struct {
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

func GetMe(authorization string) (*APIDiscordMe, error) {
	request, err := http.NewRequest("GET", "https://discord.com/api/v9/users/@me", nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("accept", "application/json")
	request.Header.Set("authorization", authorization)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		b, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(fmt.Sprintf("Error: failed to authenticate user. got response %s", string(b)))
	}

	var me APIDiscordMe
	err = json.NewDecoder(response.Body).Decode(&me)
	if err != nil {
		return nil, err
	}

	return &me, nil
}

// asdsa
// asdsa
// asdsaas
// asdas

type Request struct {
	Method string
	Path   string
	Body   io.Reader
	Query  []RequestQuery
}

type RequestQuery struct {
	Key   string
	Value string
}

func (da *discordAPI) sendRequest(p *Request) (*http.Response, error) {
	request, err := http.NewRequest(p.Method, DISCORD_HOST+p.Path, p.Body)
	if err != nil {
		return nil, err
	}

	query := request.URL.Query()
	for _, rq := range p.Query {
		query.Add(rq.Key, rq.Value)
	}

	request.URL.RawQuery = query.Encode()
	request.Header.Set("authorization", da.Args.Autorization)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

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
