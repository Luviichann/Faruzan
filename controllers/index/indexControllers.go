package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (con IndexController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
