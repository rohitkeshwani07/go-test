//go:build integration

package users1

import "github.com/stretchr/testify/suite"

const (
	TestName = "users"
)

type MyTestSuite struct {
	suite.Suite
	// Add test-specific fields
}

func (s *MyTestSuite) TestSomething() {
	s.T().Log("Running user1 TestSomething")
	s.Equal(1, 1)
}

func (s *MyTestSuite) TestSomething1() {
	s.T().Log("Running user1 TestSomething1")
	s.Equal(1, 1)
}
