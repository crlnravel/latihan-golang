package todo

//// Service is an interface from which our api module can access our repository of all our models
//type Service interface {
//	InsertOption(option *entities.Option) error
//	FetchOptions() (*[]entities.Option, error)
//	FetchOptionByID(id string) (*entities.Option, error)
//	UpdateOption(option *entities.Option) error
//	RemoveOption(ID string) error
//}
//
//type service struct {
//	repository Repository
//}
//
//func NewService(r Repository) Service {
//	return &service{
//		repository: r,
//	}
//}
//
//func (s *service) InsertOption(option *entities.Option) error {
//	return s.repository.CreateOption(option)
//}
//
//func (s *service) FetchOptions() (*[]entities.Option, error) {
//	return s.repository.FindAllOptions()
//}
//
//func (s *service) FetchOptionByID(id string) (*entities.Option, error) {
//	parsed, err := uuid.FromBytes([]byte(id))
//	if err != nil {
//		return nil, err
//	}
//	return s.repository.FindOptionByID(parsed)
//}
//
//func (s *service) UpdateOption(option *entities.Option) error {
//	return s.repository.UpdateOption(option)
//}
//
//func (s *service) RemoveOption(id string) error {
//	parsed, err := uuid.FromBytes([]byte(id))
//	if err != nil {
//		return err
//	}
//	return s.repository.DeleteOption(parsed)
//}
