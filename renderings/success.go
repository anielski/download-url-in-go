package renderings

import "time"

// SuccessResponse godoc
type SuccessResponse struct {
	Success bool `json:"success"`
}

type Fetcher struct {
	Id uint `json:"id"`
}

type Fetchers struct {
	Id       uint   `json:"id"`
	Url      string `json:"url"`
	Interval uint   `json:"interval"`
}

type Response struct {
	Response  string     `json:"response"`
	Duration  float32    `json:"duration"`
	CreatedAt *time.Time `json:"created_at"`
}

type History struct {
	Response []*Response
}
