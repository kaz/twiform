package state

import (
	"github.com/ChimeraCoder/anaconda"
)

type (
	Store interface {
		Load() (*State, error)
		Save(*State) error
	}

	State struct {
		Credentilas struct {
			ConsumerKey       string `json:"consumer_key"`
			ConsumerSecret    string `json:"consumer_secret"`
			AccessTokenKey    string `json:"access_token_key"`
			AccessTokenSecret string `json:"access_token_secret"`
		} `json:"credentials"`

		Followers map[int64]anaconda.User `json:"followers"`
		Friends   map[int64]anaconda.User `json:"friends"`
	}
)
