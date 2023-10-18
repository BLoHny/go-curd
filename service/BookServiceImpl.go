package service

import (
	"context"

	"github.com/blohny/data/request"
	response "github.com/blohny/data/respond"
	error "github.com/blohny/helper"
	"github.com/blohny/model"
	"github.com/blohny/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookServiceImpl(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{BookRepository: bookRepository}
}

func (b *BookServiceImpl) Create(ctx context.Context, request request.BookCreatRequest) {
	book := model.Book{
		Name: request.Name,
	}
	b.BookRepository.Save(ctx, book)
}

func (b *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	book, err := b.BookRepository.FindById(ctx, bookId)
	error.PanicIfError(err)
	b.BookRepository.Delete(ctx, book.Id)
}

func (b *BookServiceImpl) FindAll(ctx context.Context) []response.BookResponse {
	books := b.BookRepository.FindAll(ctx)

	var bookResp []response.BookResponse

	for _, value := range books {
		book := response.BookResponse{Id: value.Id, Name: value.Name}
		bookResp = append(bookResp, book)
	}
	return bookResp

}

func (b *BookServiceImpl) FindById(ctx context.Context, bookId int) response.BookResponse {
	book, err := b.BookRepository.FindById(ctx, bookId)
	error.PanicIfError(err)
	return response.BookResponse(book)
}

func (b *BookServiceImpl) Update(ctx context.Context, request request.BookUpdateRequest) {
	book, err := b.BookRepository.FindById(ctx, request.Id)
	error.PanicIfError(err)

	book.Name = request.Name
	b.BookRepository.Update(ctx, book)
}
