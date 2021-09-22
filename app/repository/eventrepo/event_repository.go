package eventrepo

import (
	"context"
	"log"

	"github.com/DLzer/go-gin-api-boilerplate/app/domain/event"
	"github.com/jackc/pgx/v4/pgxpool"
)

// The repo interface
type Repo interface {
	CreateEvent(event *event.Event) (*event.Event, error)
	GetAllEvents() ([]event.Event, error)
	GetEventById(id string) (event.Event, error)
	DeleteEventById(id string) (bool, error)
}

// The repo
type eventRepo struct {
	db *pgxpool.Pool
}

// Return event repo
func NewEventRepo(db *pgxpool.Pool) Repo {
	return &eventRepo{
		db: db,
	}
}

// Create new event
func (e *eventRepo) CreateEvent(event *event.Event) (*event.Event, error) {

	query := `INSERT INTO events (type, time, identity) VALUES ($1, $2, $3) RETURNING id;`
	err := e.db.QueryRow(context.Background(), query, event.Type, event.Time, event.Identity).Scan(&event.ID)
	if err != nil {
		log.Print("Issue inserting event: ", err)
	}

	return event, nil

}

// Return all events
func (e *eventRepo) GetAllEvents() ([]event.Event, error) {

	eventResponse := []event.Event{}

	rows, err := e.db.Query(context.Background(), "SELECT * FROM events ORDER BY ID DESC")
	if err != nil {
		return eventResponse, err
	}

	for rows.Next() {
		var event event.Event

		if err := rows.Scan(&event.ID, &event.Type, &event.Time, &event.Identity); err != nil {
			return eventResponse, err
		}

		eventResponse = append(eventResponse, event)

	}

	return eventResponse, nil

}

// Return a single event by Id
func (e *eventRepo) GetEventById(id string) (event.Event, error) {

	var eventResponse event.Event

	if err := e.db.QueryRow(context.Background(), "SELECT * FROM events WHERE id = $1", id).Scan(
		&eventResponse.ID, &eventResponse.Type, &eventResponse.Time, &eventResponse.Identity); err != nil {
		log.Print("Error scanning rows: ", err)
		return eventResponse, err
	}

	return eventResponse, nil

}

// Delete event by Id
func (e *eventRepo) DeleteEventById(id string) (bool, error) {

	if _, err := e.db.Exec(context.Background(), `DELETE FROM events WHERE id = ?`, id); err != nil {
		log.Print("Error deleting event: ", err)
		return false, err
	}

	return true, nil

}
