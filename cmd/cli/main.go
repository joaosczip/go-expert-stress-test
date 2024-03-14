package main

import (
	"time"

	stress "github.com/joaosczip/go-expert-stress-test/pkg"
)

func main() {
	tester := stress.NewStressTester(10, 1000, "https://instagram.com", time.Duration(3*time.Second))
	tester.Run()
}
