package drawers

import (
	"fmt"
	"github.com/ghadd/candy-wars/config"
	"github.com/ghadd/candy-wars/models"
	"github.com/nfnt/resize"
	"gopkg.in/fogleman/gg.v1"
	"image"
	"image/color"
	"log"
	"os"
)

const (
	forestPath = "photos/forest.png"
)

var (
	forestLoaded = false
	forestImage  image.Image
)

//CreateMapViewPhoto draws a full map but only areas that have been visited will be displayed.
func CreateMapViewPhoto(locations []models.Location, players []models.Player, visited [][]bool, saveTo string) error {
	context := gg.NewContext(windowConfig, windowConfig)
	drawBackground(context, color.RGBA{R: 219, G: 255, B: 204, A: 255})
	drawGrid(context, config.DefaultFieldDimension)
	objects := locations
	for i := 0; i < len(players); i++ {
		objects = append(objects, &players[i])
	}

	for _, l := range objects {
		locX, locY := l.GetLocation()
		f, err := os.Open(l.GetSmallPic())
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			log.Fatal(err)
		}

		crop := resize.Resize(uint(windowConfigF)/9, uint(windowConfigF)/9, img, resize.Lanczos3)

		if visited[locX][locY] == true {
			context.DrawImage(crop, int(scale(float64(locX), 0, float64(config.DefaultFieldDimension), 0, windowConfigF)), int(scale(float64(locY), 0, float64(config.DefaultFieldDimension), 0, windowConfigF)))
		} else {
			drawForest(context, locX, locY)
		}
	}
	return context.SavePNG(fmt.Sprintf("temp/%s.png", saveTo))
}

func drawForest(context *gg.Context, locX int, locY int) {
	if !forestLoaded {
		fl, err := os.Open(forestPath)
		if err != nil {
			log.Fatal(err)
		}

		img, _, err := image.Decode(fl)
		if err != nil {
			log.Fatal(err)
		}
		forestImage = resize.Resize(uint(windowConfigF)/uint(config.DefaultFieldDimension), uint(windowConfigF)/uint(config.DefaultFieldDimension), img, resize.Lanczos3)
		forestLoaded = true
	}

	context.DrawImage(forestImage, int(scale(float64(locX), 0, float64(config.DefaultFieldDimension), 0, windowConfigF)), int(scale(float64(locY), 0, float64(config.DefaultFieldDimension), 0, windowConfigF)))
}
