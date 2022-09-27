package msgcleaner

import (
	discordapi "discord-msg-cleaner/pkg/discord-api"
	"fmt"
	"sync"
	"time"
)

func ClearChannelMessages(channelID string) {
	discordApi := discordapi.DiscordApi

	fmt.Println("Getting total messages in channel ...")
	time.Sleep(time.Second * 2)

	channel, err := discordApi.GetChannel(channelID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	totalMsg, err := discordApi.GetTotalMessages(&discordapi.GuildQuery{
		ChannelID: channelID,
		GuildID:   channel.GuildID,
	})
	if err != nil {
		if obj, ok := err.(*discordapi.ErrorTimeout); ok {
			fmt.Println("Timeout", obj.RetryAfter(), ". try again later")
			return
		}

		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Total messages %d ... \n", totalMsg)
	if totalMsg == 0 {
		fmt.Println("Nothing to delete")
		return
	}

	time.Sleep(time.Second * 2)

	fmt.Println("Downloading messages id ...")

	var offset int
	var errCounter int
	var wg sync.WaitGroup
	var messageIds []discordapi.MessageID

	for {
		fmt.Println("Getting request message offset", offset)

		time.Sleep(time.Millisecond * 700)
		asd, err := discordApi.GetMessagesID(&discordapi.GuildQuery{
			Offset:    offset,
			ChannelID: channelID,
			GuildID:   channel.GuildID,
		})

		if err != nil {
			if obj, ok := err.(*discordapi.ErrorTimeout); ok {
				fmt.Println("Timeout", obj.RetryAfter(), "")
				wg.Add(1)
				go deleteMessages(messageIds, &wg)
				time.Sleep(obj.RetryAfter())
				messageIds = nil
				continue
			}

			errCounter += 1
			if errCounter > 3 {
				fmt.Printf("Error %dx while getting messages at offset %d", errCounter, offset)
				offset += 25
				errCounter = 0
			}

			continue

		}

		offset += 25
		messageIds = append(messageIds, asd...)

		if len(asd) != 25 {
			break
		}
	}
	wg.Add(1)
	go deleteMessages(messageIds, &wg)
	wg.Wait()

	fmt.Println("Success delete messages in guild")
}
