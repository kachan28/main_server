package service

type Service struct{}

func InitializeService() (*Service, error) {
	return &Service{}, nil
}
