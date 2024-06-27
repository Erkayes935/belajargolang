package handler

import (
	"PERTEMUAN7/helper"
	"net/http"
	"pertemuan7/model"
	"pertemuan7/service"
	"strconv"
	"time"

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
	hashedPassword, err := helper.HashPassword(productCreate.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Message: "something went wrong",
			Success: false,
		})
		return
	}

	// call service
	err = u.UserService.Create(&model.Product{
		Name:     productCreate.Name,
		Price:    productCreate.Price,
		Email:    productCreate.Email,
		Password: hashedPassword,
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

func (u *UserHandler) Login(ctx *gin.Context) {
	// binding payload
	payload := &model.UserLogin{}
	if err := ctx.Bind(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Message: "bad request param",
			Success: false,
		})
		return
	}
	// fetch user by email
	user, err := u.UserService.GetByEmail(payload.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Message: "something went wrong",
			Success: false,
		})
		return
	}
	if user.ID <= 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, model.Response{
			Message: "user not found",
			Success: false,
		})
		return
	}
	// compare password
	isMatched := helper.CheckPasswordHash(payload.Password, user.Password)
	if !isMatched {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, model.Response{
			Message: "invalid email or password",
			Success: false,
		})
		return
	}
	// generate TOKEN if password is correct
	authToken, _ := helper.GenerateUserJWT(user.Name, user.Email, 2*time.Hour)
	sessionToken, _ := helper.GenerateUserJWT(user.Name, user.Email, 48*time.Hour)
	// return
	ctx.JSON(http.StatusOK, model.Response{
		Message: "logged in",
		Success: true,
		Data: model.Token{
			AuthToken:    authToken,
			SessionToken: sessionToken,
		},
	})
}
