package game

import (
	"github.com/ghadd/candy-wars/api"
)

func GameUpdate(client *api.Client) {
	StartGames(client)
}