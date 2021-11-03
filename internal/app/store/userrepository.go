package store

import "github.com/Muhammad-D/http_rest_api/internal/app/model"

//UserRepository...
type UserRepository struct {
	store *Store
}

//Create METHOD...
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	//Email and password validations before a user account creation
	if err := u.Validation(); err != nil {
		return nil, err
	}

	//Password Encryption before a user account creation
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

//FindByEmail METHOD...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{
		Email: email,
	}

	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email=$1",
		u.Email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}

	return u, nil
}
