package moviesubject_test

import (
	"github.com/heibizi/go-moviesubject"
	"os"
	"testing"
)

var ms = moviesubject.NewMovieSubject()

func TestMain(m *testing.M) {
	ms.SetTmdbApiParams(moviesubject.TmdbApiParams{
		ApiKey: os.Getenv("GO_MOVIESUBJECT_TMDB_APIKEY"),
	})
	m.Run()
}

func TestAllSubjects(t *testing.T) {
	subjects := ms.AllSubjects()
	t.Log(subjects)
}

func TestDoubanSubjectCollection(t *testing.T) {
	subject := moviesubject.Subject{Code: "movie_showing", Name: "豆瓣正在热映",
		Category: "douban_subject_collection"}
	items, err := ms.Items(subject, 0, 20)
	if err != nil {
		t.Error(err)
	}
	t.Log(items)
}

func TestTmdbTrending(t *testing.T) {
	subject := moviesubject.Subject{Code: "tmdb_movie_trending_week", Name: "TMDB 电影本周趋势",
		Category: "tmdb_trending"}
	items, err := ms.Items(subject, 0, 20)
	if err != nil {
		t.Error(err)
	}
	t.Log(items)
}

func TestTmdbMovie(t *testing.T) {
	subject := moviesubject.Subject{Code: "tmdb_movie_popular", Name: "TMDB 热门电影",
		Category: "tmdb_movie"}
	items, err := ms.Items(subject, 0, 20)
	if err != nil {
		t.Error(err)
	}
	t.Log(items)
}

func TestTmdbTv(t *testing.T) {
	subject := moviesubject.Subject{Code: "tmdb_tv_popular", Name: "TMDB 热门电视剧",
		Category: "tmdb_tv"}
	items, err := ms.Items(subject, 0, 20)
	if err != nil {
		t.Error(err)
	}
	t.Log(items)
}
