package memory

import "petstore"

type PetStorage struct {
	Storage[*petstore.Pet, string]
}

func NewPetStorage() PetStorage {
	return PetStorage{Storage: NewStorage[*petstore.Pet, string]()}
}

func (p PetStorage) ByStatus(s petstore.Status) ([]*petstore.Pet, error) {
	var matches []*petstore.Pet
	for _, v := range p.content {
		if v.Status == s {
			matches = append(matches, v)
		}
	}
	return matches, nil
}
