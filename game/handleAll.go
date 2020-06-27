package game

import (
	"github.com/ghadd/candy-wars/api"
)

func HandleUpdate(client *api.Client, update api.Update) {
	if update.HasMessage() {
		handleUpdateMessage(client, update)
	} else if update.HasCallBackQuery() {
		handleUpdateCallBackQuery(client, update)
	}
}