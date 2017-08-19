package data

import "strings"

type BaseMovieInfo struct {
	ID string
	Title string
	Genres []string
}

type MovieItem struct {
	BaseMovieInfo
	Users map[string]*MovieRating
}

type MovieRating struct {
	Rating float32
	Tag    string
}

func DecodeMovieInfo(vals []string) *MovieItem {
	if len(vals) < 3 {
		return nil
	}

	item := new(MovieItem)
	item.ID = vals[0]
	item.Title = vals[1]
	item.Users = make(map[string]*MovieRating)

	item.Genres = strings.Split(vals[2], "|")

	return item
}


func (item *MovieItem) AddRating(userID string, rating float32) {
	if rate := item.Users[userID]; rate != nil {
		rate.Rating = rating
	} else {
		item.Users[userID] = &MovieRating{Rating:rating}
	}
}

func (item *MovieItem) AddTag(userID string, tags string) {
	if len(tags) == 0 {
		return
	}

	if rate := item.Users[userID]; rate != nil {
		rate.Tag = tags
	} else {
		item.Users[userID] = &MovieRating{Tag: tags}
	}
}