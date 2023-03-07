package store

import (
	"testing"
)

func TestStore(t *testing.T, databaceURL string) (*Store, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = databaceURL
	sf := New(config)

	store, _ := sf.NewStore()

	return store, func(tables ...string) {
		if len(tables) > 0 {
			err := store.testRepository.TRUNCATE(tables...)
			defer store.Close()
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}
