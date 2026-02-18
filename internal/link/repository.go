package link

import (
	"adv/pkg/db"

	"gorm.io/gorm"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(&link)
	return ErrHandling(link, result)
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, "hash = ?", hash)
	return ErrHandling(&link, result)
}

func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	result := repo.Database.DB.Updates(link)
	return ErrHandling(link, result)
}

func ErrHandling(link *Link, result *gorm.DB) (*Link, error) {
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}
