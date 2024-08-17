package routing

import (
	"github.com/Adosh74/geek-blog/internal/providers/routes"
	"github.com/gin-gonic/gin"
)

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoute() {
	routes.RegisterRoute(GetRouter())
}
