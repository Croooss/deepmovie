package data

import (
	"os"
	"encoding/csv"
	log "github.com/cihub/seelog"
	"deeplearning/config"
)

const (
	FinishLoadTag = iota
	FinishLoadMovie
	FinishLoadRating
)

type MovieData struct {
	Movies []*MovieInfo
	Ratings map[string][]*MovieRating
	Tags  map[string][]*MovieTag
}

func LoadData(files *config.DataFile) *MovieData {
	if files == nil {
		return nil
	}
	loadChan := make(chan int, 3)

	log.Debug("start load data")
	movieData := new(MovieData)
	movieData.Movies = make([]*MovieInfo, 0)
	movieData.Ratings = make(map[string][]*MovieRating)
	movieData.Tags = make(map[string][]*MovieTag)

	go movieData.loadMovieTag(files.Tag, loadChan)
	go movieData.loadMovieInfo(files.Movie, loadChan)
	go movieData.loadMovieRating(files.Rating, loadChan)

	var count int
	for {
		select {
		case rd, ok := <- loadChan:
			if !ok {
				loadChan = nil
			} else {
				count ++
				if rd == FinishLoadTag {
					log.Debug("load tag finished")
				} else if rd == FinishLoadMovie {
					log.Debug("load movie finished")
				} else if rd == FinishLoadRating {
					log.Debug("load rating finished")
				}
				if count == 3{
					loadChan = nil
				}
			}
		}
		if loadChan == nil {
			break
		}
	}

	log.Debug("load data finished")
	return movieData
}


func (data *MovieData) loadMovieInfo(fileName string, readChan chan<- int){
	log.Debugf("start load file: %s \n", fileName)
	dats, err := readDataFile(fileName)
	if err != nil {
		log.Error("invalid csv file (%s) \n", fileName)
		os.Exit(1)
	}

	for _, vals := range dats {
		info := DecodeMovieInfo(vals)
		data.Movies = append(data.Movies, info)
	}

	readChan <- FinishLoadMovie
}

func (data *MovieData) loadMovieTag(fileName string, readChan chan<- int) {
	log.Debugf("start load file: %s \n", fileName)
	dats, err := readDataFile(fileName)
	if err != nil {
		log.Error("invalid csv file(%s)\n", fileName)
		os.Exit(1)
	}

	for _, vals := range dats {
		tag := DecodeMovieTag(vals)
		data.Tags[tag.UserID] = append(data.Tags[tag.UserID], tag)
	}

	readChan <- FinishLoadTag
}

func (data *MovieData) loadMovieRating(fileName string, readChan chan<- int) {
	log.Debugf("start load file: %s \n", fileName)
	dats, err := readDataFile(fileName)
	if err != nil {
		log.Error("invalid csv file(%s)\n", fileName)
		os.Exit(1)
	}

	for _, vals := range dats {
		rate := DecodeMovieRating(vals)
		data.Ratings[rate.UserID] = append(data.Ratings[rate.UserID], rate)
	}

	readChan <- FinishLoadRating
}

func readDataFile(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	return reader.ReadAll()
}
