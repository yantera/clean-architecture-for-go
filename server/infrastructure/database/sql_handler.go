package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"api/server/adapters/gateways"
)

type SQLHandler struct {
	db *gorm.DB
}

func InitSQLHandler() gateways.SQLHandler {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	containerName := os.Getenv("DB_CONTAINER_NAME")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, containerName, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.db = db
	return sqlHandler
}

func (handler *SQLHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.db.First(out, where...)
}

func (handler *SQLHandler) Take(out interface{}, where ...interface{}) *gorm.DB {
	return handler.db.Take(out, where...)
}

func (handler *SQLHandler) Last(out interface{}, where ...interface{}) *gorm.DB {
	return handler.db.Last(out, where...)
}

func (handler *SQLHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.db.Find(out, where...)
}
