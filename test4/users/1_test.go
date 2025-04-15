package users

import (
	"log"
	"testing"

	"github.com/rohitkeshwani07/go-test/test3/runner"
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
	log.Println("TestSomethingHere")
	// To fail the test with testify:
	// s.Fail("Test failed")
}

func (s *UserTestSuite) TestSomething2() {
	log.Println("TestSomethingHere22")
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
