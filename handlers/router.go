package handlers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Context struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

type Result struct {
	Success bool
	Data    interface{}
	Code    int
	Message string
}

func (c *Context) Api(group string) {
	public := c.Echo.Group(group)
	{
		public.POST("/member", c.InsertMember)
		public.PUT("/member", c.UpdateMember)
		public.DELETE("/member/:id", c.DeleteMember)
		public.GET("/member", c.GetAllMember)
		public.GET("/product/:id", c.GetProductById)
		public.POST("/like", c.InsertLikeReview)
		public.DELETE("/cancel-like", c.DeleteLikeReview)
	}

}
