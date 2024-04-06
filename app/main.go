package main

import (
	"fmt"

	"github.com/kocubinski/semver-test/lib"
)

func main() {
	fmt.Println("App version: ", lib.Version())
}
