package controller

import (
	"jwt-client/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kataras/golog"
)

func ErrorHandler(e error, c *gin.Context) bool {
	if e != nil {
		golog.Error(e)

		r := model.Response{Message: e.Error(), Status: false}
		c.JSON(http.StatusInternalServerError, r)
		return false
	}
	return true
}
