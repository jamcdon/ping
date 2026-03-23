package Routes

import (
	"github.com/jamcdon/ping/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/daily/:symbol/:days", Controllers.Daily)
	}

	return router
}