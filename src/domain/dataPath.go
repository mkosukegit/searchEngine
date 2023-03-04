package domain

import (
	"os"
	"path"
)

var DoraemonDocumentTsvPath = path.Join(os.Getenv("PROJECT_ROOT"), "./testData.tsv")

var MovieDocumentTsvPath = path.Join(os.Getenv("PROJECT_ROOT"), "cmd/seeds/data/en/movie.tsv")
