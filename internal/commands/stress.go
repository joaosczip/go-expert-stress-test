package commands

import (
	"time"

	stress "github.com/joaosczip/go-expert-stress-test/pkg"
)

type StressCommand struct {
	stress *stress.StressTester
}

func NewStressCommand(concurrency, requests int, url string, timeout int) *StressCommand {
	return &StressCommand{
		stress: stress.NewStressTester(concurrency, requests, url, time.Duration(timeout)*time.Second),
	}
}

func (s *StressCommand) Run() {
	s.stress.Run()
}
