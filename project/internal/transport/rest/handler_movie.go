package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfsyahrz/submission/project/internal/usecase/movie"
	"github.com/mfsyahrz/submission/project/pkg/util"
	"github.com/spf13/cast"
)

type MovieHandler struct {
	service movie.MovieService
}

func NewMovieHandler(e *echo.Echo, service movie.MovieService) {
	handler := &MovieHandler{
		service: service,
	}

	e.GET("/movie/health", handler.Ping)
	e.GET("/movie", handler.GetMovie)
	e.GET("/movie/detail", handler.GetMovieDetail)

}

func (h *MovieHandler) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "healthy")
}

func (h *MovieHandler) GetMovie(c echo.Context) error {

	var resp Response

	search := c.QueryParam("search")
	page := cast.ToInt64(c.QueryParam("p"))

	if search == "" {
		msg := "Please review your input, search cannot be empty"
		resp.Message = msg
		return c.JSON(http.StatusBadRequest, resp)
	}

	data, err := h.service.GetMovie(
		c.Request().Context(),
		search,
		util.SetPageForPagination(page),
	)
	if err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Success = true
	resp.Message = "Success"
	resp.Data = data

	return c.JSON(http.StatusOK, resp)
}

func (h *MovieHandler) GetMovieDetail(c echo.Context) error {

	var resp Response

	imdbID := c.QueryParam("imdbID")

	if imdbID == "" {
		msg := "Please review your input, imdbID param cannot be empty"
		resp.Message = msg
		return c.JSON(http.StatusBadRequest, resp)
	}

	data, err := h.service.GetMovieDetail(
		c.Request().Context(),
		imdbID,
	)
	if err != nil {
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			})
	}

	resp.Success = true
	resp.Message = "Success"
	resp.Data = data

	return c.JSON(http.StatusOK, resp)
}
