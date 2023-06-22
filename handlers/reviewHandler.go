package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *Context) InsertLikeReview(ctx echo.Context) error {
	Results := &Result{
		Success: true,
		Data:    nil,
		Code:    http.StatusOK,
		Message: "hit api at '/api/product' successfully",
	}

	return ctx.JSON(200, Results)
}

func (c *Context) DeleteLikeReview(ctx echo.Context) error {
	Results := &Result{
		Success: true,
		Data:    nil,
		Code:    http.StatusOK,
		Message: "hit api at '/api/product' successfully",
	}

	return ctx.JSON(200, Results)
}
