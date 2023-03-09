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

	fmt.Print(sf)

	store, err := sf.NewStore()

	if err != nil {
		t.Fatal(err)
	}

	//return store, func(tables ...string) {
	//	if len(tables) > 0 {
	//		err := store.testRepository.TRUNCATE(tables...)
	//		defer store.Close()
	//		if err != nil {
	//			t.Fatal(err)
	//		}
	//	}
	//}

	return store, func(tables ...string) {
		defer store.Close()
		if len(tables) > 0 {
			_, err := store.db.Exec(
				fmt.Sprintf("TRUNCATE %s CASCADE",
					strings.Join(tables, ", ")),
			)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}
