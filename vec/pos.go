package vec

import (
	"bufio"
)

func Origin() Pos {
	return Pos{Vec{0, 0, 0}}
}

func (p *Pos) Displace(d Dir) {
	p.X += d.X
	p.Y += d.Y
	p.Z += d.Z
}

func (p *Pos) Offset(target Pos) Dir {
	return Dir{Vec{
		target.X - p.X,
		target.Y - p.Y,
		target.Z - p.Z,
	}}
}

func (p Pos) Copy() Pos {
	return p
}

func ReadPos(s *bufio.Scanner) (Pos, error) {
	vec, err := Read(s)
	return Pos{vec}, err
}
