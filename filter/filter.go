package filter

import (
	"deepmovie/filter/movie_filter"
	"deepmovie/data"
	"github.com/liangdas/mqant/log"
)

type Filter struct {
	MovieFilter *movie_filter.Filter
}

func NewFilter(data *data.MovieData) *Filter{
	filter := new(Filter)

	filter.MovieFilter = movie_filter.NewMovieFilter(data)

	return filter
}


func (filter *Filter) Run(userID string, count int, callBack chan<- []string) {
	movieChan := make(chan []string)
//	userChan := make(chan []string)

	go filter.MovieFilter.Run(userID, movieChan)

	var movieRecomms []string
	var userRecomms  []string
	for {
		select {
		//case movies, ok := <- userChan:
		//	if !ok {
		//		userChan = nil
		//	} else {
		//		userRecomms = movies
		//	}
		case movies, ok := <- movieChan:
			if !ok {
				movieChan = nil
			} else {
				movieRecomms = movies
				movieChan = nil
			}
		}
		if movieChan == nil {
			break
		}
	}

	result := filter.Remix(userRecomms, movieRecomms, count)

	callBack <- result
}


//1.Remix Result
func (filter *Filter) Remix(userRecomms []string, movieRecomms []string, count int) []string {
	log.Debug("基于用户的推荐数：%d", len(userRecomms))
	log.Debug("基于电影的推荐数：%d", len(movieRecomms))

	var result []string = make([]string, 0)
	u_count := len(userRecomms)
	m_count := len(movieRecomms)
	if u_count == 0 {
		if m_count < count {
			result = append(result, movieRecomms...)
			result = append(result, filter.MovieFilter.GetRandomMovie(count - m_count)...)
		} else {
			result = append(result, movieRecomms[:count]...)
		}

		return result
	}

	if m_count == 0 {
		if u_count < count {
			result = append(result, movieRecomms...)
			result = append(result, filter.MovieFilter.GetRandomMovie(count - u_count)...)
		} else {
			result = append(result, movieRecomms[:count]...)
		}

		return result
	}

	if m_count + u_count > count {
		// 优先查询两种算法结果的重叠部分
		//　重叠结果数量不足，先从基于电影的推荐里面取５个备用, 未取到５个则从另一个算法里面取
		var back_result []string = make([]string, count)
		if len(back_result) < count {
			if count <= m_count {
				back_result = movieRecomms[:5]
			} else {
				back_result = movieRecomms[:m_count]
			}
		}

		if len(back_result) < count {
			back_result = userRecomms[:count-len(back_result)]
		}

		for _, u_val := range userRecomms {
			for _, m_val := range movieRecomms {
				if u_val == m_val {
					result = append(result, m_val)
				}
			}
		}

		//复杂度小于20
		if len(result) < count {
			for _, back := range back_result {
				var find bool = false
				for _, have := range result {
					if have == back {
						find = true
						break
					}
				}
				if !find {
					result = append(result, back)
				}
				if len(result) == count {
					break
				}
			}
		} else {
			return result[:5]
		}
	} else {
		result = append(result, movieRecomms...)
		result = append(result, userRecomms...)

		result = append(result, filter.MovieFilter.GetRandomMovie(count-len(result))...)
	}

	return result
}

