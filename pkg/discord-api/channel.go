package discordapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type APIDiscordChannelMessageAuthor struct {
	ID string
}

type APIDiscordChannelMessage struct {
	ID     string
	Author APIDiscordChannelMessageAuthor
}

func (a *discordAPI) GetChannelMessageId(channel string) error {
	response, err := a.sendRequest(&Request{
		Method: "GET",
		Path:   "/api/v9/channels/" + channel + "/messages",
		Body:   nil,
	})
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		b, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return errors.New(fmt.Sprintf("Error: failed to authenticate user. got response %s", string(b)))
	}

	var asd []APIDiscordChannelMessage
	err = json.NewDecoder(response.Body).Decode(&asd)
	if err != nil {
		return err
	}

	var messageIds []string
	// for _, message := range asd {
	// 	if message.Author.ID == a.ID {

	// 		messageIds = append(messageIds, message.ID)
	// 	}
	// }

	fmt.Println(messageIds)

	return nil
}

func (a *discordAPI) DeleteMessageById(r *Return) error {
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
		var t Timeout
		_ = json.NewDecoder(response.Body).Decode(&t)

		return &ErrorTimeout{
			retryAfter: time.Second * (time.Duration(t.RetryAfter + 2)),
		}
	} else {
		if response.StatusCode != 204 {
			return errors.New(fmt.Sprintf("err %d", response.StatusCode))
		}
	}

	return nil
}
