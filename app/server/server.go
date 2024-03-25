package server

import (
	"book-simple-api/controller"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init() {
	router = CreateRouter(controller.GetAllRoutes())
}

func CreateRouter(routes controller.Routes) *gin.Engine {
	r := gin.Default()
	for _, route := range routes {
		switch route.MethodType {
		case controller.GET:
			r.GET(route.Path, route.HandlerFunc)
		case controller.POST:
			r.POST(route.Path, route.HandlerFunc)
		case controller.PUT:
			r.PUT(route.Path, route.HandlerFunc)
		case controller.DELETE:
			r.DELETE(route.Path, route.HandlerFunc)
		}
	}
	return r
}

func Listen() {
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
