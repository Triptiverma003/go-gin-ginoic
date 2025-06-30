// package database

// import (
// 	"log"
// 	"os"

// 	"github.com/Triptiverma003/go-gin-auth/model"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
	
// )

// var DBConn *gorm.DB

// func ConnectDB(){
// 	dsn := os.Getenv("dsn")

// 	db , err := gorm.Open(mysql.Open(dsn) , &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Error),
// 	})

// 	if err != nil {
// 		panic("Error in database connection")
// 	}

// 	log.Println("Database connection successfull")

// 	db.AutoMigrate(new(model.User))

// 	DBConn = db
// }

package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Triptiverma003/go-gin-auth/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {
	user := os.Getenv("db_username")
	password := os.Getenv("db_password")
	dbname := os.Getenv("db_name")
	host := os.Getenv("dbhostname")
	port := os.Getenv("port_number")

	// ✅ Construct the DSN from separate environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		log.Fatal("❌ Error in database connection:", err)
	}

	log.Println("✅ Database connection successful")

	db.AutoMigrate(new(model.User)) // Auto-migrate User model

	DBConn = db
}