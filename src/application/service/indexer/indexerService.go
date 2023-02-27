package indexer

import (
	"context"
	"log"
	"search/src/application/repository"
	"search/src/domain"
	errors "search/src/domain/error"
	"search/src/domain/error/code"

	"golang.org/x/sync/errgroup"
)

type IndexerService interface {
	CreateIndex(ctx context.Context) (err error)
}

type indexerService struct {
	plainDocumentRepository repository.PlainDocumentRepository
}

func NewUserInfoService(plainDocumentRepository repository.PlainDocumentRepository) IndexerService {
	return &indexerService{plainDocumentRepository}
}

func (indexerService *indexerService) CreateIndex(ctx context.Context) (err error) {

	// initialize term usecase
	zlibInvertIndexCompresser := domain.NewZlibInvertIndexCompresser()

	// create tokenizer
	jaKagomeTokenizer := domain.NewJaKagomeTokenizer()
	indexer := domain.NewIndexer(jaKagomeTokenizer, zlibInvertIndexCompresser)

	// load documents
	docs, err := domain.LoadDocumentsFromTsv(domain.DoraemonDocumentTsvPath)
	if err != nil {
		log.Fatalf("failed loading documents: %v", err)
	}
	size := 200
	docLists := domain.SplitSlice(*docs, size)
	goroutines := len(docLists)

	// index documents concurrently
	log.Println("start indexing documents")
	var eg errgroup.Group
	for i := 0; i < goroutines; i++ {
		docList := docLists[i]
		eg.Go(func() error {
			for _, doc := range docList {
				// ドキュメントを作成
				docCreate := domain.NewDocumentCreate(doc, []string{})
				createdDoc, _ := indexer.CreateDocument(ctx, docCreate)

				// ドキュメントを登録
				documentDetail, createErr := indexerService.plainDocumentRepository.CreateDocument(ctx, createdDoc)
				if createErr != nil {
					return errors.NewMyError(createErr.Code, createErr.Error())
				}

				// ポスティング作成
				wordToPostingMap, _ := indexer.CreatePosting(ctx, documentDetail)

				// 登録済みインデックスを取得
				termCompresseds, termErr := indexerService.plainDocumentRepository.FindTermCompressedsByWords(ctx, &docCreate.TokenizedContent)
				if termErr != nil {
					return errors.NewMyError(termErr.Code, termErr.Error())
				}

				upsertTermCompresseds := indexer.CreateWordToTermsMap(ctx, termCompresseds, wordToPostingMap)

				// 転値インデックスの登録
				upsertErr := indexerService.plainDocumentRepository.BulkUpsertTerm(ctx, upsertTermCompresseds)
				if upsertErr != nil {
					return errors.NewMyError(code.Unknown, upsertErr)
				}
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		log.Println(err)
	}

	return nil
}
