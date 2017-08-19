package user_filter

import "deepmovie/data"

type Filter struct {

}

func NewUserFilter(data *data.MovieData) *Filter{
	filter := new(Filter)

	return filter
}

func (filter *Filter) Run(userid string, callBack chan<- []string) {
	//TODO
}