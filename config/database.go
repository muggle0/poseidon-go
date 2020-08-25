package config

import (
	"github.com/jinzhu/gorm"
)

var database *Database

type Database struct {
	Datatype string `yaml:"datatype"`
	Hostname string `yaml:"hostname"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Prefix   string `yaml:"prefix"`
	DB       *gorm.DB
}

func (database *Database) Connection() {
	db, _ := gorm.Open(database.Datatype, database.Username+":"+database.Password+"@/"+database.Database+"?charset=utf8&parseTime=True&loc=Local")
	//全局禁用表复数
	database.DB = db
	db.SingularTable(true)
	db.Close()
	// 设置连接池
}

func GetInstance() *Database {
	return database
}

func (database *Database) Close() {
	database.DB.Close()
}
