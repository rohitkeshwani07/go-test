package users2

import (
	"log"
	"os"
	"testing"

	"github.com/rohitkeshwani07/go-test/test3/runner"
)

func TestMain(m *testing.M) {
	runner.Setup()

	code := m.Run()

	runner.Teardown()

	os.Exit(code)
}

func TestSomething(t *testing.T) {
	log.Println("TestSomethingHere2")
}
