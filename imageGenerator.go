package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	Width int
	Height int
}

func (img Image)  ColorModel() color.Model {
	return color.RGBAModel;
}

func (img Image)  Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func (img Image) At(x, y int) color.Color {
	v := uint8(x ^ y) // Customizable color generation logic
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
