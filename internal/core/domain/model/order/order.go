package order

import (
	"delivery/internal/core/domain/model/kernel"
	"delivery/internal/pkg/errs"
	"errors"

	"github.com/google/uuid"
)

type Order struct {
	id        uuid.UUID
	courierID *uuid.UUID
	location  kernel.Location
	volume    int
	status    Status
}

func NewOrder(location kernel.Location, volume int) (*Order, error) {
	if !location.IsValid() {
		return nil, errs.NewValueIsRequiredError("location")
	}
	if volume <= 0 {
		return nil, errs.NewValueIsRequiredError("volume")
	}

	return &Order{
		id:       uuid.New(),
		location: location,
		volume:   volume,
		status:   StatusCreated,
	}, nil
}

func RestoreOrder(id uuid.UUID, courierID *uuid.UUID, location kernel.Location, volume int, status Status) *Order {
	return &Order{
		id:        id,
		courierID: courierID,
		location:  location,
		volume:    volume,
		status:    status,
	}
}

func (o *Order) ID() uuid.UUID { return o.id }

func (o *Order) CourierID() *uuid.UUID { return o.courierID }

func (o *Order) Location() kernel.Location { return o.location }

func (o *Order) Volume() int { return o.volume }

func (o *Order) Status() Status { return o.status }

func (o *Order) Assign(couriedID uuid.UUID) error {
	if o.courierID != nil {
		return errors.New("courier already assigned for this order")
	}
	o.courierID = &couriedID
	return nil
}

func (o *Order) Complete() error {
	if o.status != StatusCreated {
		return errors.New("only orders in 'Created' status can be completed")
	}
	o.status = StatusCompleted
	return nil
}
