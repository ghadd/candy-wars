package database

import (
	"fmt"
	"testing"

	"github.com/ghadd/candy-wars/game_model"

	"github.com/ghadd/candy-wars/api"
	"github.com/ghadd/candy-wars/models"
)

func TestDBHandler(t *testing.T) {
	var tests = []api.User{
		api.User{
			ID:       12345,
			Username: "player1",
			State:    0,
		},
		api.User{
			ID:       54321,
			Username: "player2",
			State:    1,
		},
		api.User{
			ID:       56789,
			Username: "player3",
			State:    2,
		},
		api.User{
			ID:       98765,
			Username: "player4",
			State:    3,
		},
	}

	db, err := NewDBHandler()
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	//Insert test
	for _, tt := range tests {
		testname := fmt.Sprintf("Insert user (%d, %s, %d) into database, if there is not one", tt.ID, tt.Username, tt.State)
		t.Run(testname, func(t *testing.T) {
			if c, err := db.ContainsUser(tt); !c {
				if err == nil {
					err = db.InsertUser(tt)
					if err != nil {
						t.Errorf("Got error: %v", err)
					}
				}
			}
			user, err := db.GetUserByID(tt.ID)
			if err != nil {
				t.Errorf("Got error: %v", err)
			}
			if user.Username != tt.Username || user.State != tt.State {
				t.Errorf("got User (%d, %s, %d), want User (%d, %s, %d)", user.ID, user.Username, user.State, tt.ID, tt.Username, tt.State)
			}
		})
	}
	//User is registered test
	for _, tt := range tests {
		testname := fmt.Sprintf("Check, if user (%d, %s) has been already inserted into database", tt.ID, tt.Username)
		t.Run(testname, func(t *testing.T) {
			flag, err := db.NameExists(tt.Username)
			if !flag || err != nil {
				t.Errorf("got %v, want %v", flag, true)
			}
		})
	}
	//Update test
	for i, tt := range tests {
		newName := fmt.Sprintf("player%d", i)
		testname := fmt.Sprintf("Update user (%d, %s) to user (%d, %s)", tt.ID, tt.Username, tt.ID, newName)
		t.Run(testname, func(t *testing.T) {
			err := db.Update("users", "nickname", newName, "telegram_id", tt.ID)
			if err != nil {
				t.Errorf("Got error: %v", err)
			}
			user, err := db.GetUserByID(tt.ID)
			if user.Username != newName || err != nil {
				t.Errorf("got User (%d, %s, %d), want User (%d, %s, %d)", user.ID, user.Username, user.State, tt.ID, newName, tt.State)
			}
		})
	}

	//Players test
	var players = []models.Player{
		*models.NewPlayer(tests[0], 2, 4),
		*models.NewPlayer(tests[1], 1, 7),
		*models.NewPlayer(tests[2], 0, 4),
		*models.NewPlayer(tests[3], 5, 5),
	}

	for i, tt := range players {
		testname := fmt.Sprintf("Insert new player from user (%d, %s)", tests[i].ID, tests[i].Username)
		t.Run(testname, func(t *testing.T) {
			if c, err := db.ContainsPlayer(tt); !c {
				if err == nil {
					err = db.InsertPlayer(tt)
					if err != nil {
						t.Errorf("Got error: %v", err)
					}
				}
			}
			player, err := db.GetPlayerByID(tt.PlayerId)
			if err != nil {
				t.Errorf("Got error: %v", err)
			}
			if player.PlayerId != tests[i].ID || player.ObjectName != tests[i].Username {
				t.Errorf("got Player (%d, %s ...), want User (%d, %s ...)", player.PlayerId, player.ObjectName, tests[i].ID, tests[i].Username)
			}
		})
	}
	//UpdateGame test
	var games = []*game_model.Game{}
	for i := 0; i < len(tests); i++ {
		game, _ := game_model.NewGame(&tests[i])
		game.GameID = i
		games = append(games, game)
	}
	for _, tt := range games {
		t.Run("Inserting games", func(t *testing.T) {
			err := db.InsertGame(*tt)
			if err != nil {
				t.Errorf("Got error: %v", err)
			}
		})
	}
	for i, tt := range games {
		err := db.InsertGame(*tt)
		newUID := i * 11111
		testname := fmt.Sprintf("Update game (%d, %d) to game (%d, %d)", tt.GameID, tt.PlayerID, tt.GameID, newUID)
		t.Run(testname, func(t *testing.T) {
			tt.PlayerID = newUID
			err = db.UpdateGame(*tt)
			if err != nil {
				t.Errorf("Got error: %v", err)
			}
			game, err := db.GetGameByID(tt.GameID)
			if err != nil {
				t.Errorf("Got error: %v", err)
			}
			if tt.GameID != game.GameID || newUID != game.PlayerID {
				t.Errorf("got Game (%d, %d), want Game (%d, %d)", game.GameID, game.PlayerID, tt.GameID, tt.PlayerID)
			}
		})
	}
}
