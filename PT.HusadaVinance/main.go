package main

import (
	controller "PT.HusadaVinance/Controllers"
	connection "PT.HusadaVinance/Models/Connection"

	"github.com/gin-gonic/gin"
)

func main() {

	port := ":8080"
	r := gin.Default()
	connection.ConnectDatabase()

	 r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if c.Request.Method == "OPTIONS" {
            return
        }
        c.Next()
    })



	//###BEGIN WEB API PT. HUSADA VINANCE TBK
	// Get data
	r.GET("/api/PT.HusadaVinance/GetKeuangan", controller.GetHusadaVinanceControllersByID)

	//Insert data
	r.POST("/api/PT.HusadaVinance/InsertKeuangan", controller.InsertHusadaVinanceControllers)

	// Update data
	r.PUT("/api/PT.HusadaVinance/UpdateKeuangan", controller.UpdateHusadavinanceControllers)

	//Delete data
	r.DELETE("/api/PT.HusadaVinance/DeleteKeuangan", controller.DeleteHusadaVinanceControllers)
	//###END WEB API PT.HUSADA VINANCE TBK


	r.Run(port)
}