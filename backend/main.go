package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tech-rounak/true-beacon/controllers"
	"github.com/tech-rounak/true-beacon/routes"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Origin, Cache-Control, X-Requested-With,token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")
	fmt.Println("----------------- PORT -----------------", port)
	if port == "" {
		port = "8000"
	}
	var SECRET_KEY string = os.Getenv("SECRET_KEY")
	fmt.Println("---------------- SECRET_KEY ---------------", SECRET_KEY)
	router := gin.New()
	router.Use(CORSMiddleware())
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.ShareRoutes(router)

	if err := controllers.ParseCsvAndInsertIntoDB(); err != nil {
		log.Fatal(err)
	}
	router.Run(":" + port)

}
