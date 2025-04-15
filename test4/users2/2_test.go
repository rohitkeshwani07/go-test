package users2

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
	log.Println("TestSomethingHere2")
	// To fail the test with testify:
	// s.Fail("Test failed")
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
