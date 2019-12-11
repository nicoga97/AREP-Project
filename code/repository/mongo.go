package repository

import (
	"arep/config"
	"arep/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type MongoRepository struct {
	client           *mongo.Client
	storesCollection *mongo.Collection
}

var mongoStoresCollection = config.MongoStoresCollection

func NewMongoRepository(config *config.MongoConfiguration) (*MongoRepository, error) {
	client, err := CreateClient(config)
	if err != nil {
		return nil, err
	}
	database := client.Database(config.DbName)
	repository := &MongoRepository{
		client:           client,
		storesCollection: database.Collection(mongoStoresCollection),
	}

	return repository, nil
}

func (r *MongoRepository) UpdateStore(ctx context.Context, storeID string, enabled bool) error {
	filter := bson.D{{"id", storeID}}
	update := bson.D{{"$set",
		bson.D{
			{"enabled", enabled},
		},
	}}
	_, err := r.storesCollection.UpdateOne(ctx, filter, update)

	return err
}

func (r *MongoRepository) GetStores(ctx context.Context, storeIDs []int64) (*[]model.Store, error) {
	var stores []model.Store
	cursor, err := r.storesCollection.Find(ctx, bson.M{"_id": bson.M{"$in": storeIDs}})
	if cursor != nil {
		defer cursor.Close(ctx)
	}
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &stores)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return &stores, nil
}

func CreateClient(config *config.MongoConfiguration) (*mongo.Client, error) {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(config.ConnString).
		SetReadPreference(readpref.SecondaryPreferred())
	db, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	err = db.Ping(ctx, readpref.SecondaryPreferred())
	if err != nil {
		return nil, err
	}

	return db, nil
}
