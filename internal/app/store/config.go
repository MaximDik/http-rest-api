package store

type Config struct {
	DatabaseURL string `toml:"database_url"` //connect string to db.
}

func NewConfig() *Config {
	return &Config{}
}
