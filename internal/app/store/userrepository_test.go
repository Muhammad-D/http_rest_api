package store_test

import (
	"testing"

	"github.com/Muhammad-D/http_rest_api/internal/app/model"
	"github.com/Muhammad-D/http_rest_api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@example.com"

	_, err := s.User().FindByEmail(email)

	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: "user@example.com",
	})

	u, err := s.User().FindByEmail(email)

	assert.NotNil(t, u)
	assert.NoError(t, err)

}
