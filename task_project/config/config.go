package config

type Config struct {
	PostgresHost string
	PostgresUser string
	PostgresDatabase string 
	PostgresPassword string
	PostgresPort string
}

func Load() Config {
	var cfg Config

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "jahongir2"
	cfg.PostgresDatabase = "pj"
	cfg.PostgresPassword = "1005"
	cfg.PostgresPort = "5432"

	return cfg
}

