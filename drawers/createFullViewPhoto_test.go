package drawers

import (
	"github.com/ghadd/candy-wars/api"
	"github.com/ghadd/candy-wars/config"
	"github.com/ghadd/candy-wars/models"
	"math/rand"
	"reflect"
	"testing"
)

func TestCreateFullViewPhoto(t *testing.T) {
	funcs := []interface{}{
		models.NewSign,
		models.NewMonster,
		models.NewBlock,
		models.NewChest,
		models.NewEmptyField,
	}

	locs := make([]models.Location, 81, 81)
	for i := 0; i < config.DefaultFieldDimension; i++ {
		for j := 0; j < config.DefaultFieldDimension; j++ {
			f := reflect.ValueOf(funcs[rand.Intn(5)])
			params := []reflect.Value{reflect.ValueOf(i), reflect.ValueOf(j)}

			locs[j*config.DefaultFieldDimension+i] = f.Call(params)[0].Interface().(models.Location)
		}
	}

	type args struct {
		locations []models.Location
		players   []models.Player
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
				locs,
				[]models.Player{
					*models.NewPlayer(api.User{ID: 123}, 5, 5),
				},
				"full-test",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateFullViewPhoto(tt.args.locations, tt.args.players, tt.args.saveTo); (err != nil) != tt.wantErr {
				t.Errorf("CreateFullViewPhoto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
