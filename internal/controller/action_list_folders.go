package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Controller) ListFolders(ctx echo.Context) error {
	dirs, err := c.srv.ActionListFolders()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, dirs)
}
