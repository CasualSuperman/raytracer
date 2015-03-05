package image

import (
	"fmt"
	"io"
)

/* Used to create an image. Has a height and a width. */
type Image struct {
	base          []Pixel
	height, width int
}

/* A type that can point into the Image's internal byte array. */
type Pixel struct {
	R, G, B uint8
}

/* Create a new image with the given dimensions. */
func NewImage(x, y int) *Image {
	i := new(Image)
	i.base = make([]Pixel, x*y)
	i.height = y
	i.width = x
	return i
}

/* Write the image to the given io.Write in PPM format. */
func (i Image) PPM(w io.Writer) {
	fmt.Fprintln(w, "P6")
	fmt.Fprintf(w, "%d %d\n", i.width, i.height)
	fmt.Fprintln(w, "255")
	for _, pixel := range i.base {
		pixel.Write(w)
	}
}

/* Return the image's width. */
func (i Image) Width() int {
	return i.width
}

/* Return the image's height. */
func (i Image) Height() int {
	return i.height
}

// Get a pixel with references into the image, modifying this pixel updates the
// image.
func (i Image) GetPixel(x, y int) *Pixel {
	index := ((y * i.width) + x)
	return &i.base[index]
}

func (i Image) Iter() func() *Pixel {
	ch := make(chan *Pixel)
	go func() {
		for _, p := range i.base {
			ch <- p
		}
		close(ch)
	}()
	return func() *Pixel {
		return <-ch
	}
}

func (p Pixel) Write(w io.Writer) {
	w.Write([]byte{p.R, p.G, p.B})
}
