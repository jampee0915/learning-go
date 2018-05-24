package main

import (
	"./controllers"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

func main() {
	router := gin.Default()

	router.GET("/:corporateId", func(c *gin.Context) {
		// Paramを処理する
		n := c.Param("corporateId")
		corporateId, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			c.JSON(400, err)
			return
		}
		if corporateId <= 0 {
			c.JSON(400, gin.H{"status": "corporateId should be begger then 0"})
			return
		}

		// データを処理する
		ctrl := controllers.NewCorporates()
		result := ctrl.Get(corporateId)
		if result == nil || reflect.ValueOf(result).IsNil() {
			c.JSON(404, gin.H{})
			return
		}

		c.JSON(200, result)
	})

	router.Run(":8080")
}
