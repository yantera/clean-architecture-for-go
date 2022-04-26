package guest

import (
	"api/server/entities"
)

type GuestInteractor struct {
	GuestPort GuestPort
}

func (interactor *GuestInteractor) GuestByID(id int) (guest entities.Guest, err error) {
	guest, err = interactor.GuestPort.FindByID(id)
	return
}

func (interactor *GuestInteractor) Guests() (guests entities.Guests, err error) {
	guests, err = interactor.GuestPort.FindAll()
	return
}
