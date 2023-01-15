package config

// Config config struct
type Config struct {
	server   ServerConfigInterface
	database DatabaseConfigInterface
	redis    RedisConfiginterface
}

// Interface config interface
type Interface interface {
	Server() ServerConfigInterface
	Database() DatabaseConfigInterface
	Redis() RedisConfiginterface
}

// Initialize initialize config
func Initialize() Interface {
	return &Config{
		server:   NewServerConfig(),
		database: NewDatabaseConfig(),
		redis:    NewRedisConfig(),
	}
}

// Server get server config
func (config *Config) Server() ServerConfigInterface {
	return config.server
}

// Database get database config
func (config *Config) Database() DatabaseConfigInterface {
	return config.database
}

// Redis get redis cofig
func (config *Config) Redis() RedisConfiginterface {
	return config.redis
}
