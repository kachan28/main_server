package controller

import (
	"io/ioutil"
	pb "main_server/internal/models/pb"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/proto"
)

func (c *Controller) Create(ctx echo.Context) error {
	bodyBytes, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	b := &pb.BackupCreate{}
	err = proto.Unmarshal(bodyBytes, b)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	err = c.validate.Struct(b)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	err = c.srv.CreateArchive(b)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusCreated)
}
