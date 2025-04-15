package users

import "errors"

type IUserRepository interface {
	GetUser(id string) (*User, error)
}

type UserRepository struct {
	users []*User
}

func NewUserRepository(users []*User) *UserRepository {
	return &UserRepository{users: users}
}

func (r *UserRepository) GetUser(id string) (*User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
