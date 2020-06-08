package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter init router
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hey man ! It's OK~")
	})
	return router
}
