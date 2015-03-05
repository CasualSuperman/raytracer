package vec

type Vec struct {
	X, Y, Z float64
}

type Pos struct{ Vec }
type Dir struct{ Vec }

type Ray struct {
	Pos
	Vec
}
