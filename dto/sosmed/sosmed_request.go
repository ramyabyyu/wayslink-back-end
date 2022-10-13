package sosmeddto

type SosmedRequest struct {
	LinkID      int    `json:"link_id"`
	TitleSosmed string `json:"title_sosmed"`
	Url         string `json:"url"`
	Image       string `json:"image"`
}
