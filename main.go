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
		fmt.Println(err.Error())
		os.Exit(1)
	}

	a = arg
}

func main() {
	fmt.Println("Getting user information ...")
	me, err := discordapi.GetMe(a.Autorization)
	if err != nil {
		printErrAndExit(err)
	}

	discordapi.Init(me, a)
	discordApi := discordapi.DiscordApi

	if a.GuildID != "" {
		guildIsValid := discordApi.GuildIsValid()
		if guildIsValid == false {
			printErrAndExit(errors.New("Error: guild id is invalid"))
		}

		msgcleaner.ClearGuildMessage()
	}

}

func printErrAndExit(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}
