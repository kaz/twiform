package sync

import (
	"fmt"
	"log"
	"strings"
)

func (s *Synchronizer) Apply() error {
	purges := s.getPurges()

	for _, user := range purges {
		fmt.Printf("* @%s\n", user.ScreenName)
	}
	fmt.Printf("\n[!] %d accounts above will be purged. Are you sure? [Y/n]: ", len(purges))

	input := ""
	if _, err := fmt.Scanln(&input); err != nil {
		return fmt.Errorf("fmt.Scanln: %w", err)
	}

	if strings.ToLower(input)[0] != 'y' {
		fmt.Println("Cancelled")
		return nil
	}

	for _, user := range purges {
		if _, err := s.client.BlockUserId(user.Id, nil); err != nil {
			return fmt.Errorf("s.client.BlockUserId failed: %w", err)
		}
		if _, err := s.client.UnblockUserId(user.Id, nil); err != nil {
			return fmt.Errorf("s.client.UnblockUserId failed: %w", err)
		}
		log.Println("Purged @" + user.ScreenName)
	}
	return nil
}
