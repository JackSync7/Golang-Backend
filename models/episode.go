package models

type Episode struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string `json:"thumbnail" form:"thumbnail" gorm:"type: varchar(255)"`
	Year          int    `json:"year" form:"year"`
	LinkFilm      string `json:"linkFilm" form:"link" gorm:"type: varchar(255)"`
	Film          Film   `json:"film"`
	FilmID        int    `json:"film_id" `
}

type EpisodeResponse struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string `json:"thumbnail" form:"thumbnail" gorm:"type: varchar(255)"`
	Year          int    `json:"year" form:"year"`
	LinkFilm      string `json:"linkFilm" form:"link" gorm:"type: varchar(255)"`
	Film          Film   `json:"film"`
}
