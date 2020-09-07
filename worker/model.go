package worker

type Page struct {
	HasMore   bool      `json:"hasMore"`
	Page      int       `json:"page"`
	PageCount int       `json:"pageCount"`
	Pictures  []Picture `json:"pictures"`
}

type Picture struct {
	ID             string `json:"id"`
	Author         string `json:"author"`
	Camera         string `json:"camera"`
	CroppedPicture string `json:"cropped_picture"`
	FullPicture    string `json:"full_picture"`
	Tags           string `json:"tags"`
}

type AuthResponse struct {
	Auth  bool   `json:"auth"`
	Token string `json:"token"`
}

type AuthRequest struct {
	ApiKey string `json:"apiKey"`
}
