package weathers

import (
	"WeatherInfo/configManagers"
	"WeatherInfo/weathers/weatherServices"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterEndpoints(g *gin.Engine) {
	g.LoadHTMLFiles("./weathers/htmlTemplates/index.html")

	g.GET("/", serveIndexHtml)
	g.GET("/index", serveIndexHtml)

	g.GET("api/:location/weathers", func(context *gin.Context) {
		getWeatherInfo(context)
	})
}

func serveIndexHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func getWeatherInfo(context *gin.Context) {
	location := context.Param("location")

	if location == "" {
		context.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Location should not be empty"})
		return
	}

	req := weatherServices.GetWeatherRequest{
		Location: location,
	}

	configManager := configManagers.NewConfigManager()
	service := weatherServices.NewWeatherService(configManager)
	res, err := service.GetWeatherInfo(req)

	if err != nil {
		if errors.Is(err, weatherServices.ErrorLocationNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"errorMessage": err.Error()})
		} else if errors.Is(err, weatherServices.ErrorUnauthorizedRequest) {
			context.JSON(http.StatusUnauthorized, gin.H{"errorMessage": err.Error()})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		}

		return
	}

	context.JSON(http.StatusOK, res)
	return
}
