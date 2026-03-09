package user

import (
	"adv/pkg/db"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Database.DB.Create(&user)
	return ErrHandling(user, result)
}

func (repo *UserRepository) Update(user *User) (*User, error) {
	//result := repo.Database.DB.Updates(User)
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(user)
	return ErrHandling(user, result)
}

func (repo *UserRepository) Delete(id uint) (*User, error) {
	result := repo.Database.DB.Delete(&User{}, id)
	User := &User{
		Model: gorm.Model{ID: id},
	}
	return ErrHandling(User, result)
}

func ErrHandling(user *User, result *gorm.DB) (*User, error) {
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) GetById(id uint) (*User, error) {
	var user User
	result := repo.Database.DB.First(&User{}, id)
	return ErrHandling(&user, result)
}

func (repo *UserRepository) GetByEmail(email string) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, "email = ?", email)
	return ErrHandling(&user, result)
}
