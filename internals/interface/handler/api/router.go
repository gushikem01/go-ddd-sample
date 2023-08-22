package api

import "github.com/gin-gonic/gin"

func NewRouter(
	userHandler UserHandler,
) *gin.Engine {
	e := gin.New()
	e = userHandler.AddRoutes(e)
	return e
}
