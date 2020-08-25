package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

var Properties struct {
	Database Database
}

func InitConfig() {
	dir, _ := os.Getwd()
	join := path.Join(dir, "/config/test.yml")
	bytes, _ := ioutil.ReadFile(join)
	fmt.Printf("读取项目配置，运行环境: %s\n", gin.Mode())
	yaml.Unmarshal(bytes, &Properties)

}
