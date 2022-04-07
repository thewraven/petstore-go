package main

import (
	"petstore/api"
	"petstore/internal/backend"
	"petstore/internal/storage/memory"

	"github.com/gin-gonic/gin"
)

func main() {
	ps := memory.NewPetStorage()
	petService := backend.NewPetService(ps)
	var storeHandler api.StoreHandler
	createRouter(api.NewPetHandler(petService), storeHandler, gin.Default())
}
