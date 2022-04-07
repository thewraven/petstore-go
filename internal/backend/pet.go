package backend

import (
	"image"
	"petstore"
	"petstore/internal/storage"
)

type PetService struct {
	st storage.PetStorage
}

func (ps PetService) AddPet(p *petstore.Pet) error {
	return nil
}
func (ps PetService) AddPhoto(petID string, img image.Image, metadata string) error {
	return nil
}
func (ps PetService) UpdatePet(p *petstore.Pet) error {
	return nil
}
func (ps PetService) FindPets(s petstore.Status) ([]petstore.Pet, error) {
	return nil, nil
}
func (ps PetService) FindPet(id string) (*petstore.Pet, error) {
	return nil, nil
}

func (ps PetService) RemovePet(id string) error {
	return nil
}

func NewPetService(st storage.PetStorage) PetService {
	return PetService{st: st}
}
