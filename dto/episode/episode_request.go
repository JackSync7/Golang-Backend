package episodedto

import "dumbmerch/models"

type EpisodeRequest struct {
	ID            int         `json:"id" gorm:"primary_key:auto_increment"`
	Title         string      `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string      `json:"thumbnail" form:"thumbnail" gorm:"type: varchar(255)"`
	Year          int         `json:"year" form:"year"`
	LinkFilm      string      `json:"linkFilm" form:"link" gorm:"type: varchar(255)"`
	FilmID        int         `json:"film_id" `
	Film          models.Film `json:"film"`
}
