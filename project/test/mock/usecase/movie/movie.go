// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/usecase/movie/movie.go

// Package mock_movie is a generated GoMock package.
package mock_movie

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	movie "github.com/mfsyahrz/submission/project/internal/domain/movie"
)

// MockMovieService is a mock of MovieService interface.
type MockMovieService struct {
	ctrl     *gomock.Controller
	recorder *MockMovieServiceMockRecorder
}

// MockMovieServiceMockRecorder is the mock recorder for MockMovieService.
type MockMovieServiceMockRecorder struct {
	mock *MockMovieService
}

// NewMockMovieService creates a new mock instance.
func NewMockMovieService(ctrl *gomock.Controller) *MockMovieService {
	mock := &MockMovieService{ctrl: ctrl}
	mock.recorder = &MockMovieServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieService) EXPECT() *MockMovieServiceMockRecorder {
	return m.recorder
}

// GetMovie mocks base method.
func (m *MockMovieService) GetMovie(ctx context.Context, search string, page int64) (*movie.PaginationMovie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovie", ctx, search, page)
	ret0, _ := ret[0].(*movie.PaginationMovie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovie indicates an expected call of GetMovie.
func (mr *MockMovieServiceMockRecorder) GetMovie(ctx, search, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovie", reflect.TypeOf((*MockMovieService)(nil).GetMovie), ctx, search, page)
}

// GetMovieDetail mocks base method.
func (m *MockMovieService) GetMovieDetail(ctx context.Context, imdbID string) (*movie.MovieDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieDetail", ctx, imdbID)
	ret0, _ := ret[0].(*movie.MovieDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieDetail indicates an expected call of GetMovieDetail.
func (mr *MockMovieServiceMockRecorder) GetMovieDetail(ctx, imdbID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieDetail", reflect.TypeOf((*MockMovieService)(nil).GetMovieDetail), ctx, imdbID)
}