package service

import (
	"os"
	"path/filepath"
)

func (s *Service) DeleteFile(folderName, fileName string) error {
	return os.Remove(filepath.Join(folderName, fileName))
}
