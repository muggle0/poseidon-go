package util

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

func YamlValue(key string, filepath string) interface{} {
	bytes, _ := ioutil.ReadFile(filepath)
	var config map[interface{}]interface{}
	yaml.Unmarshal(bytes, &config)
	config = config[gin.ReleaseMode].(map[interface{}]interface{})
	keys := strings.Split(key, ".")
	length := len(keys)
	if length == 1 {
		return config[key]
	} else {
		var value interface{}
		for _, v := range keys {
			if value == nil {
				value = config[v]
			} else {
				value = value.(map[interface{}]interface{})[v]
			}
		}
		return value
	}
}
