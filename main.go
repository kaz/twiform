package main

import (
	"github.com/kaz/twiform/state"
	"github.com/kaz/twiform/sync"
)

func main() {
	syn, err := sync.New(state.NewJsonStore("state.json"))
	if err != nil {
		panic(err)
	}

	syn.Plan()
}
