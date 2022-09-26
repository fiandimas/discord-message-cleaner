package discordapi

import (
	"encoding/json"
	"net/http"
	"time"
)

func (a *discordAPI) DeleteMessageById(r *UserMessagesID) error {
	response, err := a.sendRequest(&Request{
		Method: "DELETE",
		Path:   "/api/v9/channels/" + r.ChannelID + "/messages/" + r.MessageID,
		Body:   nil,
	})
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusTooManyRequests {
		var t RequestTimeout
		_ = json.NewDecoder(response.Body).Decode(&t)

		return &ErrorTimeout{
			retryAfter: time.Second * (time.Duration(t.RetryAfter + 2)),
		}
	} else {
		if response.StatusCode != 204 {
			var re RequestError
			_ = json.NewDecoder(response.Body).Decode(&re)

			return &ErrorRequest{
				code: re.Code,
			}
		}
	}

	return nil
}
