package state

import (
	"encoding/json"
	"fmt"
	"os"
)

type (
	jsonStore struct {
		filePath string
	}
)

func NewJsonStore(filePath string) Store {
	return &jsonStore{filePath}
}

func (s *jsonStore) Load() (*State, error) {
	fp, err := os.Open(s.filePath)
	if err != nil {
		return nil, fmt.Errorf("os.Open failed: %w", err)
	}
	defer fp.Close()

	entry := &State{}
	if err := json.NewDecoder(fp).Decode(entry); err != nil {
		return nil, fmt.Errorf("json.NewDecoder.Decode failed: %w", err)
	}

	return entry, nil
}

func (s *jsonStore) Save(entry *State) error {
	fp, err := os.Create(s.filePath)
	if err != nil {
		return fmt.Errorf("os.Create failed: %w", err)
	}
	defer fp.Close()

	encoder := json.NewEncoder(fp)
	encoder.SetIndent("", "\t")

	if err := encoder.Encode(entry); err != nil {
		return fmt.Errorf("encoder.Encode failed: %w", err)
	}

	return nil
}
