package courier

import (
	"delivery/internal/core/domain/model/kernel"
	"delivery/internal/core/domain/model/order"
	"errors"
	"math"

	"github.com/google/uuid"
)

type Courier struct {
	id            uuid.UUID
	name          string
	speed         int
	location      kernel.Location
	storagePlaces []*StoragePlace
}

func NewCourier(name string, speed int, location kernel.Location) (*Courier, error) {
	if name == "" {
		return nil, errors.New("courier should have a name")
	}
	if !location.IsValid() {
		return nil, errors.New("location is not valid")
	}

	return &Courier{
		id:            uuid.New(),
		name:          name,
		speed:         speed,
		location:      location,
		storagePlaces: []*StoragePlace{},
	}, nil
}

func RestoreCourier(id uuid.UUID, name string, speed int, location kernel.Location, storagePlaces []*StoragePlace) *Courier {
	return &Courier{
		id:            id,
		name:          name,
		speed:         speed,
		location:      location,
		storagePlaces: storagePlaces,
	}
}

func (c *Courier) ID() uuid.UUID { return c.id }

func (c *Courier) Name() string { return c.name }

func (c *Courier) Speed() int { return c.speed }

func (c *Courier) Location() kernel.Location { return c.location }

func (c *Courier) StoragePlaces() []*StoragePlace { return c.storagePlaces }

func (c *Courier) AddStoragePlace(name string, volume int) error {
	sp, err := NewStoragePlace(name, volume)
	if err != nil {
		return err
	}
	c.storagePlaces = append(c.storagePlaces, sp)
	return nil
}

func (c *Courier) CanTakeOrder(order *order.Order) (bool, error) {
	for _, sp := range c.storagePlaces {
		if can, err := sp.CanStore(order.Volume()); can {
			return can, err
		}
	}
	return false, nil
}

func (c *Courier) TakeOrder(order *order.Order) error {
	for _, sp := range c.storagePlaces {
		if can, err := sp.CanStore(order.Volume()); can {
			return sp.Store(order.ID(), order.Volume())
		} else if err != nil {
			return err
		}
	}
	return errors.New("courier cannot take the order")
}

func (c *Courier) CompleteOrder(order *order.Order) error {
	sp, err := c.findStoragePlaceByOrderID(order.ID())
	if err != nil {
		return err
	}
	if err = sp.Clear(order.ID()); err != nil {
		return err
	}
	return nil
}

func (c *Courier) CalculateTimeToLocation(target kernel.Location) (float64, error) {
	if !target.IsValid() {
		return 0, errors.New("target cannot be empty")
	}
	d, err := c.location.DistanceTo(&target)
	if err != nil {
		return 0, err
	}

	return float64(d) / float64(c.speed), nil
}

func (c *Courier) Move(target kernel.Location) error {
	if !target.IsValid() {
		return errors.New("target cannot be empty")
	}

	dx := float64(target.X() - c.location.X())
	dy := float64(target.Y() - c.location.Y())
	remainingRange := float64(c.speed)

	if math.Abs(dx) > remainingRange {
		dx = math.Copysign(remainingRange, dx)
	}
	remainingRange -= math.Abs(dx)

	if math.Abs(dy) > remainingRange {
		dy = math.Copysign(remainingRange, dy)
	}

	newX := c.location.X() + int(dx)
	newY := c.location.Y() + int(dy)

	newLocation, err := kernel.NewLocation(newX, newY)
	if err != nil {
		return err
	}

	c.location = *newLocation

	return nil
}

func (c *Courier) findStoragePlaceByOrderID(orderID uuid.UUID) (*StoragePlace, error) {
	for _, sp := range c.storagePlaces {
		if sp.OrderID() != nil && *sp.OrderID() == orderID {
			return sp, nil
		}
	}
	return nil, errors.New("no storage place found for this order")
}
