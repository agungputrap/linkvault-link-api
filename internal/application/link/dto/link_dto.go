package dto

type CreateLinkRequest struct {
	Title       string   `json:"title"`
	Url         string   `json:"url"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

type LinkResponse struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	Url         string   `json:"url"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

type UpdateLinkRequest struct {
	Title       string   `json:"title"`
	Url         string   `json:"url"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
