package drawers

import (
	"testing"
)

func Test_scale(t *testing.T) {
	tests := []struct {
		name    string
		v       float64
		min1    float64
		min2    float64
		max1    float64
		max2    float64
		want    float64
		wantErr bool
	}{
		{
			"positive", 2, 0, 0, 10, 100, 20, false,
		},
		{
			"negative", 3, 0, 0, 10, 100, 40, true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v2 := scale(test.v, test.min1, test.max1, test.min2, test.max2)
			if (v2 == test.want) == test.wantErr {
				t.Errorf("got value 2 %v, want value 2 %v", v2, test.want)
			}
		})
	}
}

func Test_drawBackground(t *testing.T) {
	tests := []struct {
		name       string
		x, y, w, h int
	}{
		{
			"uselessness lvl: aqua", 0, 0, 1023, 1024,
		},
	}

	var wantimage bool
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.w != 1024 {
				wantimage = false
			} else if test.h != 1024 {
				wantimage = false
			} else {
				wantimage = true
			}
			if wantimage != true {
				t.Errorf("wanted image got successful true, got successful %t", wantimage)
			}
		})
	}

}

func Test_drawGrid(t *testing.T) {
	tests := []struct {
		name      string
		dimension int
	}{
		{
			"test", -3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.dimension < 0 {
				t.Errorf("got value %d negative, want %d positive", test.dimension, test.dimension)
			}
		})
	}
}
