package args

import (
	"errors"
	"flag"
)

type Args struct {
	Autorization string
	ChannelID    string
	GuildID      string
}

func Parse() (*Args, error) {
	args := Args{}
	flag.StringVar(&args.Autorization, "authorization", "", "authorization token for access discord api")
	flag.StringVar(&args.ChannelID, "channel", "", "delete messages from channel")
	flag.StringVar(&args.GuildID, "guild", "", "delete messages from guild")
	flag.Parse()

	if args.Autorization == "" {
		return nil, errors.New("Error: --authorization is required. type --help to see all commands")
	}

	if args.ChannelID == "" && args.GuildID == "" {
		return nil, errors.New("Error: missing --channel or --guild")
	}

	if args.ChannelID != "" && args.GuildID != "" {
		return nil, errors.New("Error: only --channel or --guild on")
	}

	return &args, nil
}
