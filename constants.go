package moviesubject

type category string

const (
	categoryDoubanDouList           category = "douban_dou_list"           // 豆瓣豆列
	categoryDoubanSubjectCollection category = "douban_subject_collection" // 豆瓣片单
	categoryTmdbTrending            category = "tmdb_trending"             // TMDB 趋势
	categoryTmdbMovie               category = "tmdb_movie"                // TMDB 电影
	categoryTmdbTv                  category = "tmdb_tv"                   // TMDB 电视剧
)

type subjectCode string

const (
	tmdbMoviePopular    subjectCode = "tmdb_movie_popular"
	tmdbMovieNowPlaying subjectCode = "tmdb_movie_now_playing"
	tmdbMovieUpcoming   subjectCode = "tmdb_movie_upcoming"
	tmdbMovieTopRated   subjectCode = "tmdb_movie_top_rated"

	tmdbTvPopular     subjectCode = "tmdb_tv_popular"
	tmdbTvAiringToday subjectCode = "tmdb_tv_airing_today"
	tmdbTvOnTheAir    subjectCode = "tmdb_tv_on_the_air"
	tmdbTvTopRated    subjectCode = "tmdb_tv_top_rated"

	tmdbMovieTrendingWeek subjectCode = "tmdb_movie_trending_week"
	tmdbMovieTrendingDay  subjectCode = "tmdb_movie_trending_day"
	tmdbTvTrendingWeek    subjectCode = "tmdb_tv_trending_week"
	tmdbTvTrendingDay     subjectCode = "tmdb_tv_trending_day"
)

var (
	subjects = []Subject{
		{Code: "movie_showing", Name: "豆瓣正在热映", Category: string(categoryDoubanSubjectCollection)},
		{Code: "movie_hot_gaia", Name: "豆瓣热门电影", Category: string(categoryDoubanSubjectCollection)},
		{Code: "tv_hot", Name: "豆瓣热门剧集", Category: string(categoryDoubanSubjectCollection)},
		{Code: "tv_animation", Name: "豆瓣热门动画", Category: string(categoryDoubanSubjectCollection)},
		{Code: "movie_soon", Name: "豆瓣即将上映", Category: string(categoryDoubanSubjectCollection)},
		{Code: "movie_top250", Name: "豆瓣电影Top250", Category: string(categoryDoubanSubjectCollection)},
		{Code: "show_hot", Name: "豆瓣热门综艺", Category: string(categoryDoubanSubjectCollection)},
		{Code: "tv_chinese_best_weekly", Name: "豆瓣华语口碑剧集周榜", Category: string(categoryDoubanSubjectCollection)},
		{Code: "tv_global_best_weekly", Name: "豆瓣全球口碑剧集周榜", Category: string(categoryDoubanSubjectCollection)},

		{Code: string(tmdbMovieTrendingWeek), Name: "TMDB 电影本周趋势", Category: string(categoryTmdbTrending)},
		{Code: string(tmdbMovieTrendingDay), Name: "TMDB 电影本日趋势", Category: string(categoryTmdbTrending)},

		{Code: string(tmdbTvTrendingWeek), Name: "TMDB 电视剧本周趋势", Category: string(categoryTmdbTrending)},
		{Code: string(tmdbTvTrendingDay), Name: "TMDB 电视剧本日趋势", Category: string(categoryTmdbTrending)},

		{Code: string(tmdbMoviePopular), Name: "TMDB 热门电影", Category: string(categoryTmdbMovie)},
		{Code: string(tmdbMovieNowPlaying), Name: "TMDB 最新电影", Category: string(categoryTmdbMovie)},
		{Code: string(tmdbMovieUpcoming), Name: "TMDB 即将上映电影", Category: string(categoryTmdbMovie)},
		{Code: string(tmdbMovieTopRated), Name: "TMDB 高分电影", Category: string(categoryTmdbMovie)},

		{Code: string(tmdbTvPopular), Name: "TMDB 热门电视剧", Category: string(categoryTmdbTv)},
		{Code: string(tmdbTvOnTheAir), Name: "TMDB 最新电视剧", Category: string(categoryTmdbTv)},
		{Code: string(tmdbTvAiringToday), Name: "TMDB 今日播出电视剧", Category: string(categoryTmdbTv)},
		{Code: string(tmdbTvTopRated), Name: "TMDB 高分电视剧", Category: string(categoryTmdbTv)},

		{Code: "movie_scifi", Name: "豆瓣高分经典科幻片", Category: string(categoryDoubanSubjectCollection)},
		{Code: "movie_comedy", Name: "豆瓣高分经典喜剧片", Category: string(categoryDoubanSubjectCollection)},
		{Code: "movie_action", Name: "豆瓣高分经典动作片", Category: string(categoryDoubanSubjectCollection)},
		{Code: "movie_love", Name: "豆瓣高分经典爱情片", Category: string(categoryDoubanSubjectCollection)},
		{Code: "tv_domestic", Name: "豆瓣国产剧", Category: string(categoryDoubanSubjectCollection)},
		{Code: "tv_american", Name: "豆瓣美剧", Category: string(categoryDoubanSubjectCollection)},
		{Code: "tv_japanese", Name: "豆瓣日剧", Category: string(categoryDoubanSubjectCollection)},
		{Code: "tv_korean", Name: "豆瓣韩剧", Category: string(categoryDoubanSubjectCollection)},
		{Code: "tv_variety_show", Name: "豆瓣综艺", Category: string(categoryDoubanSubjectCollection)},
		{Code: "show_domestic", Name: "豆瓣国内综艺", Category: string(categoryDoubanSubjectCollection)},
		{Code: "show_foreign", Name: "豆瓣国外综艺", Category: string(categoryDoubanSubjectCollection)},
	}
)
