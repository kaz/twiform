package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kaz/twiform/state"
	"github.com/kaz/twiform/sync"
)

func main() {
	s, err := sync.New(state.NewJsonStore("state.json"))
	if err != nil {
		panic(err)
	}

	switch os.Args[1] {
	case "serve":
		serve()
	case "clean":
		s.Clean()
	case "plan":
		if err := s.Sync(); err != nil {
			panic(err)
		}
		s.Plan()
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

func serve() {
	fmt.Println("serving at http://localhost:8080")
	panic(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
