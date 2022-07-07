package main

import (
	"fmt"
	"net/http"

	"go-code/api"
	"go-code/utility"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	ans := utility.CompareVersion("1.10", "1.2.3")
	fmt.Println(ans)
	router := gin.New()

	//Default logger middleware by gin to log
	router.Use(gin.Logger())

	//Default panic recovery middleware by gin and writes 500 if there is a panic
	//It will avoid giving panic error messages
	router.Use(gin.Recovery())

	// setup cors config
	corsConfig := cors.Options{
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		MaxAge:           3 * 60 * 60,
	}

	router.Use(cors.New(corsConfig))

	router.GET("/healthCheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "App is running successfully !!!"})
	})

	app := router.Group("/app")
	{
		api.Routes(app)
	}
	router.Run()

}
