package api

import (
	"encoding/json"
	"net/http"
	"petstore"
	"petstore/internal"
)

type PetHandler struct {
	srv internal.PetService
}

func NewPetHandler(petsrv internal.PetService) PetHandler {
	return PetHandler{srv: petsrv}
}

func (h PetHandler) CreatePet(w http.ResponseWriter, r *http.Request) {
	var p petstore.Pet
	json.NewDecoder(r.Body).Decode(&p)
	err := h.srv.AddPet(&p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (pr PetHandler) GetPet(w http.ResponseWriter, r *http.Request) {
	id := ""
	p, err := pr.srv.FindPet(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if p == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}
