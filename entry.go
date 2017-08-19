package deeplearning

import (
	"os"
	log "github.com/cihub/seelog"
	"deeplearning/data"
	"deeplearning/config"
)

type Application struct {
	config  *config.AppConfig
	MovieData *data.MovieData
}

func ServerStart() {
	app := new(Application)
	conf := new(config.AppConfig)
	err := config.LoadConfig("server.conf", conf)
	if err != nil {
		log.Errorf("load config file failed: %v", err)
		os.Exit(1)
	}
	app.config = conf

	app.Run()

	defer log.Flush()
}

func (app *Application) Run() {
	data := data.LoadData(app.config.Files)
	if data == nil {
		log.Error("load data file failed")
		os.Exit(1)
	}

	app.MovieData = data
}

