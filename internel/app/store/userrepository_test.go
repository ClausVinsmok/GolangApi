package store_test

import (
	"golang-rest-api/internel/app/model"
	"golang-rest-api/internel/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {

	s, teardown := store.TestStore(t, databaceURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindEmail(t *testing.T) {

	s, teardown := store.TestStore(t, databaceURL)
	defer teardown("users")

	email := "user@gamil.com"

	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
