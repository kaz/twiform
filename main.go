package main

import (
	"fmt"
	"os"

	"github.com/kaz/twiform/plan"
	"github.com/kaz/twiform/state"
	"github.com/kaz/twiform/sync"
)

func main() {
	s, err := sync.New(state.NewJsonStore("state.json"))
	if err != nil {
		panic(err)
	}

	switch os.Args[1] {
	case "clean":
		s.Clean()
	case "plan":
		if err := s.Sync(); err != nil {
			panic(err)
		}
		plan.StartServer()
	case "apply":
		if err := s.Sync(); err != nil {
			panic(err)
		}
		if err := s.Apply(); err != nil {
			panic(err)
		}
	default:
		fmt.Println("no such subcommand: ", os.Args[1])
	}
}
