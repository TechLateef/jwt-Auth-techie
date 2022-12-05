package repository

import (
	entity "github.com/techlateef/jwt-Auth-techies/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	GetAllUsers() []entity.Users
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(userRepo *gorm.DB) UserRepository {
	return &userRepository{
		connection: userRepo,
	}
}

func (db *userRepository) CreateUser(user entity.User) entity.User {
	db.connection.Save(&user)
	db.connection.Preload("Users").Find(&user)
	return user
}

func (db *userRepository) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email= ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userRepository) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userRepository) GetAllUsers() []entity.Users {
	var users []entity.Users
	db.connection.Find(&users)
	return users
}
