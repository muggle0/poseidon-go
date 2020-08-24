package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

var properties map[interface{}]interface{}

func InitConfig() {
	dir, _ := os.Getwd()
	join := path.Join(dir, "/config/test.yml")
	bytes, _ := ioutil.ReadFile(join)
	fmt.Printf("读取项目配置，运行环境: %s\n", gin.Mode())
	yaml.Unmarshal(bytes, &properties)
	properties = properties[gin.Mode()].(map[interface{}]interface{})
	fmt.Println(properties)
	initDatabase(&properties)
}
