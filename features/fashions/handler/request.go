package handler

type InputRequest struct {
	FashionName   string `form:"fashion_name"`
	FashionPoints int    `form:"fashion_points"`
	Status        string `form:"status"`
	// FashionURLImage string `json:"fashion_url_image" form:"fashion_url_image"`
}
