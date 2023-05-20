package model

import "github.com/google/uuid"

type User struct {
	UUID              uuid.UUID
	Email             string
	EncryptedPassword string
}
