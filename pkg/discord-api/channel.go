package discordapi

import (
	"encoding/json"
	"errors"
)

func (a *discordAPI) ChannelIsValid(channelID string) bool {
	response, err := a.sendRequest(&Request{
		Method: "GET",
		Path:   "/api/v9/channels/" + channelID,
		Body:   nil,
	})
	if err != nil {
		return false
	}
	defer response.Body.Close()

	return response.StatusCode == 200
}

func (a *discordAPI) GetChannel(channelID string) (*DetailChannel, error) {
	response, err := a.sendRequest(&Request{
		Method: "GET",
		Path:   "/api/v9/channels/" + channelID,
		Body:   nil,
	})
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("failed to get channel")
	}

	var out DetailChannel
	err = json.NewDecoder(response.Body).Decode(&out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
