package routers

import (
	"faruzan/controllers/index"

	"github.com/gin-gonic/gin"
)

func IndexRoutersInit(r *gin.Engine) {
	indexRouters := r.Group("/")
	{
		indexRouters.GET("/", index.IndexController{}.Index)
	}
}
