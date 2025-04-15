package runner_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/rohitkeshwani07/go-test/test3/runner"
)

func TestMain(m *testing.M) {
	runner.Setup()
	code := m.Run()
	runner.Teardown()
	os.Exit(code)
}

// TestRunnerTest demonstrates running another Go test command from a test
// go clean -testcache && go test ./test3/runner/... -v -run TestRunnerTest
func TestRunnerTest(t *testing.T) {
	// Define the test command to run
	// This will run the "TestTarget" function in the "target_test.go" file
	cmd := exec.Command("go", "test", "../users", "../users2", "-v", "-args", "test")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()

	// Check if the command executed successfully
	if err != nil {
		t.Fatalf("Failed to run target test: %v", err)
	}
}

// Example target test file would look like:
/*
package target

import (
	"testing"
)

func TestTarget(t *testing.T) {
	// This is the test that will be executed by the runner
	t.Log("Target test executed successfully")
}
*/
