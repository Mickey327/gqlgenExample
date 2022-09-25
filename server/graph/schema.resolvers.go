package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"log"

	"github.com/Mickey327/graphqlapp/graph/generated"
	"github.com/Mickey327/graphqlapp/graph/model"
)

// AddBook is the resolver for the addBook field.
func (r *mutationResolver) AddBook(ctx context.Context, book model.BookInput) (bool, error) {
	q := `
		INSERT INTO book(name, description, author_id) 
		VALUES ($1, $2, $3)
	`
	if _, err := r.Postgres.Pool.Query(ctx, q, book.Title, book.Description, book.AuthorID); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code)
			log.Fatal(newErr)
			return false, nil
		}
		return false, err
	}
	log.Println("INSERTED BOOK")
	return true, nil
}

// GetAllBooks is the resolver for the getAllBooks field.
func (r *queryResolver) GetAllBooks(ctx context.Context) ([]*model.Book, error) {
	q := `SELECT b.id, b.name, b.description, a.id, a.first_name, a.last_name FROM BOOK b LEFT JOIN AUTHOR a ON a.id = b.author_id`
	rows, err := r.Postgres.Pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	books := make([]*model.Book, 0)
	for rows.Next() {
		var book model.Book
		var author model.Author
		err = rows.Scan(&book.ID, &book.Title, &book.Description, &author.ID, &author.FirstName, &author.LastName)
		if err != nil {
			return nil, err
		}
		book.Author = &author
		books = append(books, &book)
	}
	return books, nil
}

// GetBook is the resolver for the getBook field.
func (r *queryResolver) GetBook(ctx context.Context, id string) (*model.Book, error) {
	q := `SELECT b.id, b.name, b.description, a.id, a.first_name, a.last_name FROM BOOK b LEFT JOIN AUTHOR a ON a.id = b.author_id WHERE b.id = $1`
	var book model.Book
	var author model.Author
	err := r.Postgres.Pool.QueryRow(ctx, q, id).Scan(&book.ID, &book.Title, &book.Description, &author.ID, &author.FirstName, &author.LastName)
	if err != nil {
		return nil, err
	}
	book.Author = &author

	return &book, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
