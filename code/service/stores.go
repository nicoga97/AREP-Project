package service

import (
	"arep/model"
	"arep/repository"
	"context"
)

type StoreService struct {
	Repository repository.StoreRepository
}

func (r *StoreService) UpdateEnabledStore(c context.Context, storeID string, enabled bool) error {
	return r.Repository.UpdateStore(c, storeID, enabled)
}

func (r *StoreService) GetStores(c context.Context, storeIDs []int64) (*[]model.Store, error) {
	return r.Repository.GetStores(c, storeIDs)
}
