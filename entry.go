package deepmovie
import (
	"os"
	log "github.com/cihub/seelog"
	"deepmovie/data"
	"deepmovie/config"
	"deepmovie/filter"
)

type Application struct {
	config  *config.AppConfig
	MovieData *data.MovieData
	filter  *filter.Filter
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
	app.filter = filter.NewFilter(app.MovieData)

	result := make(chan []string)
	go app.filter.Run(app.config.UserID, app.config.Count, result)

	for {
		select {
		case ret, ok := <-result:
			if !ok {
				result = nil
			} else {
				if len(ret) != app.config.Count {
					log.Debug("not enough movies recommended\n")
				}

				for _, val := range ret {
					log.Debug("recommending user(%s) movie: %s\n", app.config.UserID, val)
				}
			}
		}
		if result == nil {
			break
		}
	}

	log.Debug("recommending end!")
	log.Flush()
}

