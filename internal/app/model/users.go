package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UUID              uuid.UUID
	Email             string
	Password          string
	EncryptedPassword string
}

func (u *User) BeforeCreate() error {
	return nil
}

func encryptStr(str string) (string, error) {
	encString, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(encString), nil
}
