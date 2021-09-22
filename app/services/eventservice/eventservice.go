package eventservice

import (
	"github.com/DLzer/go-gin-api-boilerplate/app/domain/event"
	"github.com/DLzer/go-gin-api-boilerplate/app/repository/eventrepo"
)

// The service interface
type EventService interface {
	CreateEvent(event *event.Event) (*event.Event, error)
	GetAllEvents() ([]event.Event, error)
	GetEventById(id string) (event.Event, error)
	DeleteEventById(id string) (bool, error)
}

// The service
type eventService struct {
	Repo eventrepo.Repo
}

// Return new event service
func NewEventService(repo eventrepo.Repo) EventService {
	return &eventService{
		Repo: repo,
	}
}

// Creates a new event
func (es *eventService) CreateEvent(event *event.Event) (*event.Event, error) {
	return es.Repo.CreateEvent(event)
}

// Gets all events
func (es *eventService) GetAllEvents() ([]event.Event, error) {
	return es.Repo.GetAllEvents()
}

// Gets an event by its Id
func (es *eventService) GetEventById(id string) (event.Event, error) {
	return es.Repo.GetEventById(id)
}

// Deletes an event by its Id
func (es *eventService) DeleteEventById(id string) (bool, error) {
	return es.Repo.DeleteEventById(id)
}
