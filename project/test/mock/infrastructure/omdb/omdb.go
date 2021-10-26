// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/infrastructure/omdb/omdb.go

// Package mock_omdb is a generated GoMock package.
package mock_omdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	omdb "github.com/mfsyahrz/submission/project/internal/infrastructure/omdb"
)

// MockOMDB is a mock of OMDB interface.
type MockOMDB struct {
	ctrl     *gomock.Controller
	recorder *MockOMDBMockRecorder
}

// MockOMDBMockRecorder is the mock recorder for MockOMDB.
type MockOMDBMockRecorder struct {
	mock *MockOMDB
}

// NewMockOMDB creates a new mock instance.
func NewMockOMDB(ctrl *gomock.Controller) *MockOMDB {
	mock := &MockOMDB{ctrl: ctrl}
	mock.recorder = &MockOMDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOMDB) EXPECT() *MockOMDBMockRecorder {
	return m.recorder
}

// GetMovie mocks base method.
func (m *MockOMDB) GetMovie(ctx context.Context, search string, page int64) (*omdb.GetMovieResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovie", ctx, search, page)
	ret0, _ := ret[0].(*omdb.GetMovieResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovie indicates an expected call of GetMovie.
func (mr *MockOMDBMockRecorder) GetMovie(ctx, search, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovie", reflect.TypeOf((*MockOMDB)(nil).GetMovie), ctx, search, page)
}

// GetMovieDetail mocks base method.
func (m *MockOMDB) GetMovieDetail(ctx context.Context, imdbID string) (*omdb.GetMovieDetailResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieDetail", ctx, imdbID)
	ret0, _ := ret[0].(*omdb.GetMovieDetailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieDetail indicates an expected call of GetMovieDetail.
func (mr *MockOMDBMockRecorder) GetMovieDetail(ctx, imdbID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieDetail", reflect.TypeOf((*MockOMDB)(nil).GetMovieDetail), ctx, imdbID)
}
