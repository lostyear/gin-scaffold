package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lostyear/gin-scaffold/http/manager"
	bc "github.com/lostyear/go-toolkits/http/controller"
)

type Controller struct {
	bc.BaseController
}

func (ctrl Controller) HandlerRegister(r *gin.RouterGroup) {
	r.GET("records", ctrl.getAll)
	r.POST("records", ctrl.createNew)
	r.PUT("records/:id", ctrl.update)
	r.DELETE("records/:id", ctrl.delete)
}

func (ctrl *Controller) getAll(c *gin.Context) {
	list := (&manager.Manager{}).GetAll()
	c.JSON(http.StatusOK, list)
}

func (ctrl *Controller) createNew(c *gin.Context) {
	var data manager.Manager
	if err := c.BindJSON(&data); err != nil {
		ctrl.BadRequest(c, "can not convert req data")
		return
	}
	result := data.CreateNew()
	c.JSON(http.StatusOK, result)
}

func (ctrl *Controller) update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.BadRequest(c, "id not support")
	}
	var data manager.Manager
	if err := c.BindJSON(&data); err != nil {
		ctrl.BadRequest(c, "can not convert req data")
		return
	}
	result := data.Update(id)
	c.JSON(http.StatusOK, result)
}

func (ctrl *Controller) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.BadRequest(c, "id not support")
	}
	var data manager.Manager
	if err := c.BindJSON(&data); err != nil {
		ctrl.BadRequest(c, "can not convert req data")
		return
	}
	result := data.Delete(id)
	c.JSON(http.StatusOK, result)
}
