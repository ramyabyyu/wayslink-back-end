package sosmeddto

type SosmedResponse struct {
	ID int `json:"id"`
	LinkID      int    `json:"link_id"`
	TitleSosmed string `json:"title_sosmed"`
	Url         string `json:"url"`
	Image       string `json:"image"`
}