package discordapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const DISCORD_HOST = "https://discord.com"

var DiscordApi IDiscordAPI

type IDiscordAPI interface {
	// Discord Channel API
	ChannelIsValid(string) bool
	GetChannel(string) (*DetailChannel, error)

	// Function for
	GetMessagesID(*GuildQuery) ([]MessageID, error)
	DeleteMessageById(*MessageID) error
	GetTotalMessages(*GuildQuery) (int, error)

	// Discord Guild API
	GuildIsValid(string) bool
	GetDetailGuild(string) (*DetailGuild, error)
}

type discordAPI struct {
	*Me
	Authorization string
}

func Init(authorization string, me *Me) {
	DiscordApi = &discordAPI{
		Authorization: authorization,
		Me:            me,
	}
}

// Get user information by given --authorization flag
func GetMe(authorization string) (*Me, error) {
	request, err := http.NewRequest("GET", DISCORD_HOST+"/api/v9/users/@me", nil)
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

		return nil, errors.New(fmt.Sprintf("error: failed to authenticate user. got response %s", string(b)))
	}

	var me Me
	err = json.NewDecoder(response.Body).Decode(&me)
	if err != nil {
		return nil, err
	}

	return &me, nil
}

// Utility for sending http request
func (a *discordAPI) sendRequest(p *Request) (*http.Response, error) {
	request, err := http.NewRequest(p.Method, DISCORD_HOST+p.Path, p.Body)
	if err != nil {
		return nil, err
	}

	query := request.URL.Query()
	for _, rq := range p.Query {
		query.Add(rq.Key, rq.Value)
	}
	
	request.URL.RawQuery = query.Encode()
	request.Header.Set("authorization", a.Authorization)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
