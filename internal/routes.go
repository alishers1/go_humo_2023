package internal

import (
	"tasks/pkg"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	r := gin.Default()

	v1 := r.Group("/calculator")
	{
		v1.POST("/subtract", pkg.Subtract)
		v1.POST("/add", pkg.Add)
		v1.POST("/multiply", pkg.Multiply)
		v1.POST("/divide", pkg.Divide)
	}

	r.Run(":8080")
}
