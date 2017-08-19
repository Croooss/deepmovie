package data

type MovieTag struct {
	UserID string
	MovieID string
	Tag  string

	//时间戳暂时不处理
	//timestamp int64
}

func DecodeMovieTag(vals []string) *MovieTag{
	tag := new(MovieTag)

	if len(vals) < 3 {
		return nil
	}

	tag.UserID = vals[0]
	tag.MovieID = vals[1]
	tag.Tag = vals[2]

	return tag
}