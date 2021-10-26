package search_activity

import "context"

type Repository interface {
	Save(ctx context.Context, entity SearchActivity) (err error)
}
