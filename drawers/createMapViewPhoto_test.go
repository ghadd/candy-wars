package drawers

import (
	"github.com/ghadd/candy-wars/api"
	"github.com/ghadd/candy-wars/models"
	"testing"
)

func TestCreateMapViewPhoto(t *testing.T) {
	vis := make([][]bool, 9, 9)
	for i := range vis {
		vis[i] = make([]bool, 9, 9)
	}
	vis[3][4] = true
	vis[5][5] = true

	type args struct {
		locations []models.Location
		players   []models.Player
		visited   [][]bool
		saveTo    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test",
			args{
				[]models.Location{
					models.NewMonster(2, 4),
					models.NewSign(3, 4),
				},
				[]models.Player{
					*models.NewPlayer(api.User{ID: 123}, 5, 5),
				},
				vis,
				"map-test",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateMapViewPhoto(tt.args.locations, tt.args.players, tt.args.visited, tt.args.saveTo); (err != nil) != tt.wantErr {
				t.Errorf("CreateMapViewPhoto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
