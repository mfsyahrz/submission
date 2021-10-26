package movie_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mfsyahrz/submission/project/internal/domain/search_activity"
	"github.com/mfsyahrz/submission/project/internal/infrastructure/omdb"
	. "github.com/mfsyahrz/submission/project/internal/usecase/movie"
	mockSearchActivity "github.com/mfsyahrz/submission/project/test/mock/domain/search_activity"
	mockMovieInt "github.com/mfsyahrz/submission/project/test/mock/infrastructure/omdb"
)

var (
	ctx = context.Background()
)

type movieServiceMock struct {
	MovieService
	actRepo     *mockSearchActivity.MockRepository
	movieDBIntg *mockMovieInt.MockOMDB
}

func TestGetMovie(t *testing.T) {

	ctrl := gomock.NewController(t)
	ctrl.Finish()

	serviceMock := createMovieServiceMock(ctrl)

	t.Run("success", func(t *testing.T) {
		serviceMock.movieDBIntg.EXPECT().GetMovie(gomock.Any(), gomock.Any(), gomock.Any()).Return(createGetMovieResponse(), nil)
		serviceMock.actRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)

		data, err := serviceMock.GetMovie(ctx, "any", 1)
		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("failed - error get movie", func(t *testing.T) {
		serviceMock.movieDBIntg.EXPECT().GetMovie(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("failed"))
		serviceMock.actRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)

		data, err := serviceMock.GetMovie(ctx, "any", 1)
		assert.Error(t, err)
		assert.Nil(t, data)
	})
}

func TestGetMovieDetail(t *testing.T) {

	ctrl := gomock.NewController(t)
	ctrl.Finish()

	serviceMock := createMovieServiceMock(ctrl)

	t.Run("success", func(t *testing.T) {
		serviceMock.movieDBIntg.EXPECT().GetMovieDetail(gomock.Any(), gomock.Any()).Return(createGetMovieDtlResponse(), nil)
		serviceMock.actRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)

		data, err := serviceMock.GetMovieDetail(ctx, "any")
		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("failed - error get movie detail", func(t *testing.T) {
		serviceMock.movieDBIntg.EXPECT().GetMovieDetail(gomock.Any(), gomock.Any()).Return(nil, errors.New("failed"))
		serviceMock.actRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)

		data, err := serviceMock.GetMovieDetail(ctx, "any")
		assert.Error(t, err)
		assert.Nil(t, data)
	})
}

func createMovieServiceMock(ctrl *gomock.Controller) *movieServiceMock {
	mockActRepo := mockSearchActivity.NewMockRepository(ctrl)
	mockMovieIntg := mockMovieInt.NewMockOMDB(ctrl)
	s := New(mockActRepo, mockMovieIntg)

	return &movieServiceMock{
		MovieService: s,
		actRepo:      mockActRepo,
		movieDBIntg:  mockMovieIntg,
	}
}

func createGetMovieResponse() *omdb.GetMovieResponse {
	return &omdb.GetMovieResponse{
		Response: "Success",
		Search: []omdb.Search{
			omdb.Search{
				ImdbID: "id001",
				Type:   "movie",
				Title:  "any",
			},
		},
		TotalResult: "",
	}
}

func createGetMovieDtlResponse() *omdb.GetMovieDetailResponse {
	return &omdb.GetMovieDetailResponse{
		ImdbID: "id001",
		Type:   "movie",
		Title:  "any",
		Ratings: []omdb.MovieRating{
			omdb.MovieRating{
				Source: "any",
				Value:  "5",
			},
		},
		Metascore: "10",
	}
}

func createSearchActivity() search_activity.SearchActivity {
	return search_activity.SearchActivity{
		Activity: "any",
		Success:  false,
		Message:  "failed",
	}
}
