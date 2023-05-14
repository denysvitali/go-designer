package designer

type ImageUrlsThumbnails struct {
	ImageUrl      string `json:"ImageUrl"`
	ThumbnailData string `json:"ThumbnailData"`
}

type Response struct {
	B64Images          []string              `json:"b64_images"`
	ImageUrls          []string              `json:"image_urls"`
	ImageUrlsThumbnail []ImageUrlsThumbnails `json:"image_urls_thumbnail"`
}
