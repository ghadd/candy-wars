package game

import (
	"encoding/json"
	"github.com/ghadd/candy-wars/api"
	"github.com/ghadd/candy-wars/database"
	"github.com/ghadd/candy-wars/game_model"
	"log"
)

func StartGames(client *api.Client) {
	// Setting up database handler
	dbh, err := database.NewDBHandler()
	if err != nil {
		log.Println(err)
	}

	// getting all games from database
	games := dbh.GetGames()

	for _, gm := range games {
		// skipping all games which are not in Choosing clan state as they don't interest us in this case
		if gm.State != game_model.StateChoosingClan {
			continue
		}

		// otherwise getting that game' players
		players := gm.Players
		ready := len(players) > 0
		for _, player := range players {
			// if someone didn't choose their clan yet -> game is not ready to be started
			if player.Clan == "" {
				ready = false
			}
		}

		// todo perform timing join

		if ready {
			for _, p := range players {
				_, err = client.SendMessage(api.Message{
					ChatID: p.PlayerId,
					Text:   "Game is starting. Be ready.",
				})
			}

			gm.State = game_model.StateRunning
			err = dbh.Update("games", "state", gm.State, "game_id", gm.GameID)
			if err != nil {
				log.Println(err)
			}

			game_model.LocatePlayers(gm)
			playersJSON, err := json.Marshal(gm.Players)
			if err != nil {
				log.Println(err)
			}

			err = dbh.Update("games", "players_json", playersJSON, "game_id", gm.GameID)
			if err != nil {
				log.Println(err)
			}
			for _, p := range gm.Players {
				err = dbh.Update("players", "x", p.X, "player_id", p.PlayerId)
				if err != nil {
					log.Println(err)
				}
				err = dbh.Update("players", "y", p.Y, "player_id", p.PlayerId)
				if err != nil {
					log.Println(err)
				}
			}

			sendFirstGameMessage(client, gm)
		}
	}
}

func sendFirstGameMessage(client *api.Client, gm *game_model.Game) {
	players := gm.Players
	for _, player := range players {
		SendCurrentPhoto(client, api.User{ID: player.PlayerId})
	}

	SendMoveButtons(client, api.User{ID: gm.PlayerID})
}
