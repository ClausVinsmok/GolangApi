package store_test

import (
	"golang-rest-api/internel/app/model"
	"golang-rest-api/internel/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Crate(t *testing.T) {

	s, teardown := store.TestStore(t, databaceURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email:             "user@gamil.com",
		EncryptedPassword: "1111",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindEmail(t *testing.T) {

	s, teardown := store.TestStore(t, databaceURL)
	defer teardown("users")

	email := "user@gamil.com"

	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email:             "user@gamil.com",
		EncryptedPassword: "1111",
	})

	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
