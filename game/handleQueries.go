package game

import (
	"github.com/ghadd/candy-wars/api"
	"strings"
)

var (
	startsWith = strings.HasPrefix
)

func handleUpdateCallBackQuery(client *api.Client, update api.Update) {
	handleMainMenuQueries(client, update.CallBackQuery)
	handleChooseGameQueries(client, update.CallBackQuery)
	handleMainGameQueries(client, update.CallBackQuery)
}