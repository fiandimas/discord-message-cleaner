package discordapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type APIDiscordUserGuildMessages struct {
	TotalResults int                                    `json:"total_results"`
	Messages     [][]APIDiscordUserGuildMessagesContent `json:"messages"`
}

type APIDiscordUserGuildMessagesContent struct {
	ID        string `json:"id"`
	ChannelID string `json:"channel_id"`
}

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

func (a *discordAPI) GetUserGuildMessages() (*APIDiscordUserGuildMessages, error) {
	response, err := a.sendRequest(&Request{
		Method: "GET",
		Path:   "/api/v9/guilds/" + a.Args.GuildID + "/messages/search",
		Body:   nil,
		Query: []RequestQuery{
			{
				Key:   "author_id",
				Value: a.DiscordMe.ID,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("Error: failed to get guild information by given guild params")
	}

	var out APIDiscordUserGuildMessages
	err = json.NewDecoder(response.Body).Decode(&out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

// asdas
// sada
// asdsadsadas
// asdsadas

type Return struct {
	MessageID string
	ChannelID string
}

func (a *discordAPI) GetUserMessagesID() []Return {
	var offset int
	var asda []Return

	for {
		time.Sleep(time.Millisecond * 700)
		fmt.Println("Requesting offset ", offset)
		response, err := a.sendRequest(&Request{
			Method: "GET",
			Path:   "/api/v9/guilds/" + a.Args.GuildID + "/messages/search",
			Body:   nil,
			Query: []RequestQuery{
				{
					Key:   "author_id",
					Value: a.DiscordMe.ID,
				},
				{
					Key:   "offset",
					Value: strconv.Itoa(offset),
				},
			},
		})
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		defer response.Body.Close()

		if response.StatusCode == http.StatusTooManyRequests {
			type Timeout struct {
				Message    string  `json:"message"`
				RetryAfter float64 `json:"retry_after"`
				Global     bool    `json:"global"`
			}

			var timeout Timeout
			err = json.NewDecoder(response.Body).Decode(&timeout)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			retryAfter := time.Second * (time.Duration(timeout.RetryAfter + 2))
			fmt.Println("Timeout ", retryAfter, " ...")
			offset -= 25
			time.Sleep(retryAfter)

			continue
		} else {
			if response.StatusCode != 200 {
				continue
			}

			var out APIDiscordUserGuildMessages
			err = json.NewDecoder(response.Body).Decode(&out)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			for _, message := range out.Messages {
				for _, m := range message {
					asda = append(asda, Return{
						MessageID: m.ID,
						ChannelID: m.ChannelID,
					})
				}
			}

			if len(out.Messages) != 25 {
				return asda
			}

			fmt.Println("success")
		}

		offset += 25
	}
}

type QueryAsd struct {
	Offset int
}

func (a *discordAPI) GetUserGuildMessagesID(q *QueryAsd) ([]Return, error) {
	var asda []Return

	response, err := a.sendRequest(&Request{
		Method: "GET",
		Path:   "/api/v9/guilds/" + a.Args.GuildID + "/messages/search",
		Body:   nil,
		Query: []RequestQuery{
			{
				Key:   "author_id",
				Value: a.DiscordMe.ID,
			},
			{
				Key:   "offset",
				Value: strconv.Itoa(q.Offset),
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
			asda = append(asda, Return{
				MessageID: m.ID,
				ChannelID: m.ChannelID,
			})
		}
	}

	return asda, nil
}

func (a *discordAPI) GetTotalUserGuildMessages() (int, error) {
	response, err := a.sendRequest(&Request{
		Method: "GET",
		Path:   "/api/v9/guilds/" + a.Args.GuildID + "/messages/search",
		Body:   nil,
		Query: []RequestQuery{
			{
				Key:   "author_id",
				Value: a.DiscordMe.ID,
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
	json.NewDecoder(response.Body).Decode(&t)
	return t.TotalResults, nil
}
