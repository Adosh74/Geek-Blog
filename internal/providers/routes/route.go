package routes

import (
	homeRoutes "github.com/Adosh74/geek-blog/internal/modules/home/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.Engine) {
	homeRoutes.Routes(router)
}
