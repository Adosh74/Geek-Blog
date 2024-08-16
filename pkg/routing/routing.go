package routing

import "github.com/gin-gonic/gin"

func init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}
