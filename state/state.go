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
		Credentilas *Credentilas `json:"credentials"`

		Ignore []string `json:"ignore"`

		Followers map[string]anaconda.User `json:"followers"`
		Friends   map[string]anaconda.User `json:"friends"`

		Effect *Effect `json:"effect"`
	}

	Credentilas struct {
		ConsumerKey       string `json:"consumer_key"`
		ConsumerSecret    string `json:"consumer_secret"`
		AccessTokenKey    string `json:"access_token_key"`
		AccessTokenSecret string `json:"access_token_secret"`
	}

	Effect struct {
		NotFollowers    []anaconda.User `json:"not_followers"`
		NotFriends      []anaconda.User `json:"not_friends"`
		PurgeCandidates []anaconda.User `json:"purge_candidates"`
	}
)
