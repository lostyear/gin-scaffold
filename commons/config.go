package commons

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/lostyear/go-toolkits/http"
	llog "github.com/lostyear/go-toolkits/logger"
	"github.com/lostyear/go-toolkits/storage"
	"gopkg.in/yaml.v2"
)

const (
	DB_LOG_FILENAME   = "db.log"
	HTTP_LOG_FILENAME = "HTTP.log"
)

type Config struct {
	DB   storage.Config
	HTTP http.Config
	Log  llog.Config
}

func LoadConfig(path string) *Config {
	var cfg Config
	if _, err := os.Stat(path); err != nil {
		cfg = loadDefaultConfig()
	} else {
		cfg = loadConfigFile(path)
	}
	fillLogConfigs(&cfg)
	return &cfg
}

func fillLogConfigs(cfg *Config) {
	var dir string
	if strings.HasSuffix(cfg.Log.Path, "/") {
		dir = cfg.Log.Path
	} else {
		dir = cfg.Log.Path + "/"
	}

	cfg.DB.LogPath = dir + DB_LOG_FILENAME
	cfg.DB.LogLevel = cfg.Log.Level
	cfg.DB.LogMaxDays = cfg.Log.KeepDays
	cfg.DB.LogRotationHours = 1

	cfg.HTTP.LogPath = dir + HTTP_LOG_FILENAME
	cfg.HTTP.LogMaxDays = cfg.Log.KeepDays
	cfg.HTTP.LogRotationHours = 1
}

func loadConfigFile(path string) Config {
	switch strings.ToLower(path[len(path)-4:]) {
	case "json":
		return loadJSONConfig(path)
	case "toml":
		return loadTomlConfig(path)
	case "yaml":
		return loadYamlConfig(path)
	}
	log.Fatalf("unknown config file format: %s\n", path)
	return Config{}
}

func loadJSONConfig(path string) Config {
	var cfg Config
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("read config file failed! Error: %s\n", err)
	}
	if err := json.Unmarshal(b, &cfg); err != nil {
		log.Fatal("decode config file failed! Error: %s\n", err)
	}
	return cfg
}

func loadTomlConfig(path string) Config {
	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		log.Fatal("decode config file failed! Error: %s\n", err)
	}
	return cfg
}

func loadYamlConfig(path string) Config {
	var cfg Config
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("read config file failed! Error: %s\n", err)
	}
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		log.Fatal("decode config file failed! Error: %s\n", err)
	}
	return cfg
}

func loadDefaultConfig() Config {
	return Config{
		Log: llog.Config{
			Path:     "log",
			Level:    "info",
			KeepDays: 7,
		},
		DB: storage.Config{
			// mysql dsn "root:dbpwd@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
			Type:      "sqlite",
			WriterDSN: "db.sqlite",

			LogSlowMicroSeconds:  1000,
			TimeoutMilliseSecond: 10,
			MaxOpenConns:         100,
			MaxIdleConns:         10,
			ConnMaxLifeSeconds:   600,
		},
		HTTP: http.Config{
			Listen: ":8000",

			HTTPTimeoutMilliseSecond:  1000,
			ReadTimeoutMilliseSecond:  10000,
			WriteTimeoutMilliseSecond: 10000,
		},
	}
}
