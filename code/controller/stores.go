package controller

import (
	"arep/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type StoreController struct {
	service service.StoreService
}

type StoresRequest struct {
	StoreIDs []int64 `json:"store_ids"`
}

func NewStoreController(service service.StoreService) *StoreController {
	return &StoreController{
		service: service,
	}
}

func (r *StoreController) UpdateStore(c *gin.Context) {
	storeID := c.Param("store_id")
	enabled, err := strconv.ParseBool(c.Param("enabled"))
	if err != nil {
		err = r.service.UpdateEnabledStore(c, storeID, enabled)
	}
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, storeID)
	}
}

func (r *StoreController) GetStores(c *gin.Context) {
	var reqBody StoresRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil || len(reqBody.StoreIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	stores, err := r.service.GetStores(c, reqBody.StoreIDs)
	if err != nil {
		_ = c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, stores)
	}
}
