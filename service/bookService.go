package service

import (
	"context"

	"github.com/blohny/data/request"
	response "github.com/blohny/data/respond"
)

type BookService interface {
	Create(ctx context.Context, request request.BookCreatRequest)
	Update(ctx context.Context, request request.BookUpdateRequest)
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) response.BookResponse
	FindAll(ctx context.Context) []response.WebResponse
}
