package drawers

import (
	"gopkg.in/fogleman/gg.v1"
	"image/color"
)

const (
	windowConfig  = 1024
	windowConfigF = 1024.

	gridThickness = 3 // px
)

//Scale translates value v1 from one range(min1; max1) to another(min2; max2).
func scale(v1, min1, max1, min2, max2 float64) float64 {
	v2 := min2 + (max2-min2)*((v1-min1)/(max1-min1))
	return v2
}

//DrawBackground fills background with given color.
func drawBackground(context *gg.Context, c color.Color) {
	context.DrawRectangle(0, 0, windowConfigF, windowConfigF)
	context.SetColor(c)
	context.Fill()
}

//DrawGrid draws a (dimension*dimension) grid with given dimensions.
func drawGrid(context *gg.Context, dimension int) {
	for v := 1; v < dimension; v++ {
		context.DrawRectangle((windowConfigF/float64(dimension))*float64(v), 0, gridThickness, windowConfigF)
		context.SetColor(color.Black)
		context.Fill()
		context.DrawRectangle(0, (windowConfigF/float64(dimension))*float64(v), windowConfigF, gridThickness)
		context.SetColor(color.Black)
		context.Fill()
	}
}
