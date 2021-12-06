package teststore

import (
	"github.com/Muhammad-D/http_rest_api/internal/app/model"
	"github.com/Muhammad-D/http_rest_api/internal/app/store"
)

//UserRepository...
type UserRepository struct {
	store *Store
	users map[string]*model.User
}

//Create...
func (r *UserRepository) Create(u *model.User) error {
	//Email and password validations before a user account creation
	if err := u.Validation(); err != nil {
		return err
	}

	//Password Encryption before a user account creation
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

//FindByEmail...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}
