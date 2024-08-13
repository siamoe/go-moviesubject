package moviesubject

import (
	"errors"
	"github.com/heibizi/go-douban"
)

type MovieSubject struct {
	doubanService *doubanService
	tmdbService   *tmdbService
}

func NewMovieSubject() *MovieSubject {
	ms := MovieSubject{}
	ms.doubanService = &doubanService{douban.NewApiClient()}
	return &ms
}

func (ms *MovieSubject) SetTmdbApiParams(params TmdbApiParams) {
	if params.ApiKey != "" {
		ms.tmdbService = &tmdbService{
			client: newTmdbClient(params),
			params: params,
		}
	} else {
		ms.tmdbService = nil
	}
}

func (ms *MovieSubject) AllSubjects() []Subject {
	return subjects
}

var serviceRegistry = map[string]func(ms *MovieSubject, subject Subject, pageIndex int, pageSize int) (Result, error){
	string(categoryDoubanDouList): func(ms *MovieSubject, subject Subject, pageIndex int, pageSize int) (Result, error) {
		return ms.doubanService.douListItems(subject, pageIndex, pageSize)
	},
	string(categoryDoubanSubjectCollection): func(ms *MovieSubject, subject Subject, pageIndex int, pageSize int) (Result, error) {
		return ms.doubanService.subjectCollectionItems(subject, pageIndex, pageSize)
	},
	string(categoryTmdbTrending): func(ms *MovieSubject, subject Subject, pageIndex int, pageSize int) (Result, error) {
		return ms.tmdbService.trendingItems(subject, pageIndex)
	},
	string(categoryTmdbMovie): func(ms *MovieSubject, subject Subject, pageIndex int, pageSize int) (Result, error) {
		return ms.tmdbService.movieItems(subject, pageIndex)
	},
	string(categoryTmdbTv): func(ms *MovieSubject, subject Subject, pageIndex int, pageSize int) (Result, error) {
		return ms.tmdbService.tvItems(subject, pageIndex)
	},
}

func (ms *MovieSubject) Items(subject Subject, pageIndex int, pageSize int) (Result, error) {
	if service, ok := serviceRegistry[subject.Category]; ok {
		return service(ms, subject, pageIndex, pageSize)
	}
	return Result{}, errors.New("category not found")
}
