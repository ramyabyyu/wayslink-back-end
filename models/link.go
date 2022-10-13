package models

type Link struct {
	ID          int      `json:"id" gorm:"primary_key:auto_increment"`
	Title       string   `json:"title" gorm:"varchar(255)"`
	Description string   `json:"description" gorm:"text"`
	Image       string   `json:"image" gorm:"text"`
	UserID      int      `json:"user_id"`
	User        User     `json:"user"`
	Sosmeds     []Sosmed `json:"sosmeds"`
	Template    string   `json:"template"`
	UniqueLink  string   `json:"unique_link"`
}

type UserLink struct {
	ID int `json:"id"`
}

func (UserLink) TableName() string {
	return "links"
}

type LinkSosmed struct {
	ID int `json:"id"`
}
