package main

import (
	"net/http"
	"petstore/api"
	"petstore/internal/backend"
	"petstore/internal/storage/memory"
)

func main() {
	ps := memory.NewPetStorage()
	petService := backend.NewPetService(ps)
	var storeHandler api.StoreHandler
	createRouter(api.NewPetHandler(petService), storeHandler, http.DefaultServeMux)
}
