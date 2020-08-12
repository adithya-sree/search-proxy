package config

type Config struct {
	AppConfig *AppConfig
	DeezerConfig *DeezerConfig
}

type AppConfig struct {
	Port         string
	LogDirectory string
	Username     string
	Password     string
}

type DeezerConfig struct {
	ApiFqdn       string
	ApiHostHeader string
	ApiHost       string
	ApiKeyHeader  string
	ApiKey		  string
}

func GetConfig() *Config {
	return &Config{
		AppConfig: &AppConfig{
			Port:         "8888",
			LogDirectory: "/var/log/search-api.log",
			Username:     "test",
			Password:     "test",
		},
		DeezerConfig: &DeezerConfig{
			ApiFqdn:       "https://deezerdevs-deezer.p.rapidapi.com/search?q=",
			ApiHostHeader: "X-RapidAPI-Host",
			ApiHost:       "deezerdevs-deezer.p.rapidapi.com",
			ApiKeyHeader:  "X-RapidAPI-Key",
			ApiKey:        "",
		},
	}
}
