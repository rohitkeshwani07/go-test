package users

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/rohitkeshwani07/go-test/test3/runner"
)

func TestMain(m *testing.M) {
	// check if there is any arg by the name test
	isArg := false
	for _, arg := range os.Args {
		if arg == "test" {
			log.Println("test arg passed")
			isArg = true
			break
		}
	}
	if isArg == false {
		log.Println("test arg not passed")
		runner.Setup()
	}
	code := m.Run()
	if isArg == false {
		log.Println("test arg not passed")
		runner.Teardown()
	}
	os.Exit(code)
}

func TestSomething(t *testing.T) {
	log.Println("TestSomethingHere")
	time.Sleep(10 * time.Second)
	// fail the test
	t.Fatal("Test failed")
}
