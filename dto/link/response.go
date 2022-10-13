package linkdto

type LinkResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Template    string `json:"template"`
	UniqueLink  string `json:"unique_link"`
}
