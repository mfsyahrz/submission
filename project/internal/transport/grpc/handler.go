package grpc

import "github.com/mfsyahrz/submission/project/internal/usecase/movie"

type Handler struct {
	MovieHandler MovieServiceServer
}

func NewHandler(movieService movie.MovieService) *Handler {
	return &Handler{
		MovieHandler: NewMovieHandler(movieService),
	}
}
