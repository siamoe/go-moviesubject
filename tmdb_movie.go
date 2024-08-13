package moviesubject

import (
	"fmt"
	"strconv"
)

func (s *tmdbService) movieItems(subject Subject, pageIndex int) (Result, error) {
	if subject.Code == string(tmdbMoviePopular) {
		return s.moviePopulars(pageIndex)
	} else if subject.Code == string(tmdbMovieNowPlaying) {
		return s.movieNowPlayings(pageIndex)
	} else if subject.Code == string(tmdbMovieTopRated) {
		return s.movieTopRated(pageIndex)
	} else if subject.Code == string(tmdbMovieUpcoming) {
		return s.movieUpcoming(pageIndex)
	}
	return Result{}, fmt.Errorf("unknown tmdb movie subject code: %s", subject.Code)
}

func (s *tmdbService) moviePopulars(pageIndex int) (Result, error) {
	options := newTmdbOptions(s.params)
	options["page"] = strconv.Itoa(pageIndex + 1)
	t, err := s.client.GetMoviePopular(options)
	if err != nil {
		return Result{}, fmt.Errorf("get tmdb movie popular err: %v", err)
	}
	var list []Media
	for _, item := range t.Results {
		list = append(list, Media{
			Id:       strconv.FormatInt(item.ID, 10),
			Title:    item.Title,
			Type:     "movie",
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

func (s *tmdbService) movieNowPlayings(pageIndex int) (Result, error) {
	options := newTmdbOptions(s.params)
	options["page"] = strconv.Itoa(pageIndex + 1)
	t, err := s.client.GetMovieNowPlaying(options)
	if err != nil {
		return Result{}, fmt.Errorf("get tmdb movie now playing err: %v", err)
	}
	var list []Media
	for _, item := range t.Results {
		list = append(list, Media{
			Id:       strconv.FormatInt(item.ID, 10),
			Title:    item.Title,
			Type:     "movie",
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

func (s *tmdbService) movieTopRated(pageIndex int) (Result, error) {
	options := newTmdbOptions(s.params)
	options["page"] = strconv.Itoa(pageIndex + 1)
	t, err := s.client.GetMovieTopRated(options)
	if err != nil {
		return Result{}, fmt.Errorf("get tmdb movie top rated err: %v", err)
	}
	var list []Media
	for _, item := range t.Results {
		list = append(list, Media{
			Id:       strconv.FormatInt(item.ID, 10),
			Title:    item.Title,
			Type:     "movie",
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

func (s *tmdbService) movieUpcoming(pageIndex int) (Result, error) {
	options := newTmdbOptions(s.params)
	options["page"] = strconv.Itoa(pageIndex + 1)
	t, err := s.client.GetMovieUpcoming(options)
	if err != nil {
		return Result{}, fmt.Errorf("get tmdb movie upcoming err: %v", err)
	}
	var list []Media
	for _, item := range t.Results {
		list = append(list, Media{
			Id:       strconv.FormatInt(item.ID, 10),
			Title:    item.Title,
			Type:     "movie",
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
