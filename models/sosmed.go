package models

type Sosmed struct {
	ID          int    `json:"id" gorm:"primary_key:auto_increment"`
	TitleSosmed string `json:"title_sosmed" gorm:"varchar(255)"`
	Url         string `json:"url" gorm:"text"`
	Image       string `json:"image" gorm:"text"`
	LinkID      int    `json:"link_id"`
	Link        Link   `json:"link"`
}

type SosmedLink struct {
	ID int `json:"id"`
}

func (SosmedLink) TableName() string {
	return "sosmeds"
}
