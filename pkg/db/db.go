package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type MySQLOptions struct {
	Host                  string
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	LogLevel              int
}

func (o *MySQLOptions) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%v&loc=%s",
		o.Username,
		o.Password,
		o.Host,
		o.Database,
		true,
		"Local",
	)
}

func NewMySQL(opts *MySQLOptions) (*gorm.DB, error) {
	logLeve := logger.Silent
	if opts.LogLevel != 0 {
		logLeve = logger.LogLevel(opts.LogLevel)
	}
	db, err := gorm.Open(
		mysql.Open(opts.DSN()),
		&gorm.Config{
			Logger: logger.Default.LogMode(logLeve),
		},
	)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)

	// SetConnMaxLifetime 设置连接可重用的最长时间
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

	// SetMaxIdleConns 设置空闲连接池的最大连接数
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

	return db, nil
}
