package drawers

import (
	"testing"
)

func Test_loadImages(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			"base test",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := loadImages(); (err != nil) != tt.wantErr {
				t.Errorf("loadImages() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
