package repo

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// init sqlite

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

	DB = gormDB.Debug()
	err = DB.AutoMigrate(&Post{}, &Tag{}, &User{})
	if err != nil {
		log.Panicln("DB AutoMigrate Fail")
	}
}

// create

func Create(model interface{}) error {
	result := DB.Create(model)
	return result.Error
}

// retrieve

func FindOne(model interface{}, conditions ...interface{}) error {
	result := DB.First(model, conditions...)
	return result.Error
}

type Condition struct {
	Query         interface{}
	Args          []interface{}
	Orders        []interface{}
	Offset, Limit int
}

func FindAll(models interface{}, condition Condition) error {
	tx := DB.Where(condition.Query, condition.Args...)
	for _, order := range condition.Orders {
		tx = tx.Order(order)
	}

	if condition.Offset != 0 {
		tx = tx.Offset(condition.Offset)
	}
	if condition.Limit != 0 {
		tx = tx.Limit(condition.Limit)
	}
	result := tx.Find(models)
	return result.Error
}

// update

func UpdateOne(model interface{}, newValue interface{}) error {
	result := DB.Model(model).Updates(newValue)
	return result.Error
}

// delete

func Delete(model interface{}) error {
	result := DB.Delete(model)
	return result.Error
}
