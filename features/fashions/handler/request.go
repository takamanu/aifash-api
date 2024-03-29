package handler

type InputRequest struct {
	FashionName   string `json:"fashion_name" form:"fashion_name"`
	FashionPoints int    `json:"fashion_points" form:"fashion_points"`
	Status        string `json:"status" form:"status"`
	// FashionURLImage string `json:"fashion_url_image" form:"fashion_url_image"`
}
