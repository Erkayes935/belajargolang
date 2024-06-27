package handler

import (
	"net/http"
	"pertemuan6/model"
	"pertemuan6/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

func (u *UserHandler) Get(ctx *gin.Context) {
	users, err := u.UserService.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Message: "something went wrong",
			Success: false,
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Message: "users fetched",
		Success: true,
		Data:    users,
	})
}

func (u *UserHandler) Create(ctx *gin.Context) {
	// binding payload
	productCreate := model.ProductCreate{}
	if err := ctx.Bind(&productCreate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Message: "bad request param",
			Success: false,
		})
	}
	// call service
	err := u.UserService.Create(&model.Product{
		Name:  productCreate.Name,
		Price: productCreate.Price,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Message: "something went wrong",
			Success: false,
		})
		return
	}
	// response

	ctx.JSON(http.StatusCreated, model.Response{
		Message: "users created",
		Success: true,
	})
}

func (u *UserHandler) Update(ctx *gin.Context) {
	// bind id from path param
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Message: "bad request param",
			Success: false,
		})
	}
	id, _ := strconv.Atoi(idStr)
	// binding payload
	productUpdate := model.ProductUpdate{}
	if err := ctx.Bind(&productUpdate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Message: "bad request param",
			Success: false,
		})
	}
	// call service
	err := u.UserService.Update(uint64(id), &productUpdate)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Message: "something went wrong",
			Success: false,
		})
		return
	}
	// response

	ctx.JSON(http.StatusCreated, model.Response{
		Message: "users updated",
		Success: true,
	})
}

func (u *UserHandler) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, _ := strconv.Atoi(idStr)
	err := u.UserService.Delete(uint64(id))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Message: "something went wrong",
			Success: false,
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Message: "users deleted",
		Success: true,
		Data:    nil,
	})
}
