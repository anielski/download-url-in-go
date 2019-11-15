package models

import (
	"github.com/anielski/download-url-in-go/renderings"
	"github.com/jinzhu/gorm"
)

//model db
type Url struct {
	gorm.Model
	Url      string `gorm:"size:255"`
	Interval uint   `gorm:"not null"`
	Historys []History	 `gorm:"foreignkey:UrlRefer"`
}

//gets all "fetcher"
//@todo danger of get too many objects at once
func GetFetchers(db *gorm.DB) ([]*renderings.Fetchers, error) {
	var result []*renderings.Fetchers

	rows, err := db.Table("urls").Select("`id`, `url`, `interval`").Where("deleted_at is null").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var url renderings.Fetchers
		db.ScanRows(rows, &url)
		result = append(result, &url)
	}
	return result, nil
}

//get one object
func GetFetcher(id uint, db *gorm.DB) *Url {
	url := &Url{}
	db.First(url, id)
	return url
}
