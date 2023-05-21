package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UUID              uuid.UUID `json:"UUID"`
	Email             string    `json:"email"`
	Password          string    `json:"password,omitempty"`
	EncryptedPassword string    `json:"-"`
}

func (u *User) BeforeCreate() error {

	if len(u.Password) > 0 {
		enc, err := encryptStr(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
	}

	return nil
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ComparePwd(pwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pwd)) == nil
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(RequieredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
	)
}

func encryptStr(str string) (string, error) {
	encString, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(encString), nil
}
