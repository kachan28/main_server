package controller

import (
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

func (c *Controller) DeleteFolder(ctx echo.Context) error {
	bodyBytes, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	folder := struct {
		FolderName string `json:"id" valiadate:"required,gt=0"`
	}{}
	err = jsoniter.Unmarshal(bodyBytes, &folder)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	err = c.validate.Struct(&folder)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	err = c.srv.DeleteFolder(folder.FolderName)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}
