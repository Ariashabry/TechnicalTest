package handlers

import (
	"github.com/ariashabry/TechnicalTest/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func (c *Context) GetProductById(ctx echo.Context) error {

	id, _ := strconv.Atoi(ctx.Param("id"))

	p := models.Product{}

	err := p.GetById(c.DB, id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Product Not Found")
			return ctx.JSON(http.StatusNotFound, nil)
		} else {
			log.Println("Failed getting data", err.Error())
			return ctx.JSON(http.StatusInternalServerError, nil)
		}
	}

	r := models.Reviews{}
	err = r.GetReview(c.DB, id)

	if err != nil {
		Results := map[string]interface{}{
			"Success": true,
			"Data":    &p,
			"Review":  nil,
			"Code":    http.StatusOK,
			"Message": "Success Getting Data",
		}
		return ctx.JSON(http.StatusOK, Results)
	}

	Results := map[string]interface{}{
		"Success":     true,
		"DataProduct": &p,
		"DataReview":  &r,
		"Code":        http.StatusOK,
		"Message":     "Success Getting Data",
	}
	return ctx.JSON(http.StatusOK, Results)

}
