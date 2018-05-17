package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ndabAP/vue-go-example/backend/routes"
	"runtime"
	"strconv"
)

func main() {
	r := gin.Default()

	r.GET("/api/specs", func(c *gin.Context) {
		c.JSON(200, []string{runtime.GOOS, strconv.Itoa(runtime.NumCPU())})
	})

	r.POST("/api/descriptive/average", routes.Average)
	r.POST("/api/descriptive/standard-deviation", routes.StandardDeviation)
	r.POST("/api/distribution", routes.Distribution)

	r.Run(":3001")
}