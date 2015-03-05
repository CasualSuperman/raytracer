package lights

import (
	"bufio"
	"fmt"

	"github.com/CasualSuperman/raytracer/color"
	"github.com/CasualSuperman/raytracer/vec"
)

type Light interface {
	Id() int
	Color() color.Color
	Position() vec.Pos
	Illuminated(*vec.Pos) bool
}

var lightReaders map[int]func(*bufio.Scanner)(Light, error) = make(map[int]func(*bufio.Scanner)(Light, error))

func RegisterLightFormat(typ int, reader func(*bufio.Scanner) (Light, error)) {
	lightReaders[typ] = reader
}

func Read(typ int, s *bufio.Scanner) (Light, error) {
	f, ok := lightReaders[typ]
	if !ok {
		return nil, fmt.Errorf("id %d is not a light", typ)
	}
	return f(s)
}

var id = 0

func getNextId() int {
	next := id
	id++
	return next
}
