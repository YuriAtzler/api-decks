package database

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func CallDatabase(ctx *gin.Context) *gorm.DB {
	if db == nil {
		stringConection := "postgresql://postgres.icgzbdzjeouftniexhku:myNbu8-xyxsyn-votnor@aws-0-sa-east-1.pooler.supabase.com:6543/postgres"
		var err error
		db, err = gorm.Open(postgres.Open(stringConection), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}
	return db
}

func CloseDatabase() error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
