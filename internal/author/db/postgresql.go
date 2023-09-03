package author

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"rest-api-learning/internal/author"
	"rest-api-learning/pkg/client/postgresql"
	"rest-api-learning/pkg/logging"
	"strings"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *repository) Create(ctx context.Context, author *author.Author) error {
	q := `
		INSERT INTO author 
		    (name) 
		VALUES
		    ($1) 
		RETURNING id 
		`
	r.logger.Trace(fmt.Sprintf("SQL Quary: %s", formatQuery(q)))
	if err := r.client.QueryRow(ctx, q, author.Name).Scan(&author.ID); err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			newErr := fmt.Errorf(fmt.Sprintf("SQL error: %s, Detait: %s, Where: %s, Code: %s, SQL state: %s",
				pgError.Message, pgError.Detail, pgError.Where, pgError.Code, pgError.SQLState()))
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (r *repository) FindAll(ctx context.Context) ([]author.Author, error) {
	q := `
		SELECT id, name 
		FROM public.author
		`
	r.logger.Trace(fmt.Sprintf("SQL Quary: %s", formatQuery(q)))
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	authors := make([]author.Author, 0)
	for rows.Next() {
		var ath author.Author

		err := rows.Scan(&ath.ID, &ath.Name)
		if err != nil {
			return nil, err
		}

		authors = append(authors, ath)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (author.Author, error) {
	q := `
		SELECT id, name 
		FROM public.author
		`
	r.logger.Trace(fmt.Sprintf("SQL Quary: %s", formatQuery(q)))
	var ath author.Author
	err := r.client.QueryRow(ctx, q, id).Scan(&ath.ID, &ath.Name)
	if err != nil {
		return author.Author{}, err
	}
	return ath, nil
}

func (r *repository) Update(ctx context.Context, author author.Author) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) author.Repository {
	return &repository{
		client: client,
		logger: logger,
	}

}
