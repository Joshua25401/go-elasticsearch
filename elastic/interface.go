package elastic

import (
	"context"
	"elasticsearch-learn/model"
)

type SearchEngines interface {
	Search(ctx context.Context, keySearch string) ([]model.User, error)
	SearchByType(ctx context.Context, keySearch string, searchType model.UserSearchType) ([]model.User, error)
	InsertData(ctx context.Context, newUser model.User) error
	UpdateData(ctx context.Context, updatedUser model.User) error
	DeleteData(ctx context.Context, deleteTarget model.User) error
	SearchEngineInfo(ctx context.Context)
}
