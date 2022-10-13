package repositories

import (
	"wayslink/models"

	"gorm.io/gorm"
)

type SosmedRepository interface {
	CreateSosmed(sosmed models.Sosmed) (models.Sosmed, error)
	FindSosmedsByLinkID(linkID int) ([]models.Sosmed, error)
}

func RepositorySosmed(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateSosmed(sosmed models.Sosmed) (models.Sosmed, error) {
	err := r.db.Create(&sosmed).Error

	return sosmed, err

}

func (r *repository) FindSosmedsByLinkID(linkID int) ([]models.Sosmed, error) {
	var sosmeds []models.Sosmed

	err := r.db.Find(&sosmeds, "link_id = ?", linkID).Error

	return sosmeds, err
}
