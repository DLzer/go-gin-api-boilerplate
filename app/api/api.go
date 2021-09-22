package api

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/DLzer/go-gin-api-boilerplate/app/configs"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var (
	r = gin.Default()
)

type Database struct {
	Conn *pgxpool.Pool
}

// Load in configs. Start the Gin Framework. Connect to the database.
// Initialize Repositories, Services, Controllers, and Route handlers.
// Run GIN on the config port
func run() {

	/*
		====== Setup configs ============
	*/
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config := configs.GetConfig()

	gin.SetMode(gin.ReleaseMode)
	log.Print("GIN running in release mode..")

	/*
		====== Database Initialization ============
	*/
	db, err := Initialize(config)
	if err != nil {
		log.Fatalf("Issue connecting to database %s", err)
	}
	defer db.Conn.Close()

	// Run the App on the configured port
	r.Run(fmt.Sprintf(":%s", config.Port))
}

// Initialize expects the configuration model to attempt connecting to the PG
// database. It will return either the successful connection, or an error.
func Initialize(c configs.Config) (Database, error) {

	db := Database{}
	conn, err := pgxpool.Connect(context.Background(), c.DB.DSN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return db, err
	}
	db.Conn = conn

	log.Println("Database connection established")
	log.Printf("Connected on %s", c.DB.DSN)

	return db, nil
}
