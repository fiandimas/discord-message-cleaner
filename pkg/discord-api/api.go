package discordapi

import (
	"discord-msg-cleaner/pkg/args"
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
	ChannelIsValid() bool

	// Discord Message API
	GetMessagesID(*GuildQuery) ([]UserMessageID, error)
	DeleteMessageById(*UserMessageID) error
	GetTotalMessages() (int, error)

	// Discord Guild API
	GuildIsValid() bool
	GetDetailGuild() (*DetailGuild, error)
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

func GetMe(authorization string) (*APIDiscordMe, error) {
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
