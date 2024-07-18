package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/phihc116/go-clean-architecture/presentation/controllers"
	"github.com/phihc116/go-clean-architecture/presentation/middlewares"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.ExceptionHandler())

	v1 := r.Group("/api/v1")
	{
		v1.GET("", controllers.NewOrderController().GetById)
	}

	return r
}
