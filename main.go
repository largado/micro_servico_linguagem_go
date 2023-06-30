package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()

	r.GET("/v0", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"funcionou!": "respondeu!",
		})
	})

	r.Run(":8080")
	
}
