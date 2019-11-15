package v1

import (
	"errors"
	"fmt"
	"github.com/anielski/download-url-in-go/app"
	"github.com/anielski/download-url-in-go/bindings"
	"github.com/anielski/download-url-in-go/handlers"
	"github.com/anielski/download-url-in-go/models"
	"github.com/anielski/download-url-in-go/renderings"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type GWP struct {
	handlers.Gwp
}

func (g GWP) SaveFetcher(c echo.Context, fetcher *bindings.Fetcher) (id uint, err error) {
	url := models.Url{
		Url: *fetcher.Url,
		Interval: *fetcher.Interval,
	}
	if fetcher.Id != nil {
		app.Logger(c).Info("Edit ID: ", *fetcher.Id)
		url.ID = *fetcher.Id
		c.Get("db").(*gorm.DB).Save(url)
		EditFetcher( url.ID, url.Url, url.Interval )
	} else {
		c.Get("db").(*gorm.DB).NewRecord(url)
		c.Get("db").(*gorm.DB).Create(&url)
		if c.Get("db").(*gorm.DB).NewRecord(url) {
			errString := "some problem with save"
			app.Logger(c).Errorf(errString)
			err = errors.New(errString)
			return 0, err
		}
		AddFetcher( url.ID, url.Url, url.Interval )
	}
	return url.ID, nil
}

func (g GWP) DeleteFetcher(c echo.Context, id *int) {
	c.Get("db").(*gorm.DB).Where("id LIKE ?", *id).Delete(models.Url{})
	DeleteFetcherWork( uint(*id) )
}

func (g GWP) GetFetchers(c echo.Context) ( []*renderings.Fetchers, error) {
	return models.GetFetchers( c.Get("db").(*gorm.DB) )
}

func (g GWP) GetHistory(c echo.Context, id *int) ( *renderings.History, error) {
	result := &renderings.History{}
	rows, err := c.Get("db").(*gorm.DB).Table("histories").Select("`response`, `duration`, `created_at`").Where(fmt.Sprintf("url_id = %d", *id) ).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var response renderings.Response
		c.Get("db").(*gorm.DB).ScanRows(rows, &response)
		result.Response = append( result.Response, &response)
	}
	return result, nil
}