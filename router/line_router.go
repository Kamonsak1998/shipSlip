package router

import (
	slip "shipSlip/slip/handle"

	"github.com/labstack/echo/v4"
)

type Router struct {
	*echo.Echo
}

func New(e *echo.Echo) *Router {
	return &Router{e}
}

func (r *Router) LineRouting() {
	r.POST("/", slip.Handler)
}
