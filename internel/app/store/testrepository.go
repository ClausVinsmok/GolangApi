package store

import (
	"database/sql"
	"fmt"
	"strings"
)

type TestRepository struct {
	db *sql.DB
}

func (r *TestRepository) TRUNCATE(tables ...string) error {
	if _, err := r.db.Exec(
		fmt.Sprintf("TRUNCATE %s CASCADE",
			strings.Join(tables, ", ")),
	); err != nil {
		return err
	}
	return nil
}
