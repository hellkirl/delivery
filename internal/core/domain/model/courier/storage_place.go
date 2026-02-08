package courier

import (
	"fmt"

	"github.com/google/uuid"
)

var (
	InvalidTotalVolume          = fmt.Errorf("total volume should be bigger than 0")
	StoragePlaceAlreadyOccupied = fmt.Errorf("storage place is already occupied")
	NotEnoughVolume             = fmt.Errorf("not enough volume in storage place")
	IncorrectOrder              = fmt.Errorf("storage place is occupied with another order")
)

type StoragePlace struct {
	id          uuid.UUID
	name        string
	totalVolume int
	orderID     *uuid.UUID
}

func NewStoragePlace(name string, totalVolume int) (*StoragePlace, error) {
	if totalVolume <= 0 {
		return nil, InvalidTotalVolume
	}
	return &StoragePlace{id: uuid.New(), name: name, totalVolume: totalVolume}, nil
}

func (sp *StoragePlace) ID() uuid.UUID { return sp.id }

func (sp *StoragePlace) Name() string { return sp.name }

func (sp *StoragePlace) TotalVolume() int { return sp.totalVolume }

func (sp *StoragePlace) OrderID() *uuid.UUID { return sp.orderID }

func (sp *StoragePlace) Equals(other *StoragePlace) bool {
	return sp.id == other.id
}

func (sp *StoragePlace) CanStore(volume int) (bool, error) {
	if volume <= 0 {
		return false, InvalidTotalVolume
	}

	return true, nil
}

func (sp *StoragePlace) Store(orderID uuid.UUID, volume int) error {
	if volume > sp.totalVolume {
		return NotEnoughVolume
	}

	if sp.orderID != nil {
		return StoragePlaceAlreadyOccupied
	}

	sp.orderID = &orderID

	return nil
}

func (sp *StoragePlace) Clear(orderID uuid.UUID) error {
	if *sp.OrderID() == orderID {
		sp.orderID = nil
	}
	return IncorrectOrder
}

func (sp *StoragePlace) IsOccupied() bool {
	return sp.orderID != nil
}
