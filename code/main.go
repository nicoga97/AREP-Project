package main

import (
	"arep/api"
	"arep/config"
	"arep/controller"
	"arep/repository"
	"arep/service"
	"log"
)

func main() {
	mongoRepository, err := repository.NewMongoRepository(config.GetMongoConfig())
	if err != nil {
		log.Fatal("Can not connect with MongoDB")
	}
	log.Print("connection success with MongoDB")
	elasticRepository, err := repository.NewElasticRepository(config.GetElasticConfig())
	if err != nil {
		log.Fatal("Can not connect with ElasticSearch")
	}
	log.Print("connection success with ElasticSearch")
	mongoService := service.StoreService{Repository: mongoRepository}
	elasticService := service.StoreService{Repository: elasticRepository}

	mongoController := controller.NewStoreController(mongoService)
	elasticController := controller.NewStoreController(elasticService)

	api.StartServer(elasticController, mongoController)
}
