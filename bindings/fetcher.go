package bindings

// fetcher godoc
type Fetcher struct {
	Id       *uint   `json:"id"`
	Url      *string `json:"url"  validate:"required"`
	Interval *uint   `json:"interval" validate:"required"`
}
