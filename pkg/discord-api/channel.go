package discordapi

func (a *discordAPI) ChannelIsValid() bool {
	response, err := a.sendRequest(&Request{
		Method: "GET",
		Path:   "/api/v9/channels/" + a.Args.ChannelID,
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
