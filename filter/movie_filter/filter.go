package movie_filter

import (
	"deepmovie/data"
	"time"
	log "github.com/cihub/seelog"
)

type Filter struct {
	Rates map[string]*data.MovieItem
	UnRates map[string]*data.MovieItem
}

func NewMovieFilter(mdat *data.MovieData) *Filter {
	if mdat == nil {
		return nil
	}

	filter := new(Filter)
	filter.Rates = make(map[string]*data.MovieItem)
	filter.UnRates = make(map[string]*data.MovieItem)

	//TODO
	//效率可能很低
	log.Debugf("movie_filter init start : %s", time.Now().String())
	for _, vals := range mdat.Ratings {
		for _, val := range vals {
			if item := mdat.Movies[val.MovieID]; item != nil {
				item.AddRating(val.UserID, val.Rating)
				filter.Rates[val.MovieID] = item
			}
		}
	}

	for _, vals := range mdat.Tags {
		for _, val := range vals {
			if item := mdat.Movies[val.MovieID]; item != nil {
				item.AddTag(val.UserID, val.Tag)
				filter.Rates[val.MovieID] = item
			}
		}
	}

	for _, val := range mdat.Movies {
		if item := filter.Rates[val.ID]; item == nil {
			filter.UnRates[val.ID] = val
		}
	}

	log.Debugf("movie_filter init finish : %s", time.Now().String())

	return filter
}

func (filter *Filter) Run(userid string, callBack chan<- []string) {
	var count int = 10
	result := make([]string, 0)

	//todo

	callBack <- result
}

func (filter *Filter) GetRandomMovie(count int) []string {
	var result []string = make([]string, 0)

	for _, val := range filter.UnRates {
		result = append(result, val.ID)
		if len(result) == count {
			break
		}
	}

	return result
}