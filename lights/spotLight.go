package lights

import (
	"bufio"

	"github.com/CasualSuperman/raytracer/vec"
)

func init() {
	RegisterLightFormat(11, readSpotLight)
}

type spotLight struct {
	baseLight

	aim   vec.Dir
	theta float64
}

func (l *spotLight) Illuminated(p *vec.Pos) bool {
	start := l.baseLight.center
	centerToHit := start.Offset(*p)
	centerToHit.Unit()

	return l.theta < vec.Dot(&centerToHit.Vec, &l.aim.Vec)
}

func readSpotLight(s *bufio.Scanner) (l Light, err error) {
	l, err = readBaseLight(s)

	if err != nil {
		return
	}

	var sl spotLight
	sl.baseLight = *l.(*baseLight)

	sl.center, err = vec.ReadPos(s)

	return &sl, err
}
