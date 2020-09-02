package controller

import (
	"jwt-client/model"
	"jwt-client/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var r = model.Response{Message: "success", Status: true}
	var e error

	r.Data, e = service.GetUser()
	if ok := ErrorHandler(e, c); !ok {
		return
	}

	c.JSON(http.StatusOK, r)
}

func GetByIDUser(c *gin.Context) {
	var r = model.Response{Message: "success", Status: true}
	var e error

	id := c.Param("id")

	r.Data, e = service.GetUserByID(id)
	if ok := ErrorHandler(e, c); !ok {
		return
	}

	c.JSON(http.StatusOK, r)
}

func CreateUser(c *gin.Context) {
	var r = model.Response{Message: "success", Status: true}
	var e error

	var p model.User
	e = c.BindJSON(&p)
	if ok := ErrorHandler(e, c); !ok {
		return
	}

	r.Data, e = service.CreateUser(p)
	if ok := ErrorHandler(e, c); !ok {
		return
	}

	c.JSON(http.StatusOK, r)
}

func UpdateUser(c *gin.Context) {
	var r = model.Response{Message: "success", Status: true}
	var e error

	var p model.User
	e = c.BindJSON(&p)
	if ok := ErrorHandler(e, c); !ok {
		return
	}

	id := c.Param("id")

	r.Data, e = service.UpdateUser(p, id)
	if ok := ErrorHandler(e, c); !ok {
		return
	}

	c.JSON(http.StatusOK, r)
}

func DeleteUser(c *gin.Context) {
	var r = model.Response{Message: "success", Status: true}
	var e error

	id := c.Param("id")

	r.Data, e = service.DeleteUser(id)
	if ok := ErrorHandler(e, c); !ok {
		return
	}

	c.JSON(http.StatusOK, r)
}
