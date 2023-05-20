package model_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/vindosVp/http-rest-api/internal/app/model"
	"testing"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}

type TestCases []TestCase

type TestCase struct {
	name    string
	u       func() *model.User
	isValid bool
}

func TestUser_Validate(t *testing.T) {
	testCases := make(TestCases, 0)

	testCases = append(testCases, TestCase{
		name: "valid",
		u: func() *model.User {
			return model.TestUser(t)
		},
		isValid: true,
	})

	testCases = append(testCases, TestCase{
		name: "with encrypted pwd",
		u: func() *model.User {
			u := model.TestUser(t)
			u.Password = ""
			u.EncryptedPassword = "encPwd"
			return u
		},
		isValid: true,
	})

	testCases = append(testCases, TestCase{
		name: "empty email",
		u: func() *model.User {
			u := model.TestUser(t)
			u.Email = ""
			return u
		},
		isValid: false,
	})

	testCases = append(testCases, TestCase{
		name: "invalid email",
		u: func() *model.User {
			u := model.TestUser(t)
			u.Email = "invalid"
			return u
		},
		isValid: false,
	})

	testCases = append(testCases, TestCase{
		name: "empty password",
		u: func() *model.User {
			u := model.TestUser(t)
			u.Password = ""
			return u
		},
		isValid: false,
	})

	testCases = append(testCases, TestCase{
		name: "short password",
		u: func() *model.User {
			u := model.TestUser(t)
			u.Password = "1"
			return u
		},
		isValid: false,
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}

}
