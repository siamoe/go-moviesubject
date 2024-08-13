package moviesubject

import (
	"fmt"
	"strconv"
)

func (s *tmdbService) trendingItems(subject Subject, pageIndex int) (Result, error) {
	mediaType := ""
	timeWindow := ""
	if subject.Code == string(tmdbMovieTrendingWeek) {
		mediaType = "movie"
		timeWindow = "week"
	} else if subject.Code == string(tmdbMovieTrendingDay) {
		mediaType = "movie"
		timeWindow = "day"
	} else if subject.Code == string(tmdbTvTrendingWeek) {
		mediaType = "tv"
		timeWindow = "week"
	} else if subject.Code == string(tmdbTvTrendingDay) {
		mediaType = "tv"
		timeWindow = "day"
	}
	options := newTmdbOptions(s.params)
	options["page"] = strconv.Itoa(pageIndex + 1)
	t, err := s.client.GetTrending(mediaType, timeWindow, options)
	if err != nil {
		return Result{}, fmt.Errorf("获取 tmdb 趋势异常: %v", err)
	}
	var list []Media
	for _, item := range t.Results {
		list = append(list, Media{
			Id:       strconv.FormatInt(item.ID, 10),
			Title:    item.Title,
			Type:     mediaType,
			Year:     getTmdbYearFrom(item.ReleaseDate),
			Vote:     float64(item.VoteAverage),
			Image:    item.PosterPath,
			Overview: item.Overview,
		})
	}
	return Result{
		PageIndex: pageIndex,
		PageSize:  20,
		Total:     t.TotalResults,
		List:      list,
	}, nil
}
