package book

import (
	"context"
	"fmt"
	"rest-api-learning/internal/author"
	"rest-api-learning/internal/book"
	"rest-api-learning/pkg/client/postgresql"
	"rest-api-learning/pkg/logging"
	"strings"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r *repository) Create(ctx context.Context, book *book.Book) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) FindOne(ctx context.Context, id string) (book.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Update(ctx context.Context, book book.Book) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *repository) FindAll(ctx context.Context) ([]book.Book, error) {
	q := `
		SELECT id, name 
		FROM public.book
		`
	r.logger.Trace(fmt.Sprintf("SQL Quary: %s", formatQuery(q)))
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	books := make([]book.Book, 0)
	for rows.Next() {
		var bk book.Book

		err := rows.Scan(&bk.ID, &bk.Name)
		if err != nil {
			return nil, err
		}

		sq := `
		SELECT
				a.id, a.name
		FROM book_authors ba
		JOIN public.author a on ba.author_id = a.id
		WHERE ba.book_id = $1;
`
		authorRows, err := r.client.Query(ctx, sq, bk.ID)
		if err != nil {
			return nil, err
		}

		authors := make([]author.Author, 0)
		for authorRows.Next() {
			var ath author.Author

			err := rows.Scan(&ath.ID, &ath.Name)
			if err != nil {
				return nil, err
			}
			authors = append(authors, ath)
		}
		bk.Authors = authors
		books = append(books, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}
func NewRepository(client postgresql.Client, logger *logging.Logger) book.Repository {
	return &repository{
		client: client,
		logger: logger,
	}

}
