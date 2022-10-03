package controller

import (
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

func (c *Controller) DeleteFile(ctx echo.Context) error {
	bodyBytes, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	file := struct {
		FolderName string `json:"folder_name" validate:"required,gt=0"`
		FileName   string `json:"file_name" validate:"required,gt=0"`
	}{}
	err = jsoniter.Unmarshal(bodyBytes, &file)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	err = c.validate.Struct(&file)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	err = c.srv.DeleteFile(file.FolderName, file.FileName)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}
