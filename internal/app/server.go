package app

import (
	"sample-server/internal/config"
	"sample-server/internal/domain/usecase"
	"sample-server/internal/gateway/handler"
	"sample-server/internal/gateway/middleware"

	"github.com/labstack/echo/v4"
)

func NewServer(conf *config.Config) *echo.Echo {
	e := echo.New()

	// APIキー認証ミドルウェア
	apiKeyConf := middleware.LoadAPIKeyAuthConfigFromConfig(conf)
	apiKeyAuth := middleware.NewAPIKeyAuthMiddleware(apiKeyConf)
	e.Use(apiKeyAuth.AuthMiddleware)

	helloUsecase := usecase.NewHelloUsecase()
	helloHandler := handler.NewHelloHandler(helloUsecase)
	e.GET("/hello", helloHandler.SayHello)

	userUsecase := usecase.NewUserUsecase()
	userHandler := handler.NewUserHandler(userUsecase)
	e.GET("/users/:id", userHandler.GetUser)

	return e
}
