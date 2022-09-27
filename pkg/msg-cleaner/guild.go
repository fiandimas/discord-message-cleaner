package msgcleaner

import (
	discordapi "discord-msg-cleaner/pkg/discord-api"
	"fmt"
	"os"
	"sync"
	"time"
)

func ClearGuildMessage(guildID string) {
	discordApi := discordapi.DiscordApi

	g, err := discordApi.GetDetailGuild(guildID)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Geting your total messages in %s \n", g.Name)
	time.Sleep(time.Second * 2)

	totalMsg, err := discordApi.GetTotalMessages(&discordapi.GuildQuery{})
	if err != nil {
		if obj, ok := err.(*discordapi.ErrorTimeout); ok {
			fmt.Printf("Timeout %d. try again later \n", obj.RetryAfter())
			os.Exit(1)
		}

		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Total messages %d \n", totalMsg)
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
	fmt.Printf("Delete %d messages ... \n", len(messageIds))
	go deleteMessages(messageIds, &wg)
	wg.Wait()

	fmt.Printf("Success delete your messages in %s \n", g.Name)
}

