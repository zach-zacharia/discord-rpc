package main

import (
	"os"
	"time"

	discordrpc "github.com/rikkuness/discord-rpc"
)

func main() {

	drpc, err := discordrpc.New(os.Getenv("DISCORD_CLIENTID"))
	if err != nil {
		panic(err)
	}
	defer drpc.Socket.Close()

	err = drpc.SetActivity(discordrpc.Activity{
		Details: "Foo",
		State:   "Bar",
		Assets: &discordrpc.Assets{
			SmallImage: "keyart_hero",
			LargeImage: "keyart_hero",
		},
		Timestamps: &discordrpc.Timestamps{
			Start: &discordrpc.Epoch{Time: time.Now()},
		},
	})
	if err != nil {
		panic(err)
	}

	for {
		time.Sleep(time.Second)
	}
}
