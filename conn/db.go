package conn

import (
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/satyakarki/did-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func ConnectDB() {
	dbConfig := config.Db()

	log.Info("Connecting to database at "+dbConfig.Host+":", dbConfig.Port)
	logMode := logger.Silent
	if dbConfig.Debug {
		logMode = logger.Info
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.Schema)

	dB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		panic(err)
	}

	sqlDB, err := dB.DB()
	if err != nil {
		panic(err)
	}

	if dbConfig.MaxIdleConn > 0 {
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConn)
	}
	if dbConfig.MaxOpenConn > 0 {
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConn)
	}
	if dbConfig.MaxConnLifetime > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.MaxConnLifetime))
	}

	db = dB
	log.Info("mysql connection successful...")

}

func Db() *gorm.DB {
	return db
}
