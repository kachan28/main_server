package service

import (
	"fmt"
	"io/ioutil"
	"main_server/internal/models"
	"os"
)

func (s *Service) ActionListBackups(folderName string) ([]string, error) {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		return nil, models.ErrFolderDoesNotExist
	}

	files := make([]string, 0)
	contents, err := ioutil.ReadDir(fmt.Sprintf("./%s", folderName))
	if err != nil {
		return nil, err
	}

	for _, content := range contents {
		files = append(files, content.Name())
	}

	return files, nil
}
