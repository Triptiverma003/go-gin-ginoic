package main

import (
	"log"
	"os"

	"github.com/Triptiverma003/go-gin-auth/database"
	routes "github.com/Triptiverma003/go-gin-auth/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init(){
	if err := godotenv.Load(); err!=nil{
		log.Println("Error in loading env...")
	}

	database.ConnectDB()
}

func main () {

	sqlDb , err := database.DBConn.DB()

	if err!= nil {
		log.Println("error in connecting Db")
	}

	defer sqlDb.Close()

	port := os.Getenv("port")

	if port == ""{
		port = "8001"
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	routes.SetUpRoutes(router)

	log.Fatal(router.Run(":" + port))

}