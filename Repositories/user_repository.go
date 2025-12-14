package repositories

import domain "taskmanager/Domain"

type UserRepository interface {
	Create(user domain.User) (*domain.User, error)
	FindByUsername(username string) (*domain.User, error)
	Promote(username string) error
	CountUsers() (int64, error)
}
