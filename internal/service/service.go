package service

type Service struct{}



func New() *Service {
	return &Service{}
}

func (s *Service) Soundex(name string) ([]string, error) {
	return nil,nil
}
