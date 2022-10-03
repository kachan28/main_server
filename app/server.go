package app

import (
	"fmt"
	"main_server/config"

	"main_server/internal/controller"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e *echo.Echo
}

func Init(conf *config.Config) (*Server, error) {
	e := echo.New()
	s := &Server{e: e}
	err := s.registerRoutes()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Server) Start(conf *config.Config) error {
	return s.e.Start(fmt.Sprintf(":%s", conf.HTTP.Port))
}

func (s *Server) registerRoutes() error {
	backupCtrl, err := controller.InitController()
	if err != nil {
		return err
	}

	s.e.POST("/api/backup/save", backupCtrl.Create)
	s.e.GET("/api/backup/list/folders", backupCtrl.ListFolders)
	s.e.GET("/api/backup/list/files", backupCtrl.ListBackups)
	s.e.POST("/api/backup/delete/folder", backupCtrl.DeleteFolder)
	s.e.POST("/api/backup/delete/file", backupCtrl.DeleteFile)
	return nil
}
