package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/gormigrate.v1"
	"log"
)

//set conn to DB
//@todo change this function as Singleton and secure against restarting @see sync.Once
func NewDB(dataSourceName string) *gorm.DB {
	var err error
	db, err := gorm.Open("mysql", dataSourceName)

	if err != nil && db.DB().Ping() != nil {
		log.Panic(err)
	}
	// Migrate the schema
	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.LogMode(true)

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// create url table
		{
			ID: "201910281800",
			Migrate: func(tx *gorm.DB) error {
				type Url struct {
					gorm.Model
					Url      *string `gorm:"size:255"`
					Interval *uint   `gorm:"not null"`
					Historys []History	 `gorm:"foreignkey:UrlRefer"`
				}
				return tx.AutoMigrate(&Url{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("url").Error
			},
		},
		// create param table
		{
			ID: "201910281801",
			Migrate: func(tx *gorm.DB) error {
				type History struct {
					gorm.Model
					UrlID    uint
					Response string `gorm:"size:1048576"`
					Duration float32
				}
				return tx.AutoMigrate(&History{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("history").Error
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
