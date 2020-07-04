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

//CreatePartViewPhoto draws a part-view map where objects are displayed on players horizon.
func CreatePartViewPhoto(locations []models.Location, players []models.Player, drawingCenterX, drawingCenterY, drawingHorizon int, saveTo string) error {
	topLeftX, topLeftY := drawingCenterX-drawingHorizon, drawingCenterY-drawingHorizon

	bgColor := color.RGBA{R: 219, G: 255, B: 204, A: 255}

	context := gg.NewContext(windowConfig, windowConfig)
	drawBackground(context, bgColor)
	horizon := 2*drawingHorizon + 1

	drawGrid(context, horizon)

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

		x := scale(float64(locX-topLeftX), 0, float64(horizon), 0, windowConfig) // windowConfig * (locX - topLeftX) / horizon
		y := scale(float64(locY-topLeftY), 0, float64(horizon), 0, windowConfig) // windowConfig * (locY - topLeftY) / horizon

		crop := resize.Resize(uint(windowConfig/horizon), uint(windowConfigF/horizon), img, resize.Lanczos3)

		if locX >= (drawingCenterX-drawingHorizon) && locX <= (drawingCenterX+drawingHorizon) &&
			locY >= (drawingCenterY-drawingHorizon) && locY <= (drawingCenterY+drawingHorizon) {
			context.DrawImage(crop, int(x), int(y))
		}
	}

	fillBoundaries(context, drawingCenterX, drawingCenterY, drawingHorizon)

	return context.SavePNG(fmt.Sprintf("temp/%s.png", saveTo))
}

func fillBoundaries(context *gg.Context, x int, y int, horizon int) {
	// todo load image to draw

	for i := x - horizon; i <= x+horizon; i++ {
		for j := y - horizon; j <= y+horizon; j++ {
			if i < 0 || j < 0 || i >= config.DefaultFieldDimension || j >= config.DefaultFieldDimension {
				topLeftX := windowConfig * (i - x + horizon) / (horizon*2 + 1)
				topLeftY := windowConfig * (j - y + horizon) / (horizon*2 + 1)
				w := windowConfig / (horizon*2 + 1)

				context.DrawRectangle(float64(topLeftX), float64(topLeftY), float64(w), float64(w))
				context.SetColor(color.Black)
				context.Fill()
			}
		}
	}
}
