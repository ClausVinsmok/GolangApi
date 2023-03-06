package store

import "database/sql"

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
		userRepository: &UserRepository{db: db},
	}, nil
}

func (s *Store) Close() {
	s.db.Close()
}
