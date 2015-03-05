package view

import (
	"bufio"

	"github.com/CasualSuperman/raytracer/lights"
	"github.com/CasualSuperman/raytracer/shapes"
)

type Model struct {
	Proj   Projection
	Lights []lights.Light
	Shapes []shapes.Shape
}

func Read(args []string, input bufio.Reader) (Model, error) {
	m := Model{}

	m.Proj, err := ReadProjection(args, input)

	if err != nil {
		return m, err
	}

	return m, nil
}
