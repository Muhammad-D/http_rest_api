package sqlstore

import (
	"database/sql"

	"github.com/Muhammad-D/http_rest_api/internal/app/store"
	_ "github.com/lib/pq"
)

//Store...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

//New...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//UserRepository...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
