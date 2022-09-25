package msgcleaner

import (
	discordapi "discord-msg-cleaner/pkg/discord-api"
	"fmt"
	"os"
	"sync"
	"time"
)

func ClearChannelMessages() {
	discordApi := discordapi.DiscordApi

	fmt.Println("Getting total messages in guilds ...")
	time.Sleep(time.Second * 2)

	asd, err := discordApi.GetTotalUserGuildMessages()
	if err != nil {
		if obj, ok := err.(*discordapi.ErrorTimeout); ok {
			fmt.Println("Timeout ", obj.RetryAfter(), ". try again later")
			os.Exit(1)
		}

		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Total messages ", asd, " ...")
	time.Sleep(time.Second * 2)

	fmt.Println("Downloading messages id ...")

	var offset int
	var errCounter int
	var wg sync.WaitGroup
	var messageIds []discordapi.UserMessagesID

	for {
		fmt.Println("Getting request message offset", offset)

		time.Sleep(time.Millisecond * 700)
		asd, err := discordApi.GetUserGuildMessagesID(&discordapi.GuildQuery{
			Offset: offset,
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
