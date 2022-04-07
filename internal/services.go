package internal

import (
	"image"
	"petstore"
)

type PetService interface {
	AddPet(p *petstore.Pet) error
	AddPhoto(petID string, img image.Image, metadata string) error
	UpdatePet(p *petstore.Pet) error
	FindPets(s petstore.Status) ([]petstore.Pet, error)
	FindPet(id string) (*petstore.Pet, error)
	RemovePet(id string) error
}
