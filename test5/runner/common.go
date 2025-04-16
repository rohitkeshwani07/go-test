package runner

import (
	"log"
	"os"

	"github.com/rohitkeshwani07/go-test/di"
	"github.com/samber/do"
)

var injector *do.Injector

func Setup() {
	if IsArgPresent("skipsetup") {
		return
	}

	setup()
}

func setup() {
	log.Println("setup")

	injector = do.New()
	di.RegisterServices(injector)

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

func GetString(key string) string {
	return "a"
}

func GetInjectorClone() *do.Injector {
	return injector.Clone()
}
