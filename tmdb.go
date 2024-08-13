package moviesubject

import (
	"fmt"
	tmdb "github.com/cyruzin/golang-tmdb"
	"net/http"
	"strconv"
	"time"
)

type tmdbService struct {
	client *tmdb.Client
	params TmdbApiParams
}

func newTmdbClient(params TmdbApiParams) *tmdb.Client {
	// apikey 为空才会报错，直接忽略
	tmdbClient, _ := tmdb.Init(params.ApiKey)
	tmdbClient.SetClientAutoRetry()
	if params.CustomBaseURL != "" {
		tmdbClient.SetCustomBaseURL(params.CustomBaseURL)
	}
	if params.Timeout == 0 {
		params.Timeout = time.Second * 60
	}
	if params.MaxIdleConn == 0 {
		params.MaxIdleConn = 10
	}
	if params.IdleConnTimeout == 0 {
		params.IdleConnTimeout = time.Second * 60
	}
	customClient := http.Client{
		Timeout: params.Timeout,
		Transport: &http.Transport{
			MaxIdleConns:    params.MaxIdleConn,
			IdleConnTimeout: params.IdleConnTimeout,
		},
	}
	tmdbClient.SetClientConfig(customClient)
	return tmdbClient
}

func newTmdbOptions(params TmdbApiParams) map[string]string {
	options := make(map[string]string)
	if params.Language != "" {
		options["language"] = params.Language
	} else {
		options["language"] = "zh"
	}
	if params.Region != "" {
		options["region"] = params.Region
	} else {
		options["region"] = "CN"
	}
	options["include_adult"] = strconv.FormatBool(params.IncludeAdult)
	return options
}

func getTmdbYearFrom(date string) int {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println(err)
		return 2006
	}
	return t.Year()
}
