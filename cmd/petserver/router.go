package main

import (
	"petstore/api"

	"github.com/gin-gonic/gin"
)

func createRouter(p api.PetHandler, s api.StoreHandler, r *gin.Engine) *gin.Engine {
	if r == nil {
		r = gin.Default()
	}
	petRoutes := r.Group("/pet")
	{
		petRoutes.POST("", gin.WrapF(p.CreatePet))
		petRoutes.GET("/:id", gin.WrapF(p.GetPet))
	}
	storeRoutes := r.Group("/store")
	{
		storeRoutes.POST("/order", s.CreateOrder)
	}
	return r
}
