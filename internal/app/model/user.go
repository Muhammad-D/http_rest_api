package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

//User...
type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

//Validation METHOD...
func (u *User) Validation() error {

	if err := validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
	); err != nil {
		return err
	}

	return nil

}

//BeforeCreate METHOD...
func (u *User) BeforeCreate() error {

	if len(u.Password) > 0 {
		enc, err := encryptPassword(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
		return nil
	}

	return nil
}

//An encrypting FUNCTION...
func encryptPassword(pw string) (string, error) {

	b, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
