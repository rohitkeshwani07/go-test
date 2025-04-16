package di

import (
	"github.com/rohitkeshwani07/go-test/users"
	"github.com/samber/do"
)

func RegisterServices(i *do.Injector) {
	do.Provide(i, func(i *do.Injector) (users.IUserRepository, error) {
		return users.NewUserRepository([]*users.User{
			{ID: "1", Name: "John Doe", Email: "john@example.com"},
			{ID: "2", Name: "Jane Doe", Email: "jane@example.com"},
		}), nil
	})

	do.Provide(i, func(i *do.Injector) (users.IUserService, error) {
		repo := do.MustInvoke[users.IUserRepository](i)
		return users.NewUserService(repo), nil
	})
}
