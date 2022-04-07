package api

import (
	"net/http"
	"petstore"
	"petstore/internal/backend"

	"github.com/gin-gonic/gin"
)

type StoreHandler struct {
	srv backend.StoreService
}

func NewStoreHandler(store backend.StoreService) StoreHandler {
	return StoreHandler{srv: store}
}

func (h StoreHandler) CreateOrder(ctx *gin.Context) {
	var o petstore.Order
	err := ctx.BindJSON(&o)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	err = h.srv.PlaceOrder(&o)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusCreated)
}
