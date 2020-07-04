package drawers

import (
	"github.com/ghadd/candy-wars/api"
	"github.com/ghadd/candy-wars/models"
	"testing"
)

func TestCreatePartViewPhoto(t *testing.T) {
	type args struct {
		locations      []models.Location
		players        []models.Player
		drawingCenterX int
		drawingCenterY int
		drawingHorizon int
		saveTo         string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"1-tile-horizon",
			args{
				locations: []models.Location{
					models.NewChest(2, 3),
					models.NewMonster(5, 7),
				},
				players: []models.Player{
					*models.NewPlayer(api.User{ID: 123}, 3, 3),
				},
				drawingCenterX: 3,
				drawingCenterY: 3,
				drawingHorizon: 1,
				saveTo:         "part-1-tile-horizon",
			},
			false,
		},
		{
			"2-tile-horizon",
			args{
				locations: []models.Location{
					models.NewChest(2, 3),
					models.NewMonster(5, 7),
				},
				players: []models.Player{
					*models.NewPlayer(api.User{ID: 123}, 3, 3),
				},
				drawingCenterX: 3,
				drawingCenterY: 3,
				drawingHorizon: 2,
				saveTo:         "part-2-tile-horizon",
			},
			false,
		},
		{
			"top-left",
			args{
				locations: []models.Location{
					models.NewChest(0, 1),
					models.NewMonster(1, 0),
				},
				players: []models.Player{
					*models.NewPlayer(api.User{ID: 123}, 0, 0),
				},
				drawingCenterX: 0,
				drawingCenterY: 0,
				drawingHorizon: 1,
				saveTo:         "part-top-left",
			},
			false,
		},
		{
			"bottom-right",
			args{
				locations: []models.Location{
					models.NewChest(8, 7),
					models.NewMonster(7, 8),
				},
				players: []models.Player{
					*models.NewPlayer(api.User{ID: 123}, 8, 8),
				},
				drawingCenterX: 8,
				drawingCenterY: 8,
				drawingHorizon: 1,
				saveTo:         "part-bottom-right",
			},
			false,
		},
		{
			"bottom-center",
			args{
				locations: []models.Location{
					models.NewChest(5, 8),
					models.NewMonster(3, 8),
				},
				players: []models.Player{
					*models.NewPlayer(api.User{ID: 123}, 4, 8),
				},
				drawingCenterX: 4,
				drawingCenterY: 8,
				drawingHorizon: 1,
				saveTo:         "part-bottom-center",
			},
			false,
		},
		{
			"SUPER-HORIZON",
			args{
				locations: []models.Location{
					models.NewChest(8, 6),
					models.NewMonster(1, 8),
				},
				players: []models.Player{
					*models.NewPlayer(api.User{ID: 123}, 5, 5),
				},
				drawingCenterX: 5,
				drawingCenterY: 5,
				drawingHorizon: 10,
				saveTo:         "part-super-horizon",
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreatePartViewPhoto(tt.args.locations, tt.args.players, tt.args.drawingCenterX, tt.args.drawingCenterY, tt.args.drawingHorizon, tt.args.saveTo); (err != nil) != tt.wantErr {
				t.Errorf("CreatePartViewPhoto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
