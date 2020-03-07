package sync

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/kaz/twiform/auth"
	"github.com/kaz/twiform/state"
)

type (
	Synchronizer struct {
		*state.State

		store  state.Store
		client *anaconda.TwitterApi
	}
)

func New(store state.Store) (*Synchronizer, error) {
	st, err := store.Load()
	if err != nil {
		return nil, fmt.Errorf("store.Load failed: %w", err)
	}

	s := &Synchronizer{st, store, nil}
	defer s.Save()

	if err := s.authorize(); err != nil {
		return nil, fmt.Errorf("s.authorize failed: %w", err)
	}

	s.client = anaconda.NewTwitterApiWithCredentials(s.Credentilas.AccessTokenKey, s.Credentilas.AccessTokenSecret, s.Credentilas.ConsumerKey, s.Credentilas.ConsumerSecret)

	if err := s.syncFollowers(); err != nil {
		return nil, fmt.Errorf("s.syncFollowers failed: %w", err)
	}

	if err := s.syncFriends(); err != nil {
		return nil, fmt.Errorf("s.syncFriends failed: %w", err)
	}

	return s, nil
}

func (s *Synchronizer) Save() {
	if err := s.store.Save(s.State); err != nil {
		panic(fmt.Errorf("s.store.Save failed: %w", err))
	}
}

func (s *Synchronizer) authorize() error {
	if s.Credentilas.AccessTokenKey != "" && s.Credentilas.AccessTokenSecret != "" {
		return nil
	}
	if s.Credentilas.ConsumerKey == "" || s.Credentilas.ConsumerSecret == "" {
		return fmt.Errorf("no CK/CS")
	}

	var err error
	s.Credentilas.AccessTokenKey, s.Credentilas.AccessTokenSecret, err = auth.Authorize(s.Credentilas.ConsumerKey, s.Credentilas.ConsumerSecret)
	if err != nil {
		return fmt.Errorf("auth.Authorize failed: %w", err)
	}

	return nil
}

func (s *Synchronizer) syncFollowers() error {
	if s.Followers != nil {
		return nil
	}

	param, err := url.ParseQuery("count=200")
	if err != nil {
		return fmt.Errorf("url.ParseQuery failed: %w", err)
	}

	s.Followers = map[int64]anaconda.User{}
	for page := range s.client.GetFollowersListAll(param) {
		if page.Error != nil {
			return fmt.Errorf("s.client.GetFollowersListAll failed: %w", page.Error)
		}

		for _, user := range page.Followers {
			s.Followers[user.Id] = user
		}
	}

	return nil
}

func (s *Synchronizer) syncFriends() error {
	if s.Friends != nil {
		return nil
	}

	param, err := url.ParseQuery("count=200")
	if err != nil {
		return fmt.Errorf("url.ParseQuery failed: %w", err)
	}

	s.Friends = map[int64]anaconda.User{}
	for page := range s.client.GetFriendsListAll(param) {
		if page.Error != nil {
			return fmt.Errorf("s.client.GetFriendsListAll failed: %w", page.Error)
		}

		for _, user := range page.Friends {
			s.Friends[user.Id] = user
		}
	}

	return nil
}
