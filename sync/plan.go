package sync

import (
	"fmt"
	"strings"

	"github.com/golang-collections/collections/set"
)

type (
	result struct {
		notFollower *set.Set
		notFriends  *set.Set
	}
)

func (s *Synchronizer) Plan() {
	r := s.plan()

	fmt.Println("# Person who is followed by me, but I am not followed back")
	fmt.Println("")
	fmt.Println("| id | name | url |")
	fmt.Println("| --- | --- | --- |")
	for key, user := range s.Friends {
		if !r.notFollower.Has(key) {
			continue
		}
		fmt.Printf("| %d | %s | https://twitter.com/%s |\n", user.Id, strings.ReplaceAll(user.Name, "|", "\\|"), user.ScreenName)
	}

	fmt.Println("")

	fmt.Println("# Person who follows me, but I do not follow back")
	fmt.Println("")
	fmt.Println("| id | name | url |")
	fmt.Println("| --- | --- | --- |")
	for key, user := range s.Followers {
		if !r.notFriends.Has(key) {
			continue
		}
		fmt.Printf("| %d | %s | https://twitter.com/%s |\n", user.Id, strings.ReplaceAll(user.Name, "|", "\\|"), user.ScreenName)
	}
}

func (s *Synchronizer) plan() *result {
	followers := set.New()
	for key, _ := range s.Followers {
		followers.Insert(key)
	}

	friends := set.New()
	for key, _ := range s.Friends {
		friends.Insert(key)
	}

	return &result{
		notFollower: friends.Difference(followers),
		notFriends:  followers.Difference(friends),
	}
}
