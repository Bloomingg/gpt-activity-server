package route

import (
	"customer/api"
	"customer/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	// gin.Default()和gin.New()的区别：gin.Default()默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	r.Use(middlewares.Cors())
	user := r.Group("/activity")
	{
		user.POST("/create", api.CreateActivity)
		user.POST("/get", api.GetActivityByID)
		user.POST("/user_join", api.UserJoin)
	}
	return r
}
