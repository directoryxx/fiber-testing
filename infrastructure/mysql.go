package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"rest-api/domain"
	"rest-api/helper"
)

func OpenDBMysql(dsn string) (db *gorm.DB,err error) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	autoMigrate(db)

	return db,err
}

func autoMigrate(db *gorm.DB)  {
	db.AutoMigrate(&domain.Role{})
	db.AutoMigrate(&domain.User{})
}