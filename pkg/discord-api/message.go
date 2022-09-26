package discordapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// asdas
// sada
// asdsadsadas
// asdsadas

func (a *discordAPI) GetMessagesID(q *GuildQuery) ([]UserMessageID, error) {
	var asda []UserMessageID

	response, err := a.sendRequest(&Request{
		Method: "GET",
		Path:   "/api/v9/guilds/" + a.Args.GuildID + "/messages/search",
		Body:   nil,
		Query: []KeyValue{
			{
				Key:   "author_id",
				Value: a.DiscordMe.ID,
			},
			{
				Key:   "offset",
				Value: strconv.Itoa(q.Offset),
			},
			{
				Key:   "include_nsfw",
				Value: "true",
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusTooManyRequests {
		var t RequestTimeout
		_ = json.NewDecoder(response.Body).Decode(&t)
		return nil, &ErrorTimeout{
			retryAfter: time.Second * (time.Duration(t.RetryAfter + 2)),
		}
	}

	var out APIDiscordUserGuildMessages
	err = json.NewDecoder(response.Body).Decode(&out)
	if err != nil {
		return nil, err
	}

	for _, message := range out.Messages {
		for _, m := range message {
			asda = append(asda, UserMessageID{
				MessageID: m.ID,
				ChannelID: m.ChannelID,
			})
		}
	}

	return asda, nil
}

func (a *discordAPI) GetTotalMessages() (int, error) {
	response, err := a.sendRequest(&Request{
		Method: "GET",
		Path:   "/api/v9/guilds/" + a.Args.GuildID + "/messages/search",
		Body:   nil,
		Query: []KeyValue{
			{
				Key:   "author_id",
				Value: a.DiscordMe.ID,
			},
			{
				Key:   "include_nsfw",
				Value: "true",
			},
		},
	})
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusTooManyRequests {
		var t RequestTimeout
		_ = json.NewDecoder(response.Body).Decode(&t)
		return 0, &ErrorTimeout{
			retryAfter: time.Second * (time.Duration(t.RetryAfter + 2)),
		}
	}

	var t APIDiscordUserGuildMessages
	err = json.NewDecoder(response.Body).Decode(&t)
	if err != nil {
		return 0, err
	}

	return t.TotalResults, nil
}

func (a *discordAPI) DeleteMessageById(r *UserMessageID) error {
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
