package main

import (
	"net/http"

	"pertemuan7/handler"
	"pertemuan7/middleware"
	"pertemuan7/model"
	"pertemuan7/repository"
	"pertemuan7/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	ge := gin.New()

	ge.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,
			map[string]any{
				"status": "OK!",
			})
	})

	dsn := "host=localhost user=postgres password=root dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		panic(err)
	}

	userLocalRepo := &repository.ProductLocalRepo{}
	userPgRepo := &repository.ProductPgRepo{DB: db}
	userService := &service.UserService{UserLocalRepo: userLocalRepo, UserPgRepo: userPgRepo}
	userHandler := &handler.UserHandler{UserService: userService}
	apiV1 := ge.Group("/api/v1")
	apiV1.Use(middleware.Middleware1)

	userGroup := apiV1.Group("/users")
	userGroup.POST("/login", userHandler.Login)

	userGroup.Use(middleware.Middleware2)
	userGroup.Use(middleware.BearerAuthorization())
	userGroup.GET("", userHandler.Get)
	userGroup.POST("", userHandler.Create)

	userGroup.Use(middleware.Middleware3)
	userGroup.PUT("/:id", userHandler.Update) // midware1 midware2 midware3 update
	if err := ge.Run(":8080"); err != nil {
		panic(err)
	}

	productGroup := apiV1.Group("/products")
	productGroup.GET("", userHandler.Get)
	productGroup.POST("", userHandler.Create)

	productGroup.PUT("/:id", userHandler.Update)
	productGroup.DELETE("/:id", userHandler.Delete)
	if err := ge.Run(":8080"); err != nil {
		panic(err)
	}
}
