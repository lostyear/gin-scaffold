package commons

import (
	log "github.com/lostyear/go-toolkits/logger"
	"github.com/lostyear/go-toolkits/storage"
	"gorm.io/gorm"
)

var cfg *Config
var db *gorm.DB

func Init(configFile string) {
	// 读取配置文件
	cfg = LoadConfig(configFile)
	// 初始化日志
	log.Init(cfg.Log)
	// 初始化数据库链接
	db = storage.InitDatabase(cfg.DB)
}

func CloseAll() {
	storage.Close(db)
	log.Close()
}

func GetConfig() *Config {
	return cfg
}

func GetDatabaseConn() *gorm.DB {
	return db
}
