package interfaces

import (
	"url-shortener/internal/domains"
)

type UserRepository interface {
	Save(user *domains.User) error
	FindByEmail(email string) (*domains.User, error)
}
