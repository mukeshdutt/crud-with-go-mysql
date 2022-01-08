package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Welcome to the golang crud app.")
	})

	r.GET("/api/employees", AllEmployee)
	r.GET("/api/employees/:id", EmployeeByID)
	r.POST("/api/employees", SaveEmployee)
	r.PUT("/api/employees/:id", UpdateEmployee)
	r.DELETE("/api/employees/:id", DeleteEmployee)

	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
