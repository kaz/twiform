package sync

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	mapset "github.com/deckarep/golang-set"
	"github.com/kaz/twiform/state"
)

func (s *Synchronizer) calcEffects() {
	followers := mapset.NewSet()
	for key := range s.Followers {
		followers.Add(key)
	}

	friends := mapset.NewSet()
	for key := range s.Friends {
		friends.Add(key)
	}

	s.Effect = &state.Effect{
		NotFollowers:    s.setToSlice(friends.Difference(followers)),
		NotFriends:      s.setToSlice(followers.Difference(friends)),
		PurgeCandidates: s.setToSlice(followers.Union(friends).Difference(followers.Intersect(friends))),
	}
}

func (s *Synchronizer) getPurges() []anaconda.User {
	purges := mapset.NewSet()
	for _, user := range s.Effect.PurgeCandidates {
		purges.Add(user.ScreenName)
	}
	for _, ignoredScreenName := range s.Ignore {
		purges.Remove(ignoredScreenName)
	}
	return s.setToSlice(purges)
}

func (s *Synchronizer) setToSlice(set mapset.Set) []anaconda.User {
	users := make([]anaconda.User, 0, set.Cardinality())
	for key := range set.Iter() {
		if ent, ok := s.Followers[key.(string)]; ok {
			users = append(users, ent)
			continue
		}
		if ent, ok := s.Friends[key.(string)]; ok {
			users = append(users, ent)
			continue
		}
		panic(fmt.Errorf("not found: %s", key))
	}
	return users
}
