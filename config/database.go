package config

import (
	"github.com/jinzhu/gorm"
	"sync"
)

var database *Database

var mu sync.Mutex

type Database struct {
	Datatype string `yaml:"datatype"`
	Hostname string `yaml:"hostname"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Prefix   string `yaml:"prefix"`
}

func (database *Database) connection() {
	db, _ := gorm.Open(database.Datatype, database.Username+":"+database.Password+"@/"+database.Database+"?charset=utf8&parseTime=True&loc=Local")
	//全局禁用表复数
	db.SingularTable(true)
	// 设置连接池
}

func GetInstance() *Database {
	if database == nil {
		mu.Lock()
		if database == nil {
			// 初始化
		}
		mu.Unlock()

	}
	return database
}
