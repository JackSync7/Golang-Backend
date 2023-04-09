package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type EpisodeRepository interface {
	FindEpisode() ([]models.Episode, error)
	GetEpisode(ID int) (models.Episode, error)
	CreateEpisode(Episode models.Episode) (models.Episode, error)
	UpdateEpisode(Episode models.Episode) (models.Episode, error)
	DeleteEpisode(Episode models.Episode, ID int) (models.Episode, error)
}

func RepositoryEpisode(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindEpisode() ([]models.Episode, error) {
	var Episode []models.Episode
	err := r.db.Preload("Film").Find(&Episode).Error

	return Episode, err
}

func (r *repository) GetEpisode(ID int) (models.Episode, error) {
	var Episode models.Episode
	err := r.db.Preload("Film").First(&Episode, ID).Error

	return Episode, err
}

func (r *repository) CreateEpisode(Episode models.Episode) (models.Episode, error) {
	err := r.db.Preload("Film").Create(&Episode).Error

	return Episode, err
}

func (r *repository) UpdateEpisode(Episode models.Episode) (models.Episode, error) {
	err := r.db.Save(&Episode).Error

	return Episode, err
}

func (r *repository) DeleteEpisode(Episode models.Episode, ID int) (models.Episode, error) {
	err := r.db.Delete(&Episode, ID).Scan(&Episode).Error

	return Episode, err
}
