package runner_test

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
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
	// lets pass some args
	// multiple test files
	cmd := exec.Command("go", "test", "../users", "../users2", "-v", "-args", "test")

	// Create buffers to capture output
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()

	// Log the output for debugging
	t.Logf("STDOUT: %s", stdout.String())
	t.Logf("STDERR: %s", stderr.String())

	// Check if the command executed successfully
	if err != nil {
		t.Fatalf("Failed to run target test: %v", err)
	}

	// Validate the output to ensure the test passed
	if !strings.Contains(stdout.String(), "PASS") {
		t.Errorf("Target test did not pass. Output: %s", stdout.String())
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
