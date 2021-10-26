package omdb_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mfsyahrz/submission/project/internal/config"
	. "github.com/mfsyahrz/submission/project/internal/infrastructure/omdb"
	"github.com/stretchr/testify/assert"
)

func TestGetMovie(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		getMovieMock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			res := &GetMovieResponse{
				Response: "Success",
				Search: []Search{
					Search{
						ImdbID: "id001",
						Type:   "movie",
						Title:  "saw",
						Year:   "2002",
					},
				},
				TotalResult: "120",
			}

			want, _ := json.Marshal(res)

			w.Header().Add("Content-Type", "application/json")
			w.Write(want)
		}))
		defer getMovieMock.Close()

		e := New(&config.OMDB{
			Host: getMovieMock.URL,
		})
		ctx := context.Background()

		data, err := e.GetMovie(ctx, "any", 1)
		assert.NoError(t, err)
		assert.Equal(t, &GetMovieResponse{
			Response: "Success",
			Search: []Search{
				Search{
					ImdbID: "id001",
					Type:   "movie",
					Title:  "saw",
					Year:   "2002",
				},
			},
			TotalResult: "120",
		}, data)

	})

	t.Run("failed - retrieve data", func(t *testing.T) {
		getMovieMock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			res := &GetMovieResponse{
				Error: "failed",
			}

			want, _ := json.Marshal(res)

			w.Header().Add("Content-Type", "application/json")
			w.Write(want)
		}))
		defer getMovieMock.Close()

		e := New(&config.OMDB{
			Host: getMovieMock.URL,
		})
		ctx := context.Background()

		data, err := e.GetMovie(ctx, "any", 1)
		assert.Error(t, err)
		assert.Nil(t, data)

	})

	t.Run("failed - error make request", func(t *testing.T) {
		getMovieMock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer getMovieMock.Close()

		e := New(&config.OMDB{
			Host: "",
		})
		ctx := context.Background()

		data, err := e.GetMovie(ctx, "any", 1)
		assert.Error(t, err)
		assert.Nil(t, data)

	})

	t.Run("failed - integration error", func(t *testing.T) {
		getMovieMock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer getMovieMock.Close()

		e := New(&config.OMDB{
			Host: getMovieMock.URL,
		})
		ctx := context.Background()

		data, err := e.GetMovie(ctx, "any", 1)
		assert.Error(t, err)
		assert.Nil(t, data)

	})

}

func TestGetMovieDetail(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		getMovieMock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			res := &GetMovieDetailResponse{
				Response: "Success",
				ImdbID:   "id001",
				Title:    "any",
				Ratings: []MovieRating{
					MovieRating{
						Source: "any",
						Value:  "any",
					},
				},
			}

			want, _ := json.Marshal(res)

			w.Header().Add("Content-Type", "application/json")
			w.Write(want)
		}))
		defer getMovieMock.Close()

		e := New(&config.OMDB{
			Host: getMovieMock.URL,
		})
		ctx := context.Background()

		data, err := e.GetMovieDetail(ctx, "any")
		assert.NoError(t, err)
		assert.Equal(t, &GetMovieDetailResponse{
			Response: "Success",
			ImdbID:   "id001",
			Title:    "any",
			Ratings: []MovieRating{
				MovieRating{
					Source: "any",
					Value:  "any",
				},
			},
		}, data)

	})

	t.Run("failed - retrieve data", func(t *testing.T) {
		getMovieMock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			res := &GetMovieDetailResponse{
				Error: "failed",
			}

			want, _ := json.Marshal(res)

			w.Header().Add("Content-Type", "application/json")
			w.Write(want)
		}))
		defer getMovieMock.Close()

		e := New(&config.OMDB{
			Host: getMovieMock.URL,
		})
		ctx := context.Background()

		data, err := e.GetMovieDetail(ctx, "any")
		assert.Error(t, err)
		assert.Nil(t, data)

	})

	t.Run("failed - error making request", func(t *testing.T) {
		getMovieMock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer getMovieMock.Close()

		e := New(&config.OMDB{
			Host: "",
		})
		ctx := context.Background()

		data, err := e.GetMovieDetail(ctx, "any")
		assert.Error(t, err)
		assert.Nil(t, data)

	})

	t.Run("failed - integration error", func(t *testing.T) {
		getMovieMock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer getMovieMock.Close()

		e := New(&config.OMDB{
			Host: getMovieMock.URL,
		})
		ctx := context.Background()

		data, err := e.GetMovieDetail(ctx, "any")
		assert.Error(t, err)
		assert.Nil(t, data)

	})

}
