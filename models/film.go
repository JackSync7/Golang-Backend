package models

type Film struct {
	ID            int      `json:"id" gorm:"primary_key:auto_increment"`
	Title         string   `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string   `json:"thumbnail" form:"thumbnail" gorm:"type: varchar(255)"`
	Year          int      `json:"year" form:"year"`
	Category      Category `json:"category" `
	CategoryID    int      `json:"category_id" `
	Description   string   `json:"description" form:"description" gorm:"type: text"`
}

type FilmResponse struct {
	ID            int      `json:"id" gorm:"primary_key:auto_increment"`
	Title         string   `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string   `json:"thumbnail" form:"thumbnail" gorm:"type: varchar(255)"`
	Year          int      `json:"year" form:"year"`
	Category      Category `json:"category" `
	CategoryID    int      `json:"category_id" `
	Description   string   `json:"description" form:"description" gorm:"type: text"`
}
