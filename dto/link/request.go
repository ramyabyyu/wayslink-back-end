package linkdto

type LinkRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Template    string `json:"template"`
}
