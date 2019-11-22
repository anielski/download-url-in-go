package handlers

import (
	"github.com/anielski/download-url-in-go/bindings"
	"github.com/anielski/download-url-in-go/renderings"
	"github.com/labstack/echo"
)

type Gwp interface {
	SaveFetcher(c echo.Context, fetcher *bindings.Fetcher) (id uint, err error)
	DeleteFetcher(c echo.Context, id *int)
	GetHistory(c echo.Context, id *int) ( *renderings.History, error)
	GetFetchers(c echo.Context) ( []*renderings.Fetchers, error)
}

// Handler godoc
type Handler struct {
	Gwp
}
