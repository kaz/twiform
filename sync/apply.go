package sync

import (
	"fmt"
	"log"
)

func (s *Synchronizer) Apply() error {
	for _, user := range s.Effect.PurgeCandidates {
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
