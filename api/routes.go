package api

import "github.com/gin-gonic/gin"

func Routes(route *gin.RouterGroup) {
	controller := new(Controller)
	route.GET("/min", controller.Min)
	route.GET("/max", controller.Max)
	route.GET("/avg", controller.Average)
	route.GET("/median", controller.Median)
	route.GET("/percentaile", controller.Percentaile)
	route.GET("/version", controller.Version)
}
