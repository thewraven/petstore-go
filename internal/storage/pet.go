package storage

import (
	"petstore"
)

type PetStorage interface {
	Save(p *petstore.Pet) error
	Update(p *petstore.Pet) error
	ByStatus(s petstore.Status) ([]*petstore.Pet, error)
	ById(id string) (*petstore.Pet, error)
}
