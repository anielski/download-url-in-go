package handlers

import (
	"github.com/anielski/download-url-in-go/bindings"
	"github.com/anielski/download-url-in-go/renderings"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// SaveFetcher godoc
// @Summary Save fetcher
// @Description insert or update fetcher
// @Tags API
// @Param verification body bindings.Fetcher true "Transaction Data"
// @Success 200 {object} renderings.Fetcher
// @Failure 400 {object} echo.HTTPError
// @Failure 413 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Accept  application/json
// @Produce  application/json
// @Router /api/fetcher [post]
func (h *Handler) SaveFetcher(c echo.Context) (err error) {
	data := new(bindings.Fetcher)
	if err = c.Bind(data); err != nil { // Bind & validate data
		return err
	}
	if err = c.Validate(data); err != nil {
		return err
	}

	id, err := h.Gwp.SaveFetcher( c, data )
	if err != nil {
		return err
	}

	resp := renderings.Fetcher{
		Id: id,
	}

	return c.JSON(http.StatusOK, resp)
}


// DeleteFetcher godoc
// @Summary Delete fetcher
// @Description set fetcher as delete
// @Tags API
// @Success 200 {object} renderings.Fetcher
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /api/fetcher/{id} [delete]
func (h *Handler) DeleteFetcher(c echo.Context) error {
	id, err := strconv.Atoi( c.Param("id") )
	if err != nil {
		return err
	}

	h.Gwp.DeleteFetcher( c, &id )

	return c.NoContent(http.StatusOK)
}

// getFetcher godoc
// @Summary Get All fetcher
// @Description get all fetcher - not delete
// @Tags API
// @Success 200 {object} renderings.History
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Produce  application/json
// @Router /api/fetcher/{id}/history [get]
func (h *Handler) GetHistory(c echo.Context) error {
	id, err := strconv.Atoi( c.Param("id") )
	if err != nil {
		return err
	}

	rows, err := h.Gwp.GetHistory( c, &id )
	if err != nil {
		return err
	}

	resp := rows

	return c.JSON(http.StatusOK, resp)
}

// getFetcher godoc
// @Summary Get All fetcher
// @Description get all fetcher - not delete
// @Tags API
// @Success 200 {object} renderings.Fetchers
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Produce  application/json
// @Router /api/fetcher [get]
func (h *Handler) GetFetchers(c echo.Context) error {
	rows, err := h.Gwp.GetFetchers( c )
	if err != nil {
		return err
	}

	resp := rows

	return c.JSON(http.StatusOK, resp)
}
