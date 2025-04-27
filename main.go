package main

import (
	"discord-msg-cleaner/pkg/args"
	discordapi "discord-msg-cleaner/pkg/discord-api"
	msgcleaner "discord-msg-cleaner/pkg/msg-cleaner"
	"errors"
	"fmt"
	"os"
	"strings"
)

var a *args.Args

func init() {
	arg, err := args.Parse()
	if err != nil {
		printErrAndExit(err)
	}

	a = arg
}

func main() {
	fmt.Println("Getting user information ...")
	me, err := discordapi.GetMe(a.Autorization)
	if err != nil {
		printErrAndExit(err)
	}
	fmt.Printf("Username: %s#%s\n", me.Username, me.Discriminator)

	discordapi.Init(a.Autorization, me)
	discordApi := discordapi.DiscordApi

	guildID := a.GuildID
	if guildID != "" {
		guildIds := strings.Split(guildID, ",")
		for _, gId := range guildIds {
			guildIsValid := discordApi.GuildIsValid(gId)
			if guildIsValid == false {
				continue
				printErrAndExit(errors.New("Error: guild id is invalid"))
			}
	
			msgcleaner.ClearGuildMessage(gId)
		}
		
	}

	channelID := a.ChannelID
	if channelID != "" {
		channelIsValid := discordApi.ChannelIsValid(channelID)
		if channelIsValid == false {
			printErrAndExit(errors.New("Error: channel id is invalid"))
		}

		msgcleaner.ClearChannelMessages(channelID)
	}

}

func printErrAndExit(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}
