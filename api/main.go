package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lolo35/l2l_login/http/controllers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	fmt.Println(port)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	gin.SetMode(gin.ReleaseMode)
	r.GET("/lines", controllers.FetchLines)
	r.GET("/area_lines", controllers.FetchLinesInArea)
	r.GET("/areas", controllers.GetAreas)
	r.POST("/machines", controllers.FetchMachinesForLines)
	r.POST("/login", controllers.Login)

	r.Run(port)
}
