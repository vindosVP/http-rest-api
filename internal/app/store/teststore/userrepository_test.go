package teststore_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/vindosVp/http-rest-api/internal/app/logger"
	"github.com/vindosVp/http-rest-api/internal/app/model"
	"github.com/vindosVp/http-rest-api/internal/app/store"
	"github.com/vindosVp/http-rest-api/internal/app/store/teststore"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {

	_ = logger.ConfigureLogger("debug")
	s := teststore.New()

	u := model.TestUser()
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {

	_ = logger.ConfigureLogger("debug")
	s := teststore.New()

	email := "user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser()
	u.Email = email
	err = s.User().Create(u)

	if err != nil {
		t.Fatal(err)
	}

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
