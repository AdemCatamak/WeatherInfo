package configManagers

type ConfigManager interface {
	GetStr(key string) string
}

func NewConfigManager() ConfigManager {
	manager := newConfigManagerViper()
	return manager
}
