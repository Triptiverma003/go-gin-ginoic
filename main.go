// package main

// import (
// 	"log"
// 	"os"

// 	"github.com/Triptiverma003/go-gin-auth/database"
// 	routes "github.com/Triptiverma003/go-gin-auth/router"
// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// )

// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("Error in loading env...")
// 	}

// 	database.ConnectDB()
// }

// func main() {
// 	sqlDb, err := database.DBConn.DB()
// 	if err != nil {
// 		log.Println("error in connecting Db")
// 	}
// 	defer sqlDb.Close()

// 	port := os.Getenv("port")
// 	if port == "" {
// 		port = "8001"
// 	}

// 	gin.SetMode(gin.ReleaseMode)
// 	router := gin.New()

// 	// ✅ Register CORS middleware BEFORE routes and router.Run
// 	router.Use(cors.New(cors.Config{
// 		AllowOrigins:     []string{"http://localhost:3000"}, // Frontend origin
// 		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
// 		AllowHeaders:     []string{"Origin", "Auth-token", "token", "Content-Type"},
// 		AllowCredentials: true,
// 	}))

// 	routes.SetUpRoutes(router)

// 	// ✅ Start server AFTER middleware is registered
// 	log.Fatal(router.Run(":" + port))
// }

package main

import (
	"log"
	"os"

	"github.com/Triptiverma003/go-gin-auth/database"
	routes "github.com/Triptiverma003/go-gin-auth/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found or failed to load")
	}

	database.ConnectDB()
}

func main() {
	sqlDb, err := database.DBConn.DB()
	if err != nil {
		log.Fatal("❌ Failed to get underlying SQL DB:", err)
	}
	defer func() {
		if err := sqlDb.Close(); err != nil {
			log.Println("⚠️ Error closing DB:", err)
		}
	}()

	// Get port from env or use default
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "8001"
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Replace with Vercel domain on deploy
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Auth-token", "token", "Content-Type"},
		AllowCredentials: true,
	}))

	// Route setup
	routes.SetUpRoutes(router)

	// Optional: health check route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong from auth"})
	})

	log.Printf("✅ Auth server running on port %s", serverPort)
	log.Fatal(router.Run(":" + serverPort))
}