package main

import (
	"WeatherInfo/configManagers"
	"WeatherInfo/weathers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Println("main function was triggered")

	configManager := configManagers.NewConfigManager()

	port := configManager.GetStr("PORT")

	g := gin.Default()
	g.Use(gin.Recovery())

	weathers.RegisterEndpoints(g)

	log.Fatal(g.Run(":" + port))
}
