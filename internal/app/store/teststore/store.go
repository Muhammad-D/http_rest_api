package teststore

import (
	"github.com/Muhammad-D/http_rest_api/internal/app/model"
	"github.com/Muhammad-D/http_rest_api/internal/app/store"
)

//Store...
type Store struct {
	userRepository *UserRepository
}

//New...
func New() *Store {
	return &Store{}
}

//UserRepository...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
