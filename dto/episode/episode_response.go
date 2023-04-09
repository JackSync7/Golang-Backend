package episodedto

import "dumbmerch/models"

type EpisodeResponse struct {
	Title         string              `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string              `json:"thumbnail" form:"thumbnail" gorm:"type: varchar(255)"`
	Year          int                 `json:"year" form:"year"`
	LinkFilm      string              `json:"linkFilm" form:"link" gorm:"type: varchar(255)"`
	Film          models.FilmResponse `json:"film"`
	FilmID        int                 `json:"film_id" `
}
