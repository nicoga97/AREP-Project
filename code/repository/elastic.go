package repository

import (
	"arep/config"
	"arep/model"
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"log"
	"strconv"
)

type ElasticRepository struct {
	client *elastic.Client
}

var ElasticStoresIndex = config.ElasticStoresIndex

func NewElasticRepository(config *config.ElasticConfiguration) (*ElasticRepository, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(config.DbUrl),
		elastic.SetBasicAuth(config.UserName, config.Password))
	if err != nil {
		return nil, err
	}
	repository := &ElasticRepository{
		client: client,
	}

	return repository, nil
}

func (r *ElasticRepository) UpdateStore(ctx context.Context, storeID string, enabled bool) error {
	update := struct {
		Enabled bool `json:"enabled"`
	}{Enabled: enabled}
	_, err := r.client.Update().
		Index(ElasticStoresIndex).
		Type("_doc").
		Id(storeID).
		Doc(update).
		Do(ctx)

	return err
}

func (r *ElasticRepository) GetStores(ctx context.Context, storeIDs []int64) (*[]model.Store, error) {
	items := make([]*elastic.MultiGetItem, len(storeIDs))
	for _, storeID := range storeIDs {
		item := elastic.NewMultiGetItem()
		item.Index(ElasticStoresIndex)
		item.Id(strconv.FormatInt(storeID, 10))
		item.Type("_doc")
		items = append(items, item)
	}
	resp, err := r.client.Mget().Add(items...).Do(ctx)
	if err != nil {
		return nil, err
	}
	stores := make([]model.Store, len(storeIDs))
	for _, doc := range resp.Docs {
		var store model.Store
		err := json.Unmarshal(*doc.Source, &store)
		if err != nil {
			log.Print("error parsing store", err.Error())
		} else {
			stores = append(stores, store)
		}
	}

	return &stores, nil
}
