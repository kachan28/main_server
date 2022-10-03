package service

import "os"

func (s *Service) DeleteFolder(folderName string) error {
	return os.RemoveAll(folderName)
}
