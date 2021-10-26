package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	domain "github.com/mfsyahrz/submission/project/internal/domain/search_activity"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) domain.Repository {
	return &repository{db}
}

func (r *repository) Save(ctx context.Context, entity domain.SearchActivity) (err error) {
	rows, err := r.db.Exec(`INSERT INTO search_activity (activity, success, message, json_data, request_dt)
				 VALUES 
				 ($1 , $2 , $3, $4, now())`,
		entity.Activity, entity.Success, entity.Message, entity.JsonData)
	if err != nil {
		return err
	}

	_, err = rows.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
