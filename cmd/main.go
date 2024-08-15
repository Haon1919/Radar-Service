package main

import (
	  "database/sql"
    "fmt"
    "log"
		"net/url"
    "os"

		"radar_service/internal/handlers"
    "github.com/gin-gonic/gin"
		_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
  
	encodedPassword := url.QueryEscape(dbPassword)

	dbConnStr := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
        dbUser, encodedPassword, dbHost, dbPort, dbName)
	
  db, err := sql.Open("sqlserver", dbConnStr)
  if err != nil {
  	log.Fatal("Error creating connection pool: ", err.Error())
  }
  defer db.Close()

  // Verify the connection
  err = db.Ping()
  if err != nil {
    log.Fatal("Error connecting to the database: ", err.Error())
  }

  log.Println("Successfully connected to SQL Server")
	
	router := gin.Default()

  // Create an instance of your handler with the db connection
  h := handlers.NewHandler(db)

  // Register routes with handlers
  router.GET("/ping", h.PingHandler)

  router.Run(":8080")
}

