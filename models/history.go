package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

//model db
type History struct {
	gorm.Model
	UrlID  uint
	Response string `gorm:"size:1048576"`
	Duration float64
}

//saves the history of one object in db as a row
func SaveHistory(id uint, bytes []byte, secs float64, db *gorm.DB) error {
	history := History{
		UrlID:      id,
		Response: string(bytes),
		Duration: secs,
	}
	db.NewRecord(history)
	db.Create(&history)
	if db.NewRecord(history) {
		errString := "some problem with save"
		err := errors.New(errString)
		return err
	}
	return nil
}
