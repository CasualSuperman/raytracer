package vec

import (
	"bufio"
	"math"
)

func (d *Dir) Scale(factor float64) {
	d.X *= factor
	d.Y *= factor
	d.Z *= factor
}

func (d *Dir) Invert() {
	d.Scale(-1)
}

func (d *Dir) Unit() {
	length := d.Length()

	if length != 1 && length != 0 {
		d.Scale(1/length)
	} else if length == 0 {
		nan := math.NaN()
		*d = Dir{Vec{nan, nan, nan}}
	}
}

func (d *Dir) Length() float64 {
	return math.Sqrt(Dot(&d.Vec, &d.Vec))
}

func Cross(d1, d2 *Dir) Dir {
	return Dir{Vec{
		d1.Y*d2.Z - d1.Z*d2.Y,
		d1.Z*d2.X - d1.X*d2.Z,
		d1.X*d2.Y - d1.Y*d2.X,
	}}
}

func ReadDir(s *bufio.Scanner) (Dir, error) {
	vec, err := Read(s)
	return Dir{vec}, err
}
