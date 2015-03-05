package image

import (
	"bufio"
	"fmt"
	"math"
)

type Color struct {
	R, G, B float64
}

func Read(s *bufio.Scanner) (c Color, err error) {
	end := s.Scan()
	if end == false {
		return c, s.Err()
	}

	var count int

	count, err = fmt.Sscanf(s.Text(), "%f %f %f", &c.R, &c.G, &c.B)

	if err != nil {
		return c, err
	}

	if count != 3 {
		return c, fmt.Errorf("tried to read a color, only got %d values.", count)
	}

	return c, nil
}

/* Scale each component of the color by the given factor. */
func (c *Color) Scale(factor float64) {
	c.R *= factor
	c.G *= factor
	c.B *= factor
}

/* Cap each channel of the color at the given value. */
func (c *Color) Cap(max float64) {
	c.R = math.Min(max, c.R)
	c.G = math.Min(max, c.G)
	c.B = math.Min(max, c.B)
}

/* Return the sum of all channels. Useful for checking if the color is black. */
func (c *Color) Magnitude() float64 {
	return c.R + c.G + c.B
}

/* Return a string representation of the color. */
func (c Color) String() string {
	return fmt.Sprintf("R: %.2f, G: %.2f, B: %.2f", c.R, c.G, c.B)
}
