package msgcleaner

import (
	discordapi "discord-msg-cleaner/pkg/discord-api"
	"fmt"
	"os"
	"sync"
	"time"
)

func ClearGuildMessage() {
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
	var messageIds []discordapi.Return

	for {
		fmt.Println("Getting request message offset", offset)

		time.Sleep(time.Millisecond * 700)
		asd, err := discordApi.GetUserGuildMessagesID(&discordapi.QueryAsd{
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

func deleteMessages(messageIds []discordapi.Return, wg *sync.WaitGroup) {
	if len(messageIds) == 0 {
		fmt.Println("No message can be deleted right now ...")
		wg.Done()
		return
	}

	discordApi := discordapi.DiscordApi
	var counter int
	for {
		time.Sleep(time.Millisecond * 700)

		if counter == len(messageIds) {
			fmt.Println("Success delete", counter, "messages")
			break
		}

		err := discordApi.DeleteMessageById(&messageIds[counter])
		if err != nil {
			if obj, ok := err.(*discordapi.ErrorTimeout); ok {
				time.Sleep(obj.RetryAfter())
			}

			if obj, ok := err.(*discordapi.ErrorRequest); ok {
				switch obj.Code() {
				case 50083:
					counter += 1
					fmt.Println("This messages is in archived thread. try open thread and re-run command again later")
				}
			}
			continue
		}

		counter += 1
	}
	wg.Done()
}
