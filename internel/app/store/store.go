package store

import (
	"database/sql"
)

type Store struct {
	db             *sql.DB
	testRepository *TestRepository
	userRepository *UserRepository
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{db: s.db}
	return s.userRepository
}

func (s *Store) Test() *TestRepository {
	if s.testRepository != nil {
		return s.testRepository
	}

	s.testRepository = &TestRepository{db: s.db}
	return s.testRepository
}

func (s *Store) Close() {
	s.db.Close()
}
