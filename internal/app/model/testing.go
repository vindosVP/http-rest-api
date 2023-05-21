package model

func TestUser() *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}
