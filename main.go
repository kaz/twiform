package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kaz/twiform/state"
	"github.com/kaz/twiform/sync"
)

func main() {
	if os.Args[1] == "serve" {
		serve()
	}

	s, err := sync.New(state.NewJsonStore("state.json"))
	if err != nil {
		panic(err)
	}

	switch os.Args[1] {
	case "plan":
		s.Plan()
	case "apply":
		s.Apply()
	default:
		fmt.Println("no such subcommand: ", os.Args[1])
	}
}

func serve() {
	fmt.Println("serving at http://localhost:8080")
	panic(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
