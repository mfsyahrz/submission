package omdb

import (
	"context"
	"errors"

	"github.com/go-resty/resty/v2"
	"github.com/mfsyahrz/submission/project/internal/config"
	"github.com/spf13/cast"
)

type OMDB interface {
	GetMovie(ctx context.Context, search string, page int64) (*GetMovieResponse, error)
	GetMovieDetail(ctx context.Context, imdbID string) (*GetMovieDetailResponse, error)
}

type omdb struct {
	config     *config.OMDB
	httpClient *resty.Client
}

func New(cfg *config.OMDB) OMDB {
	client := resty.New().SetDebug(true)
	return &omdb{cfg, client}
}

func (o *omdb) GetMovie(ctx context.Context, search string, page int64) (*GetMovieResponse, error) {
	var (
		serverResp GetMovieResponse
	)

	resp, err := o.httpClient.R().
		SetResult(&serverResp).
		SetError(&serverResp).
		SetQueryParams(map[string]string{
			"apikey": o.config.APIKey,
			"s":      search,
			"page":   cast.ToString(page),
		}).
		Get(o.config.Host)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New("failed to connect to moviedb")
	}

	if serverResp.Error != "" {
		return nil, errors.New(serverResp.Error)
	}

	return &serverResp, nil

}

func (o *omdb) GetMovieDetail(ctx context.Context, imdbID string) (*GetMovieDetailResponse, error) {
	var (
		serverResp GetMovieDetailResponse
	)

	resp, err := o.httpClient.R().
		SetResult(&serverResp).
		SetError(&serverResp).
		SetQueryParams(map[string]string{
			"apikey": o.config.APIKey,
			"i":      imdbID,
		}).
		Get(o.config.Host)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New("failed to connect to moviedb")
	}

	if serverResp.Error != "" {
		return nil, errors.New(serverResp.Error)
	}

	return &serverResp, nil
}
