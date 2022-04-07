package api

import (
	"net/http"
	"petstore/internal/backend"
)

type StoreHandler struct {
	srv backend.StoreService
}

func NewStoreHandler(store backend.StoreService) StoreHandler {
	return StoreHandler{srv: store}
}

func (h StoreHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

}
