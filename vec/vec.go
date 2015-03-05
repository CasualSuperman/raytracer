package vec

import (
	"bufio"
	"fmt"
)

func Dot(v1, v2 *Vec) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func IsZero(num float64) bool {
	return -0.00001 < num && num < 0.00001
}

func (v *Vec) String() string {
	return fmt.Sprintf("<%.4f, %.4f, %.4f>", v.X, v.Y, v.Z)
}

func Read(s *bufio.Scanner) (Vec, error) {
	var v Vec
	end := s.Scan()

	if end == false {
		return v, s.Err()
	}

	count, err := fmt.Sscanf(s.Text(), "%f %f %f", &v.X, &v.Y, &v.Z)

	if err != nil {
		return v, err
	}

	if count != 3 {
		return v, fmt.Errorf("unable to read vec: only got %d values", count)
	}

	return v, nil
}
