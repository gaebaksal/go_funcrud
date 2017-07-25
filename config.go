package main

type Config struct {
	DB *_DBConfig
}

type _DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &_DBConfig{
			Dialect:  "mysql",
			Username: "guest",
			Password: "roqkrtkf!",
			Name:     "todoapp",
			Charset:  "utf8",
		},
	}
}
