package dto

type CreateURLRequest struct {
	URL string `json:"url"`
}

type CreateURLResponse struct {
	ShortCode   string `json:"short_code"`
	OriginalURL string `json:"original_url"`
}
