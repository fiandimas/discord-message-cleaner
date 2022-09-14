package main

import (
	"discord-msg-cleaner/pkg/args"
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
}
