package link

import (
	"adv/pkg/db"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	//result := repo.Database.DB.Updates(link)
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)
	return ErrHandling(link, result)
}

func (repo *LinkRepository) Delete(id uint) (*Link, error) {
	result := repo.Database.DB.Delete(&Link{}, id)
	link := &Link{
		Model: gorm.Model{ID: id},
	}
	return ErrHandling(link, result)
}

func ErrHandling(link *Link, result *gorm.DB) (*Link, error) {
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) CheckIfExist(id uint) error {
	result := repo.Database.DB.First(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
