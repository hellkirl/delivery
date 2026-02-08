package kernel

import (
	"errors"
	"math/rand/v2"
)

type Location struct {
	x uint8
	y uint8
}

func NewLocation(x uint8, y uint8) (*Location, error) {
	if x < 1 || x > 10 {
		return nil, errors.New("x must be between 1 and 10")
	}
	if y < 1 || y > 10 {
		return nil, errors.New("y must be between 1 and 10")
	}
	return &Location{x: x, y: y}, nil
}

func (l *Location) X() uint8 {
	return l.x
}

func (l *Location) Y() uint8 {
	return l.y
}

func (l *Location) Equals(loc *Location) bool {
	return l.x == loc.x && l.y == loc.y
}

func (l *Location) CalculateDistance(other *Location) uint8 {
	dx := int(l.x) - int(other.x)
	dy := int(l.y) - int(other.y)
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	return uint8(dx + dy)
}

func (l *Location) Random() *Location {
	randCoordinate := func() int {
		min, max := 1, 10
		return min + rand.IntN(max-min+1)
	}
	x, y := uint8(randCoordinate()), uint8(randCoordinate())
	loc, _ := NewLocation(x, y)
	return loc
}
