package drawers

import (
	"fmt"
	"github.com/ghadd/candy-wars/config"
	"github.com/ghadd/candy-wars/models"
	"github.com/nfnt/resize"
	"gopkg.in/fogleman/gg.v1"
	"image/color"
	"log"
	"strings"
)

//CreateFullViewPhoto draws a full map with all the locations no matter if they are not visible. Only for admins.
func CreateFullViewPhoto(locations []models.Location, players []models.Player, saveTo string) error {
	context := gg.NewContext(windowConfig, windowConfig)
	drawBackground(context, color.RGBA{R: 219, G: 255, B: 204, A: 255})
	drawGrid(context, config.DefaultFieldDimension)

	objects := locations
	for i := 0; i < len(players); i++ {
		objects = append(objects, &players[i])
	}

	for _, l := range objects {
		locX, locY := l.GetLocation()

		imageName := l.GetSmallPic()[strings.Index(l.GetSmallPic(), "/")+1:]
		if !loaded {
			err := loadImages()
			if err != nil {
				log.Println(err)
			}
		}
		img := images[imageName]

		crop := resize.Resize(uint(windowConfigF)/9, uint(windowConfigF)/9, img, resize.Lanczos3)

		context.DrawImage(crop, int(scale(float64(locX), 0, float64(config.DefaultFieldDimension), 0, windowConfigF)), int(scale(float64(locY), 0, float64(config.DefaultFieldDimension), 0, windowConfigF)))
	}

	return context.SavePNG(fmt.Sprintf("temp/%s.png", saveTo))
}
