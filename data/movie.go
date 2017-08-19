package data

import "strings"

type MovieInfo struct {
	ID string
	Title string
	Genres []string
}

func DecodeMovieInfo(vals []string) *MovieInfo {
	if len(vals) < 3 {
		return nil
	}

	info := new(MovieInfo)
	info.ID = vals[0]
	info.Title = vals[1]

	info.Genres = strings.Split(vals[2], "|")

	return info
}
