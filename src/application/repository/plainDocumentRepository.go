package repository

import (
	"context"
	"fmt"
	"search/src/domain"
	"search/src/domain/error"
	"search/src/domain/error/code"
	"search/src/middleware/db"
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
	in, err := db.Db.Prepare("INSERT INTO document(content) VALUES(?)")

	if err != nil {
		fmt.Println("connect failed DB")
		return nil, error.NewMyError(code.DBConnectError, err)
	} else {
		fmt.Println("connect success DB")
	}

	defer db.Db.Close()

	result, err := in.Exec("goro")

	if err != nil {
		return nil, error.NewMyError(code.DBConnectError, err)
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return nil, error.NewMyError(code.DBConnectError, err)
	}

	fmt.Println(lastId)
	return convertDocumentEntSchemaToEntity(int(lastId), doc), nil
}

func (r *plainDocumentRepository) FindTermCompressedsByWords(ctx context.Context, words *[]string) (*[]domain.TermCompressed, *error.MyError) {

	return nil, nil
}

func (r *plainDocumentRepository) BulkUpsertTerm(ctx context.Context, terms *[]domain.TermCompressedCreate) *error.MyError {

	return nil
}

func convertDocumentEntSchemaToEntity(id int, d *domain.DocumentCreate) *domain.Document {
	return &domain.Document{
		Id:               int64(id),
		Content:          d.Content,
		TokenizedContent: d.TokenizedContent,
	}
}
