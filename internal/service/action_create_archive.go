package service

import (
	"fmt"
	pb "main_server/internal/models/pb"
	"os"
	"time"
)

func (s *Service) CreateArchive(create *pb.BackupCreate) error {
	folderName := create.GetId()

	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		err := os.Mkdir(folderName, 0755)
		if err != nil {
			return err
		}
	}

	fileName := fmt.Sprintf("%s_%s.zip", create.File.Title, time.Now().Format("02-01-2006_15:04:05"))

	archive, err := os.Create(fmt.Sprintf("./%s/%s", folderName, fileName))
	if err != nil {
		return err
	}
	defer archive.Close()

	_, err = archive.Write(create.File.Content)
	if err != nil {
		return err
	}

	return nil
}
