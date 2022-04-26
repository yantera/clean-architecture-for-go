package gateways

import (
	"api/server/entities"
)

type GuestRepository struct {
	SQLHandler
}

func (repo *GuestRepository) FindByID(id int) (guest entities.Guest, err error) {
	if err = repo.Find(&guest, id).Error; err != nil {
		return
	}
	return
}

func (repo *GuestRepository) FindAll() (guests entities.Guests, err error) {
	if err = repo.Find(&guests).Error; err != nil {
		return
	}
	return
}
