package discordapi

import (
	"encoding/json"
	"errors"
)

func (a *discordAPI) GuildIsValid() bool {
	response, err := a.sendRequest(&Request{
		Method: "GET",
		Path:   "/api/v9/guilds/" + a.Args.GuildID,
		Body:   nil,
	})
	if err != nil {
		return false
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return false
	}

	return true
}

// asdsasa
// asdsa
// asdsaas

func (a *discordAPI) GetDetailGuild() (*DetailGuild, error) {
	response, err := a.sendRequest(&Request{
		Method: "GET",
		Path:   "/api/v9/guilds/" + a.Args.GuildID,
		Body:   nil,
	})
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("ERR")
	}

	var out DetailGuild
	err = json.NewDecoder(response.Body).Decode(&out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
