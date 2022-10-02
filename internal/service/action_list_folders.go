package service

import "io/ioutil"

func (s *Service) ActionListFolders() ([]string, error) {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		return nil, err
	}

	dirs := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}

	return dirs, nil
}
