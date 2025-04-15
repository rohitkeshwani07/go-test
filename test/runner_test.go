//go:build integration

package common_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/rohitkeshwani07/go-test/test/users"
	"github.com/rohitkeshwani07/go-test/test/users1"
	"github.com/rohitkeshwani07/go-test/test/users2"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	fmt.Println("setup")
}

func teardown() {
	fmt.Println("teardown")
}

func TestAll(t *testing.T) {
	t.Run("TestSomething", users.TestSomething)
	t.Run("TestSomething", users1.TestSomething)
	t.Run("TestSomething", users2.TestSomething)
}
