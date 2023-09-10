package router

import (
	"github.com/gin-gonic/gin"

	"vualt/api/v1/test"
)

func New() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/test", test.Get)
	}
	return r
}
