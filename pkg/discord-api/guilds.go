package discordapi

import (
	"encoding/json"
	"errors"
)

func (a *discordAPI) GuildIsValid(guildID string) bool {
	response, err := a.sendRequest(&Request{
		Method: "GET",
		Path:   "/api/v9/guilds/" + guildID,
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

// Get guild information
func (a *discordAPI) GetDetailGuild(guildID string) (*DetailGuild, error) {
	response, err := a.sendRequest(&Request{
		Method: "GET",
		Path:   "/api/v9/guilds/" + guildID,
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
