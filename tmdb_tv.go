package moviesubject

import (
	"fmt"
	"strconv"
)

func (s *tmdbService) tvItems(subject Subject, pageIndex int) (Result, error) {
	if subject.Code == string(tmdbTvPopular) {
		return s.tvPopulars(pageIndex)
	} else if subject.Code == string(tmdbTvTopRated) {
		return s.tvTopRated(pageIndex)
	} else if subject.Code == string(tmdbTvAiringToday) {
		return s.tvAiringToday(pageIndex)
	} else if subject.Code == string(tmdbTvOnTheAir) {
		return s.tvOnTheAir(pageIndex)
	}
	return Result{}, fmt.Errorf("unknown tmdb tv subject code: %s", subject.Code)
}

func (s *tmdbService) tvPopulars(pageIndex int) (Result, error) {
	options := newTmdbOptions(s.params)
	options["page"] = strconv.Itoa(pageIndex + 1)
	t, err := s.client.GetTVPopular(options)
	if err != nil {
		return Result{}, fmt.Errorf("get tmdb tv popular err: %v", err)
	}
	var list []Media
	for _, item := range t.Results {
		list = append(list, Media{
			Id:       strconv.FormatInt(item.ID, 10),
			Title:    item.Name,
			Type:     "tv",
			Year:     getTmdbYearFrom(item.FirstAirDate),
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

func (s *tmdbService) tvTopRated(pageIndex int) (Result, error) {
	options := newTmdbOptions(s.params)
	options["page"] = strconv.Itoa(pageIndex + 1)
	t, err := s.client.GetTVTopRated(options)
	if err != nil {
		return Result{}, fmt.Errorf("get tmdb tv top rated err: %v", err)
	}
	var list []Media
	for _, item := range t.Results {
		list = append(list, Media{
			Id:       strconv.FormatInt(item.ID, 10),
			Title:    item.Name,
			Type:     "tv",
			Year:     getTmdbYearFrom(item.FirstAirDate),
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

func (s *tmdbService) tvOnTheAir(pageIndex int) (Result, error) {
	options := newTmdbOptions(s.params)
	options["page"] = strconv.Itoa(pageIndex + 1)
	t, err := s.client.GetTVOnTheAir(options)
	if err != nil {
		return Result{}, fmt.Errorf("get tmdb tv on the air err: %v", err)
	}
	var list []Media
	for _, item := range t.Results {
		list = append(list, Media{
			Id:       strconv.FormatInt(item.ID, 10),
			Title:    item.Name,
			Type:     "tv",
			Year:     getTmdbYearFrom(item.FirstAirDate),
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

func (s *tmdbService) tvAiringToday(pageIndex int) (Result, error) {
	options := newTmdbOptions(s.params)
	options["page"] = strconv.Itoa(pageIndex + 1)
	t, err := s.client.GetTVAiringToday(options)
	if err != nil {
		return Result{}, fmt.Errorf("get tmdb tv on the air err: %v", err)
	}
	var list []Media
	for _, item := range t.Results {
		list = append(list, Media{
			Id:       strconv.FormatInt(item.ID, 10),
			Title:    item.Name,
			Type:     "tv",
			Year:     getTmdbYearFrom(item.FirstAirDate),
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
