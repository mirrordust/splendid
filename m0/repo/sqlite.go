package repo

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	log.Println("init sqlite connection...")

	gormDB, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Panicln("failed to connect database")
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Panicln("failed to get sqlDB")
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = gormDB
	DB.AutoMigrate(&Post{}, &Tag{}, &User{})
}

func Create(model interface{}) error {
	result := DB.Create(model)
	return result.Error
}

func FindOne(model interface{}, conds ...interface{}) error {
	result := DB.First(model, conds...)
	return result.Error
}

func FindAll(models interface{}, query interface{}, args ...interface{}) error {
	result := DB.Where(query, args...).Find(models)
	return result.Error
}

func UpdateOne(model interface{}, newValue interface{}) error {
	result := DB.Model(model).Updates(newValue)
	return result.Error
}

func Delete(model interface{}) error {
	result := DB.Delete(model)
	return result.Error
}
