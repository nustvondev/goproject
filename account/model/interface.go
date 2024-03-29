package model

import "github.com/google/uuid"

type UserService interface {
	Get(uid uuid.UUID) (*User, error)
}

type UserRepository interface {
	FindByID(uid uuid.UUID) (*User, error)
}
