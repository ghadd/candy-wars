package main

import (
	"github.com/ghadd/candy-wars/api"
	"github.com/ghadd/candy-wars/config"
	"github.com/ghadd/candy-wars/game"
	"log"
	"time"
)

var (
	client *api.Client
)

func main() {
	// Setting up telegram bot client
	var err error

	log.Println("Connecting to bot api.")
	client, err = api.NewClient(config.BotToken)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successful.")

	firstUpdate := 0
	lastUpdate := 0
	var update api.Update

	for {
		updates := client.GetUpdates(lastUpdate + 1)
		if len(updates) != 0 {
			update = updates[0]

			if firstUpdate == 0 {
				firstUpdate = updates[0].UpdateID
			}
			lastUpdate = updates[0].UpdateID

			if update.UpdateID != 0 {
				log.Println("Handling update: ", update)
				// run handlers asynchronously
				go game.HandleUpdate(client, update)
			}
		}

		game.Update(client)

		time.Sleep(time.Millisecond * 100)
	}

}
