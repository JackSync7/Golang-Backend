package models

type Transaction struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment"`
	StartDate string `json:"startDate" gorm:"type: varchar(255)"`
	EndDate   string `json:"endDate" gorm:"type: varchar(255)"`
	User      User   `json:"user"`
	UserID    int    `json:"userID"`
	Attach    string `json:"attach" gorm:"type:varchar(255)"`
	Status    string `json:"status" gorm:"type: varchar(255)"`
}
