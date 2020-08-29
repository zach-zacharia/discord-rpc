package discordrpc

import (
	"os"
	"time"
)

func Example() {
	drpc, err := New(os.Getenv("DISCORD_CLIENTID"))
	if err != nil {
		panic(err)
	}
	defer drpc.Socket.Close()

	err = drpc.SetActivity(Activity{
		Details: "Foo",
		State:   "Bar",
		Assets: &Assets{
			SmallImage: "keyart_hero",
			LargeImage: "keyart_hero",
		},
		Timestamps: &Timestamps{
			Start: &Epoch{Time: time.Now()},
		},
	})
	if err != nil {
		panic(err)
	}
}
