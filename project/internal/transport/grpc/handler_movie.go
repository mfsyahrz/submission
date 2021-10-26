package grpc

import (
	"context"

	entity "github.com/mfsyahrz/submission/project/internal/domain/movie"
	"github.com/mfsyahrz/submission/project/internal/usecase/movie"
	"github.com/mfsyahrz/submission/project/pkg/util"
)

type MovieHandler struct {
	service movie.MovieService
	UnimplementedMovieServiceServer
}

func NewMovieHandler(service movie.MovieService) MovieServiceServer {
	return &MovieHandler{
		service,
		UnimplementedMovieServiceServer{},
	}
}

func (h *MovieHandler) GetMovie(ctx context.Context, req *GetRequest) (*MoviePagination, error) {
	data, err := h.service.GetMovie(ctx, req.Search, util.SetPageForPagination(req.Page))
	if err != nil {
		return &MoviePagination{
			Status:  "500",
			Message: err.Error(),
		}, nil
	}

	return &MoviePagination{
		Status:  "00",
		Message: "Success",
		Data:    buildMovieMeta(*data, req.Page),
	}, nil
}

func (h *MovieHandler) GetMovieDetail(ctx context.Context, req *IdRequest) (*MovieResponse, error) {
	data, err := h.service.GetMovieDetail(ctx, req.Id)
	if err != nil {
		return &MovieResponse{
			Status:  "500",
			Message: err.Error(),
		}, nil
	}

	return &MovieResponse{
		Status:  "00",
		Message: "Success",
		Data:    buildMovieDetail(*data),
	}, nil
}

func buildMovieMeta(data entity.PaginationMovie, page int64) *MovieMeta {
	var pagination = &MovieMeta{
		TotalResult: data.TotalResult,
		Page:        page,
	}
	for _, val := range data.Movie {
		pagination.List = append(pagination.List, &Movie{
			ImdbId: val.ImdbID,
			Type:   val.Type,
			Title:  val.Title,
			Year:   val.Year,
			Poster: val.Poster,
		})
	}
	return pagination
}

func buildMovieDetail(data entity.MovieDetail) *MovieDetail {
	return &MovieDetail{
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

func buildMovieRating(data entity.MovieRating) *MovieRating {
	return &MovieRating{
		Source: data.Source,
		Value:  data.Value,
	}
}

func buildMovieRatingList(data []entity.MovieRating) []*MovieRating {
	var list []*MovieRating
	for _, val := range data {
		list = append(list, buildMovieRating(val))
	}
	return list
}
