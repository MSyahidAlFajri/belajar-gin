package user

import "gorm.io/gorm"

type UserRepository interface {
	FindAll() []User
	FindOne(id int) User
	Save(user User) (*User, error)
	Update(user User) (*User, error)
	Delete(user User) (*User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}
