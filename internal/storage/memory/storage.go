package memory

import (
	"errors"
)

type Storage[T identifiable[K], K comparable] struct {
	content map[K]T
}

func NewStorage[T identifiable[K], K comparable]() Storage[T, K] {
	return Storage[T, K]{make(map[K]T)}
}

func (s Storage[T, K]) Save(value T) error {
	_, found := s.content[value.GetID()]
	if found {
		return errors.New("object with the same id already exists")
	}
	s.content[value.GetID()] = value
	return nil
}

func (s Storage[T, K]) Update(value T) error {
	s.content[value.GetID()] = value
	return nil
}

func (s Storage[T, K]) ById(key K) (t T, err error) {
	v, found := s.content[key]
	if !found {
		return t, errors.New("key not found")
	}
	return v, nil
}

type identifiable[K comparable] interface {
	GetID() K
	SetID(K)
}
