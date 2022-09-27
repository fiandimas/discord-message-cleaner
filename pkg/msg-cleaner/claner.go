package msgcleaner

import (
	discordapi "discord-msg-cleaner/pkg/discord-api"
	"fmt"
	"sync"
	"time"
)

func deleteMessages(messageIds []discordapi.MessageID, wg *sync.WaitGroup) {
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
					// fmt.Println("This messages is in archived thread. try open thread and re-run command again later")
				}
			}
			continue
		}

		counter += 1
	}
	wg.Done()
}
