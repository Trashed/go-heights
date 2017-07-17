package hmapgen

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	baseHeight = 255
)

var (
	normalizedData [][]int

	width, height int

	hmapImage *image.RGBA
	white     color.Color = color.RGBA{255, 255, 255, 255}

	shouldBeGray bool
)

// Image creates the height map image from the height map data in Terrain object.
func Image(t Terrain) {

	width = t.size
	height = width

	// Flag for generating a gray scale image
	shouldBeGray = t.genGrayImage

	// Normalize the data
	normalize(t.Data())
	// Save the height map image.
	saveImage(strings.Join([]string{"hmap-", strconv.Itoa(rand.Intn(10000)), ".png"}, ""))
}

func normalize(data [][]int) {
	log.Println("Normalising data. Data len: ", len(data))

	// Initialize the array for normalized values
	normalizedData = make([][]int, len(data))
	for i := range normalizedData {
		normalizedData[i] = make([]int, len(data))
	}

	min, max := func() (int, int) {
		min, max := 0, 0
		for i := range data {
			for j := range data[i] {
				if max < data[i][j] {
					max = data[i][j]
				}
				if min > data[i][j] {
					min = data[i][j]
				}
			}
		}
		return min, max
	}()

	delta := max - min
	percent := float64(baseHeight) / float64(delta)

	for i := range data {
		for j := range data[i] {
			normalizedData[i][j] = int(float64(data[i][j]-min) * percent)
		}
	}
}

func saveImage(imageName string) {
	var waterLevel = (normalizedData[0][0]+normalizedData[0][width-1]+normalizedData[height-1][0]+normalizedData[height-1][width-1])/4 + 1

	log.Println("Save image: ", imageName)

	hmapImage = image.NewRGBA(image.Rect(0, 0, height, width))
	draw.Draw(hmapImage, hmapImage.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	for i := range normalizedData {
		for j := range normalizedData[i] {

			if !shouldBeGray {
				switch {
				case normalizedData[i][j] < waterLevel:
					hmapImage.Set(i, j, color.RGBA{0, 0, uint8(normalizedData[i][j]), 255})
				case normalizedData[i][j] >= waterLevel && normalizedData[i][j] < 170:
					hmapImage.Set(i, j, color.RGBA{0, 255 - uint8(normalizedData[i][j]), 0, 255})
				case normalizedData[i][j] >= 170 && normalizedData[i][j] < 220:
					hmapImage.Set(i, j, color.RGBA{255 - uint8(normalizedData[i][j]), 220 - uint8(normalizedData[i][j]), 0, 255})
				case normalizedData[i][j] >= 220:
					hmapImage.Set(i, j, color.RGBA{uint8(normalizedData[i][j]), uint8(normalizedData[i][j]), uint8(normalizedData[i][j]), 255})
				}
			} else {
				hmapImage.Set(i, j, color.RGBA{uint8(normalizedData[i][j]), uint8(normalizedData[i][j]), uint8(normalizedData[i][j]), 255})
			}
		}
	}
	w, _ := os.Create(imageName)
	defer w.Close()
	png.Encode(w, hmapImage) //Encode writes the Image m to w in PNG format.
}
