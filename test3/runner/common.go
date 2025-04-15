package runner

import (
	"log"
	"os"
)

func Setup() {
	isRunner := false
	for _, arg := range os.Args {
		if arg == "test" {
			isRunner = true
			break
		}
	}

	if !isRunner {
		setup()
	}
}

func setup() {
	log.Println("setup")
}

func Teardown() {
	isRunner := false
	for _, arg := range os.Args {
		if arg == "test" {
			isRunner = true
			break
		}
	}

	if !isRunner {
		teardown()
	}
}

func teardown() {
	log.Println("teardown")
}
