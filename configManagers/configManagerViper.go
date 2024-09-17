package configManagers

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type configManagerViper struct {
}

func newConfigManagerViper() configManagerViper {
	manager := configManagerViper{}
	manager.setup()
	return manager
}

func (c configManagerViper) setup() {
	// Load initial config
	loadConfig()

	// Watch for changes in the config file
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
		loadConfig()
	})
}

func loadConfig() {

	log.Print("loadConfig function was triggered")

	viper.SetConfigFile("config.yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	viper.AutomaticEnv()

	log.Print("loadConfig function was completed")
}

func (c configManagerViper) GetStr(key string) string {
	val := viper.GetString(key)
	return val
}
