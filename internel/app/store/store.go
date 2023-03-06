package store

import (
	"database/sql"

	_ "github.com/lib/pq" //Anonimn import
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		db: s.db,
	}

	return s.userRepository
}
