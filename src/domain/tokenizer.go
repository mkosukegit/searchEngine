package domain

import (
	"context"
	"search/src/domain/error"
)

type Tokenizer interface {
	Tokenize(ctx context.Context, content string) (*[]TermCreate, *error.MyError)
}
