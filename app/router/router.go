package router

import (
	"github.com/gin-gonic/gin"
	"github.com/x554462/demo/app/api"
	"github.com/x554462/demo/middleware/mango"
)

func InitRouter(r *gin.Engine) {
	r.Use(mango.New())
	r.GET("/test", api.GetTest)
}
