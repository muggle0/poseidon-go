package config

import (
	"github.com/jinzhu/gorm"
	"syscall"
)

var database *Database


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
func initDatabase(*map[interface{}]interface{}) {

}

func getInstance() (*Database,err error) {
	if database==nil {
		return
	}

	return database,nil;
}