package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller struct{}

type MethodType int

const (
	GET MethodType = iota
	POST
	PUT
	DELETE
)

type Route struct {
	MethodType  MethodType
	Path        string
	HandlerFunc gin.HandlerFunc
}
type Routes []Route

// GetAllRoutes all controller routes.
func GetAllRoutes() Routes {
	var routes Routes
	routes = append(routes, BookController{}.Routes()...)
	return routes
}
