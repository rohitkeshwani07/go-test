//go:build integration

package common_test

import (
	"log"
	"testing"

	"github.com/rohitkeshwani07/go-test/test2/users"
	"github.com/rohitkeshwani07/go-test/test2/users1"
	"github.com/stretchr/testify/suite"
)

type RegisterTestSuite struct {
	suite.Suite
	// Add common fields here (e.g. DB connection, mock clients, etc.)
}

func (s *RegisterTestSuite) SetupSuite() {
	// Run once before all tests
	// Example: start DB container, connect, etc.

	log.Println("SetupSuite1")

}

func (s *RegisterTestSuite) TearDownSuite() {
	log.Println("TearDownSuite1")
	// Run once after all tests
	// Example: close DB, clean up

}

func (s *RegisterTestSuite) SetupTest() {
	// Run before each test
}

func (s *RegisterTestSuite) TearDownTest() {
	// Run after each test
}

func (s *RegisterTestSuite) TestSomething() {
	suite.Run(s.T(), new(users.MyTestSuite))
	suite.Run(s.T(), new(users1.MyTestSuite))
}

func TestSomething(t *testing.T) {
	suite.Run(t, new(RegisterTestSuite))
}
