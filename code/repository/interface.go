package repository

import (
	"arep/model"
	"context"
)

type StoreRepository interface {
	UpdateStore(context.Context, string, bool) error
	GetStores(context.Context, []int64) (*[]model.Store, error)
}
