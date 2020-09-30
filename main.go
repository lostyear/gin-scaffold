package main

import (
	"flag"

	"github.com/lostyear/gin-scaffold/commons"
	"github.com/lostyear/gin-scaffold/http"
)

var (
	configFile string
)

func Init() {
	flag.StringVar(&configFile, "c", "config.json", "use config file")
	flag.Parse()
}

func main() {
	// db.Debug().AutoMigrate(&model.PodInfo{})
	commons.Init(configFile)
	cfg := commons.GetConfig()
	http.Start(cfg.HTTP)
	commons.CloseAll()
}
