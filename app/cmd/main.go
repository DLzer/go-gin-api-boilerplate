package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/DLzer/go-gin-api-boilerplate/app/configs"
	"github.com/DLzer/go-gin-api-boilerplate/app/controllers"
	"github.com/DLzer/go-gin-api-boilerplate/app/database"
	"github.com/DLzer/go-gin-api-boilerplate/app/middleware"
	"github.com/DLzer/go-gin-api-boilerplate/app/repository/eventrepo"
	"github.com/DLzer/go-gin-api-boilerplate/app/services/eventservice"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var (
	r = gin.Default()
)

// Database stores the
// database connection pool
type Database struct {
	Conn *pgxpool.Pool
}

// Run the application
func main() {
	Run()
}

// Run loads in configs. Start the Gin Framework. Connect to the database.
// Initialize Repositories, Services, Controllers, and Route handlers.
// Rurs GIN on the config port
func Run() {

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

	/*
		====== Database Migration ============
	*/
	_, _ = Migrate(db.Conn)

	/*
		====== Middleware ============
	*/
	r.Use(middleware.CORSMiddleware())

	/*
		====== Repository Initialization ============
	*/
	eventRepo := eventrepo.NewEventRepo(db.Conn)

	/*
		====== Service Initialization ============
	*/
	eventService := eventservice.NewEventService(eventRepo)

	/*
		====== Controller Initialization ============
	*/
	eventController := controllers.NewEventController(eventService)

	/*
		====== Routes ============
	*/

	// Event routes
	events := r.Group("/events")
	{
		events.POST("/create", eventController.PostEvent)
		events.GET("", eventController.GetAllEvents)
		events.GET("/:id", eventController.GetEventById)
		events.GET("/delete/:id", eventController.DeleteEventById)
	}

	// Run the App on the configured port
	err = r.Run(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		log.Fatal("Could not run app: ", err)
	}

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

// Migrate will run the database migrations
func Migrate(db *pgxpool.Pool) (bool, error) {

	if _, err := database.Migrate(db); err != nil {
		log.Print("Error migrating: ", err)
		return false, err
	}

	log.Print("Database migration complete")
	return true, nil

}
