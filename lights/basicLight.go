package lights

import (
	"bufio"
	"fmt"

	"github.com/CasualSuperman/raytracer/color"
	"github.com/CasualSuperman/raytracer/vec"
)

type baseLight struct {
	id     int
	center vec.Pos
	color  color.Color
}

func init() {
	RegisterLightFormat(10, readBaseLight)
}

func (l *baseLight) Id() int {
	return l.id
}

func (l *baseLight) Illuminated(v *vec.Pos) bool {
	return true
}

func (l *baseLight) Color() color.Color {
	return l.color
}

func (l *baseLight) Position() vec.Pos {
	return l.center
}

func readBaseLight(s *bufio.Scanner) (l Light, err error) {
	var light baseLight
	light.id = getNextId()

	light.color, err = color.Read(s)

	if err != nil {
		return
	}

	light.center, err = vec.ReadPos(s)

	return &light, err
}

func (l baseLight) String() string {
	return fmt.Sprintf("Light:\n\tId: %d\n\tPosition:\n\t%s\n\tIntensity:\n\t%v\n", l.id, l.center.String(), l.color.String())
}
