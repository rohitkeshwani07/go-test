package runner

import (
	"log"
	"os"
)

func Setup() {
	if IsArgPresent("skipsetup") {
		return
	}

	setup()
}

func setup() {
	log.Println("setup")
}

func Teardown() {
	if IsArgPresent("skipsetup") {
		return
	}

	teardown()
}

func teardown() {
	log.Println("teardown")
}

func IsArgPresent(arg string) bool {
	for _, a := range os.Args {
		if a == arg {
			return true
		}
	}
	return false
}
