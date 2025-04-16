package runner_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/rohitkeshwani07/go-test/test3/runner"
	"github.com/stretchr/testify/suite"
)

type RunnerTestSuite struct {
	suite.Suite
}

func (s *RunnerTestSuite) SetupSuite() {
	runner.Setup()
}

func (s *RunnerTestSuite) TearDownSuite() {
	runner.Teardown()
}

func (s *RunnerTestSuite) TestRunCommand() {
	cmd := exec.Command("gotestsum", "--", "../users", "../users2", "-v", "-args", "skipsetup")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	s.NoError(err, "Failed to run target test")
}

func TestRunnerSuite(t *testing.T) {
	suite.Run(t, new(RunnerTestSuite))
}
