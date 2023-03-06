package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaceURL string) (*Store, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = databaceURL
	sf := New(config)

	store, _ := sf.NewStore()
	defer store.Close()

	return store, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := store.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {

				t.Fatal(err)
			}
		}
	}
}
