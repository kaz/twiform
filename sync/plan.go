package sync

import (
	"fmt"
	"strings"
)

func (s *Synchronizer) Plan() {
	fmt.Println("# Account who will purged")
	fmt.Println("")
	fmt.Println("| name | url |")
	fmt.Println("| --- | --- |")
	for _, user := range s.Effect.PurgeCandidates {
		fmt.Printf("| %s | https://twitter.com/%s |\n", strings.ReplaceAll(user.Name, "|", "\\|"), user.ScreenName)
	}
	fmt.Println("")

	fmt.Println("# Account who is followed by me, but I am not followed back")
	fmt.Println("")
	fmt.Println("| name | url |")
	fmt.Println("| --- | --- |")
	for _, user := range s.Effect.NotFollowers {
		fmt.Printf("| %s | https://twitter.com/%s |\n", strings.ReplaceAll(user.Name, "|", "\\|"), user.ScreenName)
	}
	fmt.Println("")

	fmt.Println("# Account who follows me, but I do not follow back")
	fmt.Println("")
	fmt.Println("| name | url |")
	fmt.Println("| --- | --- |")
	for _, user := range s.Effect.NotFriends {
		fmt.Printf("| %s | https://twitter.com/%s |\n", strings.ReplaceAll(user.Name, "|", "\\|"), user.ScreenName)
	}
	fmt.Println("")
}
