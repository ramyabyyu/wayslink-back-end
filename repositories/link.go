package repositories

import (
	"wayslink/models"

	"gorm.io/gorm"
)

type LinkRepository interface {
	CreateLink(link models.Link) (models.Link, error)
	FindLink(userID int) ([]models.Link, error)
	PreviewLink(uniqueLink string) (models.Link, error)
}

func RepositoryLink(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateLink(link models.Link) (models.Link, error) {
	err := r.db.Create(&link).Error

	return link, err
}

func (r *repository) FindLink(userID int) ([]models.Link, error) {
	var links []models.Link
	err := r.db.Find(&links, "user_id = ?", userID).Error

	return links, err
}

func (r *repository) PreviewLink(uniqueLink string) (models.Link, error) {
	var link models.Link
	err := r.db.First(&link, "unique_link = ?", uniqueLink).Error
	return link, err
}
