package store

import (
	"database/sql"

	_ "github.com/lib/pq" //Anonimn import
)

type Store struct {
	config         *Config
	userRepository *UserRepository
}

func New(c *Config) *Store {
	return &Store{
		config: c,
	}
}

//func (s *Store) Open() error {
//	db, err := sql.Open("postgres", s.config.DatabaseURL)
//
//	if err != nil {
//		return err
//	}
//
//	if err := db.Ping(); err != nil {
//		return err
//	}
//
//	s.db = db
//
//	return nil
//}

func (s *Store) Open() (*sql.DB, error) {
	db, err := sql.Open("postgres", s.config.DatabaseURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (s *Store) Close(db *sql.DB) {
	db.Close()
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
