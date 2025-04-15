//go:build integration

package users

import "github.com/stretchr/testify/suite"

const (
	TestName = "users"
)

type MyTestSuite struct {
	suite.Suite
	// Add test-specific fields
}

func (s *MyTestSuite) TestSomething() {
	//panic("test")
	s.T().Log("Running users TestSomething")
	s.Equal(1, 1)
}

func (s *MyTestSuite) TestSomething1() {
	s.T().Log("Running usersTestSomething1")
	s.Equal(1, 1)
}
