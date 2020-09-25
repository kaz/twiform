package sync

import (
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

	purged := followers.Union(friends).Difference(followers.Intersect(friends))
	for _, ig := range s.Ignore {
		purged.Remove(ig)
	}

	s.Effect = &state.Effect{
		NotFollowers:    s.setToUserList(friends.Difference(followers)),
		NotFriends:      s.setToUserList(followers.Difference(friends)),
		PurgeCandidates: s.setToUserList(purged),
	}
}

func (s *Synchronizer) setToUserList(set mapset.Set) []anaconda.User {
	users := make([]anaconda.User, 0, set.Cardinality())
	for key := range set.Iter() {
		if ent, ok := s.Followers[key.(string)]; ok {
			users = append(users, ent)
			continue
		}
		users = append(users, s.Friends[key.(string)])
	}
	return users
}
