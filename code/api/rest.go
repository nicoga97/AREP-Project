package api

import (
	"arep/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func StartServer(elasticController *controller.StoreController,
	documentDBController *controller.StoreController) {

	server := gin.New()
	apiPrefix := server.Group("api")
	{
		elasticGroup := apiPrefix.Group("elastic")
		{
			elasticGroup.GET("update-store", elasticController.UpdateStore)
			elasticGroup.GET("get-stores", elasticController.GetStores)
		}sdfkkfñls
		documentDBGroup := apiPrefix.Group("document-db")
		{
			documentDBGroup.GET("update-store", documentDBController.UpdateStore)
			documentDBGroup.GET("get-stores", documentDBController.GetStores)
		}
	}

	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
