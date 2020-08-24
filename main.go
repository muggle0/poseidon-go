package main

import "poseidon-go/config"

func init() {
	config.InitConfig()
}
func main() {

	/*log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")*/

	/*	gin.SetMode(gin.DebugMode)
		new(config.Database).GetConnect()
		defer config.DB.Close()
		engine := gin.New()
		(&route.Route{Engine: engine}).Run()
		engine.Run(":8080")*/
}
