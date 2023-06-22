package handlers

import (
	"github.com/ariashabry/TechnicalTest/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func (c *Context) InsertLikeReview(ctx echo.Context) error {
	l := models.LikeReview{}
	err := ctx.Bind(&l)
	if err != nil {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: "Failed send data",
		}
		log.Println("Error like review:", err.Error)
		return ctx.JSON(http.StatusBadRequest, Results)
	}

	err = l.Like(c.DB)

	//cek error
	if err != nil {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: "Failed like review",
		}
		log.Println("Failed Like review", err.Error())
		return ctx.JSON(http.StatusInternalServerError, Results)
	}

	Results := &Result{
		Success: true,
		Data:    &l,
		Code:    http.StatusCreated,
		Message: "Success Like Review",
	}
	return ctx.JSON(http.StatusCreated, Results)
}

func (c *Context) DeleteLikeReview(ctx echo.Context) error {
	l := models.LikeReview{}
	err := ctx.Bind(&l)
	if err != nil {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: "Failed send data",
		}
		log.Println("Error cancel like review:", err.Error)
		return ctx.JSON(http.StatusBadRequest, Results)
	}

	err = l.CancelLike(c.DB)

	//cek error
	if err != nil {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: "Failed Cancel like review",
		}
		log.Println("Failed Cancel Like review", err.Error())
		return ctx.JSON(http.StatusInternalServerError, Results)
	}

	Results := &Result{
		Success: true,
		Data:    &l,
		Code:    http.StatusOK,
		Message: "Success Cancel Like Review",
	}
	return ctx.JSON(http.StatusOK, Results)
}
