package game

import (
	"github.com/ghadd/candy-wars/api"
)

// Update is being called to refresh every game status and check other events.
func Update(client *api.Client) {
	// Checks if there are not yet started games with all players with chosen clans
	// and starts those games
	StartGames(client)
}
