package controller

import (
	"errors"
	"main_server/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Controller) ListBackups(ctx echo.Context) error {
	folderName := ctx.QueryParam("folder")
	if folderName == "" {
		return ctx.String(http.StatusBadRequest, "folder param was not provided")
	}

	files, err := c.srv.ActionListBackups(folderName)
	if err != nil {
		if errors.Is(err, models.ErrFolderDoesNotExist) {
			return ctx.String(http.StatusUnprocessableEntity, err.Error())
		} else {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
	}

	return ctx.JSON(http.StatusOK, files)
}
