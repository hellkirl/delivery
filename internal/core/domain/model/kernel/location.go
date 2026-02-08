package kernel

import (
	"errors"
	"math/rand/v2"
)

type Location struct {
	x     int
	y     int
	isSet bool
}

func NewLocation(x, y int) (*Location, error) {
	if x < 1 || x > 10 {
		return nil, errors.New("x must be between 1 and 10")
	}
	if y < 1 || y > 10 {
		return nil, errors.New("y must be between 1 and 10")
	}
	return &Location{x: x, y: y, isSet: true}, nil
}

func (l *Location) X() int {
	return l.x
}

func (l *Location) Y() int {
	return l.y
}

func (l *Location) Equals(loc *Location) bool {
	return l.x == loc.x && l.y == loc.y
}

func (l *Location) IsEmpty() bool {
	return l.x >= 1 && l.x <= 10 && l.y >= 1 && l.y <= 10
}

func (l *Location) DistanceTo(other *Location) (int, error) {
	if l.IsEmpty() || other.IsEmpty() {
		return 0, errors.New("invalid location")
	}

	dx := l.x - other.x
	dy := l.y - other.y
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	return dx + dy, nil
}

func NewRandomLocation() *Location {
	randCoordinate := func() int {
		min, max := 1, 10
		return min + rand.IntN(max-min+1)
	}
	x, y := randCoordinate(), randCoordinate()
	loc, _ := NewLocation(x, y)
	return loc
}
