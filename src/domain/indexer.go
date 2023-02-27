package domain

import (
	"context"
	"search/src/domain/error"
)

type Indexer struct {
	tokenizer             Tokenizer
	invertIndexCompresser InvertIndexCompresser
}

func NewIndexer(tokenizer Tokenizer, invertIndexCompresser InvertIndexCompresser) *Indexer {
	return &Indexer{tokenizer, invertIndexCompresser}
}

func (indexer *Indexer) CreateDocument(ctx context.Context, document *DocumentCreate) (*DocumentCreate, *error.MyError) {

	tokenizedContent, tokenizeErr := indexer.tokenizer.Tokenize(ctx, document.Content)
	if tokenizeErr != nil {
		return nil, error.NewMyError(tokenizeErr.Code, tokenizeErr.Error())
	}

	document.TokenizedContent = make([]string, len(*tokenizedContent))
	for i, term := range *tokenizedContent {
		document.TokenizedContent[i] = term.Word
	}

	return document, nil
}

func (indexer *Indexer) CreatePosting(ctx context.Context, document *Document) (map[string]*Posting, *error.MyError) {
	wordToPostingMap := make(map[string]*Posting)
	for position, word := range document.TokenizedContent {
		if _, ok := wordToPostingMap[word]; ok {
			wordToPostingMap[word].PositionsInDocument = append(wordToPostingMap[word].PositionsInDocument, position)
		} else {
			positionsInDocument := []int{position}
			wordToPostingMap[word] = NewPosting(document.Id, positionsInDocument)
		}
	}

	return wordToPostingMap, nil
}

func (indexer *Indexer) CreateWordToTermsMap(ctx context.Context, termCompresseds *[]TermCompressed, wordToPostingMap map[string]*Posting) *[]TermCompressedCreate {
	terms := make([]TermCreate, len(*termCompresseds))
	wordToTermsMap := make(map[string]*TermCreate)
	for i, termCompressed := range *termCompresseds {
		invertIndex, decompressErr := indexer.invertIndexCompresser.Decompress(ctx, termCompressed.InvertIndexCompressed)
		if decompressErr != nil {
			panic(decompressErr)
		}
		terms[i].Word = termCompressed.Word
		terms[i].InvertIndex = invertIndex
		wordToTermsMap[termCompressed.Word] = &terms[i]
	}

	// PostingをAppendする
	for wordInDocument, posting := range wordToPostingMap {
		if _, ok := wordToTermsMap[wordInDocument]; ok {
			*wordToTermsMap[wordInDocument].InvertIndex.PostingList = append(*wordToTermsMap[wordInDocument].InvertIndex.PostingList, *posting)
		} else {
			invertIndex := NewInvertIndex(&[]Posting{*posting})
			wordToTermsMap[wordInDocument] = NewTermCreate(wordInDocument, invertIndex)
		}
	}

	// 圧縮
	upsertTermCompresseds := &[]TermCompressedCreate{}
	for wordInDocument := range wordToTermsMap {
		invertIndexCompressed, compressErr := indexer.invertIndexCompresser.Compress(ctx, wordToTermsMap[wordInDocument].InvertIndex)
		if compressErr != nil {
			panic(compressErr)
		}
		termCompressed := NewTermCompressedCreate(wordInDocument, invertIndexCompressed)
		*upsertTermCompresseds = append(*upsertTermCompresseds, *termCompressed)
	}

	return upsertTermCompresseds
}
