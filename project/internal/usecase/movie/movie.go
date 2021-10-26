package movie

import (
	"context"
	"encoding/json"

	"github.com/mfsyahrz/submission/project/internal/domain/movie"
	sa "github.com/mfsyahrz/submission/project/internal/domain/search_activity"
	"github.com/mfsyahrz/submission/project/internal/infrastructure/omdb"
)

const (
	GetMovieActivity    = "GetMovie"
	GetMovieDtlActivity = "GetMovieDetail"
)

type MovieService interface {
	GetMovie(ctx context.Context, search string, page int64) (*movie.PaginationMovie, error)
	GetMovieDetail(ctx context.Context, imdbID string) (*movie.MovieDetail, error)
}

type movieService struct {
	activityRepo sa.Repository
	movieDBIntg  omdb.OMDB
}

func New(activityRepo sa.Repository, movieDBIntg omdb.OMDB) MovieService {
	return &movieService{
		activityRepo,
		movieDBIntg,
	}
}

func (u *movieService) GetMovie(ctx context.Context, search string, page int64) (*movie.PaginationMovie, error) {

	var data = &omdb.GetMovieResponse{}
	var err error

	defer func() {
		u.saveActivity(ctx, data, err, GetMovieActivity)
	}()

	data, err = u.movieDBIntg.GetMovie(ctx, search, page)
	if err != nil {
		return nil, err
	}

	return buildPaginationMovie(data), nil
}

func (u *movieService) GetMovieDetail(ctx context.Context, imdbID string) (*movie.MovieDetail, error) {

	var data = &omdb.GetMovieDetailResponse{}
	var err error

	defer func() {
		u.saveActivity(ctx, data, err, GetMovieDtlActivity)
	}()

	data, err = u.movieDBIntg.GetMovieDetail(ctx, imdbID)
	if err != nil {
		return nil, err
	}

	return buildMovieDetail(data), nil
}

func (u *movieService) saveActivity(ctx context.Context, data interface{}, err error, activity string) {

	var msg string
	if err != nil {
		msg = err.Error()
	}

	_ = u.activityRepo.Save(ctx, *buildSearchActivity(data, activity, msg, err == nil))
}

func buildMovieDetail(data *omdb.GetMovieDetailResponse) *movie.MovieDetail {
	return &movie.MovieDetail{
		ImdbID:     data.ImdbID,
		ImdbRating: data.ImdbRating,
		ImdbVotes:  data.ImdbVotes,
		Type:       data.Type,
		Title:      data.Title,
		Year:       data.Year,
		Rated:      data.Rated,
		Released:   data.Released,
		Runtime:    data.Runtime,
		Genre:      data.Genre,
		Director:   data.Director,
		Writer:     data.Writer,
		Actors:     data.Actors,
		Plot:       data.Plot,
		Language:   data.Language,
		Country:    data.Country,
		Awards:     data.Awards,
		Poster:     data.Poster,
		Ratings:    buildMovieRatingList(data.Ratings),
		Metascore:  data.Metascore,
	}
}

func buildMovieRating(data omdb.MovieRating) movie.MovieRating {
	return movie.MovieRating{
		Source: data.Source,
		Value:  data.Value,
	}
}

func buildMovieRatingList(data []omdb.MovieRating) []movie.MovieRating {
	var list []movie.MovieRating
	for _, val := range data {
		list = append(list, buildMovieRating(val))
	}
	return list
}

func buildSearchActivity(data interface{}, activity, message string, isSuccess bool) *sa.SearchActivity {
	dataByte, _ := json.Marshal(data)

	return &sa.SearchActivity{
		Activity: activity,
		Success:  isSuccess,
		Message:  message,
		JsonData: string(dataByte),
	}
}

func buildPaginationMovie(data *omdb.GetMovieResponse) *movie.PaginationMovie {
	var pagination = movie.PaginationMovie{
		TotalResult: data.TotalResult,
	}
	for _, val := range data.Search {
		pagination.Movie = append(pagination.Movie, movie.Movie{
			ImdbID: val.ImdbID,
			Type:   val.Type,
			Title:  val.Title,
			Year:   val.Year,
			Poster: val.Poster,
		})
	}
	return &pagination
}
