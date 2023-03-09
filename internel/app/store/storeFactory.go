package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type StoreFactory struct {
	config *Config
}

func New(c *Config) *StoreFactory {
	return &StoreFactory{config: c}
}

func (s *StoreFactory) NewStore() (*Store, error) {
	db, err := sql.Open("postgres", s.config.DatabaseURL)

	if err != nil {
		return nil, err
	}

	return &Store{
		db:             db,
		testRepository: &TestRepository{db: db},
		userRepository: &UserRepository{db: db},
	}, nil
}
