package guest

import (
	"api/server/entities"
)

type GuestPort interface {
	FindByID(id int) (entities.Guest, error)
	FindAll() (entities.Guests, error)
}
