package http

import (
	"github.com/gin-gonic/gin"
	"github.com/lostyear/gin-scaffold/http/controller"
	thttp "github.com/lostyear/go-toolkits/http"
)

func Start(cfg thttp.Config) {
	thttp.StartHTTPServer(cfg, register, nil)
}

func register(rg *gin.RouterGroup) {
	controller.Controller{}.HandlerRegister(rg.Group("/pods"))
}
