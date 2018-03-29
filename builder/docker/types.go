package docker

type pullEvent struct {
	Status         string             `json:"status"`
	Error          string             `json:"error"`
	Progress       string             `json:"progress"`
	ProgressDetail pullProgressDetail `json:"progressDetail"`
}

type pullProgressDetail struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}
