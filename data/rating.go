package data

import "strconv"

type BaseMovieRating struct {
	UserID string
	MovieID string
	Rating float32

	//时间戳不处理
	//timestamp int64
}

type RatingData struct {

}

func DecodeMovieRating(vals []string) *BaseMovieRating {
	if len(vals) < 3 {
		return nil
	}

	rate := new(BaseMovieRating)
	rate.UserID = vals[0]
	rate.MovieID = vals[1]
	if val, err := strconv.ParseFloat(vals[2], 32); err != nil {
		return nil
	} else {
		rate.Rating = float32(val)
	}

	return rate
}

