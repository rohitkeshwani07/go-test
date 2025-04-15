package main

import (
	"fmt"
	"log"

	"github.com/samber/do"

	"github.com/rohitkeshwani07/go-test/di"
	"github.com/rohitkeshwani07/go-test/users"
)

func main() {
	injector := do.New()

	di.RegisterServices(injector)

	userService := do.MustInvoke[users.IUserService](injector)
	user, err := userService.GetUser("1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}
