package Controller

import (
	"net/http"

	"github.com/zxc10110/Fruit_Market/Model"

	"github.com/gin-gonic/gin"
)

func GetAllFruits(c *gin.Context) {
	data, err := Model.GetAllFruit()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}
