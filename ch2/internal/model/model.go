package model

import (
	"database/sql"
	"fmt"

	"github.com/chiti62/gogogo/ch2/global"
	"github.com/chiti62/gogogo/ch2/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Model struct {
	ID          uint32 `gorm:"primary_key" json:"id"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
	CreateTime  uint32 `json:"create_time"`
	UpdateTime  uint32 `json:"update_time"`
	DeletedTime uint32 `json:"delete_time"`
	IsDel       uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*sql.DB, error) {
	//data source name
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.Config.Logger = logger.Default.LogMode(logger.Info)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// sqlDB.SingularTable(true) //schema naming strategy
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return sqlDB, nil
}
