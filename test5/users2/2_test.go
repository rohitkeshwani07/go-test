package users2

import (
	"testing"

	"github.com/rohitkeshwani07/go-test/mocks"
	"github.com/rohitkeshwani07/go-test/test5/runner"
	"github.com/rohitkeshwani07/go-test/users"
	"github.com/samber/do"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
}

func (s *UserTestSuite) SetupSuite() {
	runner.Setup()
}

func (s *UserTestSuite) TearDownSuite() {
	runner.Teardown()
}

func (s *UserTestSuite) TestSomething() {
	i := runner.GetInjectorClone()
	defer i.Shutdown()

	do.Override(i, func(i *do.Injector) (users.IUserRepository, error) {
		userRepository := mocks.NewIUserRepository(s.T())
		userRepository.On("GetUser", "1").Return(&users.User{ID: "1", Name: "John Doe", Email: "john@example.com"}, nil)
		return userRepository, nil
	})

	// Get injector from TestMain
	userService, err := do.Invoke[users.IUserService](i)
	if err != nil {
		s.T().Errorf("Expected no error, got %v", err)
	}

	user, err := userService.GetUser("1")
	if err != nil {
		s.T().Errorf("Expected no error, got %v", err)
	}

	if user.ID != "1" {
		s.T().Errorf("Expected user ID to be 1, got %v", user.ID)
	}
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
