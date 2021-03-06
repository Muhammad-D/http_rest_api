package teststore_test

import (
	"testing"

	"github.com/Muhammad-D/http_rest_api/internal/app/model"
	"github.com/Muhammad-D/http_rest_api/internal/app/store"
	"github.com/Muhammad-D/http_rest_api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {

	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
}

func TestUserRepository_FindByEmail(t *testing.T) {

	s := teststore.New()
	email := "user@example.com"
	_, err := s.User().FindByEmail(email)

	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	u, err = s.User().FindByEmail(email)

	assert.NotNil(t, u)
	assert.NoError(t, err)

}
