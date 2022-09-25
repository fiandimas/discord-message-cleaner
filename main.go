package main

import (
	"discord-msg-cleaner/pkg/args"
	discordapi "discord-msg-cleaner/pkg/discord-api"
	msgcleaner "discord-msg-cleaner/pkg/msg-cleaner"
	"errors"
	"fmt"
	"os"
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

	discordapi.Init(me, a)
	discordApi := discordapi.DiscordApi

	if a.GuildID != "" {
		guildIsValid := discordApi.GuildIsValid()
		if guildIsValid == false {
			printErrAndExit(errors.New("Error: guild id is invalid"))
		}

		msgcleaner.ClearGuildMessage()
		os.Exit(1)
	}

}

func printErrAndExit(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}
