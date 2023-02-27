package domain

import (
	"context"
	"search/src/domain/error"
)

type InvertIndexCompresser interface {
	Compress(ctx context.Context, invertIndexes *InvertIndex) (*InvertIndexCompressed, *error.MyError)
	Decompress(ctx context.Context, invertIndexes *InvertIndexCompressed) (*InvertIndex, *error.MyError)
}
