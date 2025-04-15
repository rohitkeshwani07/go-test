package users

type IUserService interface {
	GetUser(id string) (*User, error)
}

type UserService struct {
	userRepository IUserRepository
}

func NewUserService(userRepository IUserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetUser(id string) (*User, error) {
	return s.userRepository.GetUser(id)
}
