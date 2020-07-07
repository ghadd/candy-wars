package drawers

import (
	"errors"
	"image"
	"io/ioutil"
	"os"
	"strings"
)

const (
	entryFolder = "photos/"
)

/*
This way of storing images optimizes their loading in runtime as it is
not needed to call os.Open millions of times in the row to open same
images;
*/

var (
	imageCount int
	loaded     = false
	setUpTemp  = false
	images     map[string]image.Image
)

func loadImages() error {
	files, err := ioutil.ReadDir(entryFolder)
	if err != nil {
		return err
	}

	imageCount = len(files)
	images = make(map[string]image.Image, imageCount)

	for _, file := range files {
		fileImage, err := os.Open(entryFolder + file.Name())
		if err != nil {
			return err
		}
		img, ext, err := image.Decode(fileImage)
		if strings.ToLower(ext) != "png" {
			return errors.New("Unacceptable file format \"" + ext + "\": use \"png\" instead.")
		}

		images[file.Name()] = img
	}

	loaded = true
	return nil
}

func setUpTempFolder() error {
	path := "temp"
	mode := os.ModeDir

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, mode)
		if err != nil {
			return err
		}
	}
	return nil
}

func SetUpAll() error {
	err := loadImages()
	if err != nil {
		return err
	}

	err = setUpTempFolder()
	if err != nil {
		return err
	}
	return nil
}
