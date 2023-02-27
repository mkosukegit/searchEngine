package repository

import (
	"context"
	"search/src/domain"
	"search/src/domain/error"
)

type PlainDocumentRepository interface {
	SelectPlainDocument(ctx context.Context) (err error.MyError)
	CreateDocument(ctx context.Context, doc *domain.DocumentCreate) (*domain.Document, *error.MyError)
	FindTermCompressedsByWords(ctx context.Context, words *[]string) (*[]domain.TermCompressed, *error.MyError)
	BulkUpsertTerm(ctx context.Context, terms *[]domain.TermCompressedCreate) *error.MyError
}

type plainDocumentRepository struct {
}

func NewPlainDocumentRepository() PlainDocumentRepository {
	return &plainDocumentRepository{}
}

func (plainDocumentRepository *plainDocumentRepository) SelectPlainDocument(ctx context.Context) (err error.MyError) {

	return err
}

func (r *plainDocumentRepository) CreateDocument(ctx context.Context, doc *domain.DocumentCreate) (*domain.Document, *error.MyError) {

	return nil, nil
}

func (r *plainDocumentRepository) FindTermCompressedsByWords(ctx context.Context, words *[]string) (*[]domain.TermCompressed, *error.MyError) {

	return nil, nil
}

func (r *plainDocumentRepository) BulkUpsertTerm(ctx context.Context, terms *[]domain.TermCompressedCreate) *error.MyError {

	return nil
}
