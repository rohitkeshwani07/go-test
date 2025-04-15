//go:build integration

package users2

import "github.com/stretchr/testify/suite"

const (
	TestName = "users"
)

type MyTestSuite struct {
	suite.Suite
	// Add test-specific fields
}

func (s *MyTestSuite) TestSomething() {
	s.T().Log("Running users2 TestSomething")
	s.Equal(1, 1)
}

func (s *MyTestSuite) TestSomething1() {
	s.T().Log("Running users2 TestSomething1")
	s.Equal(1, 1)
}
