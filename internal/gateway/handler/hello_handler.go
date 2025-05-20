package handler

import (
	"net/http"
	"sample-server/internal/domain/usecase"

	"github.com/labstack/echo/v4"
)

type HelloHandler struct {
	Usecase usecase.HelloUsecase
}

func NewHelloHandler(u usecase.HelloUsecase) *HelloHandler {
	return &HelloHandler{Usecase: u}
}

func (h *HelloHandler) SayHello(c echo.Context) error {
	name := c.QueryParam("name")
	msg := h.Usecase.SayHello(name)

	return c.JSON(http.StatusOK, map[string]string{"message": msg})
}
